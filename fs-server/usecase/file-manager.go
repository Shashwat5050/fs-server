package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type FileManager struct {
	log           *otelzap.Logger
	obs           objectStorage
	bucket        string
	VolumeDir     string
	keepPeriod    time.Duration
	cleanPeriod   time.Duration
	sftpClient    *sftp.Client
	sftpConnected bool
}

type settings func(*FileManager)

func WithVolumeDir(volumeDir string) func(*FileManager) {
	return func(f *FileManager) {
		f.VolumeDir = volumeDir
	}
}

func WithKeepPeriod(keepTime string) func(*FileManager) {
	return func(f *FileManager) {
		if keepTime == "" {
			return
		}
		kp, err := time.ParseDuration(keepTime)
		if err != nil {
			f.log.Warn("cannot parse keep time", zap.Error(err), zap.String("keepTime", keepTime))
			return
		}
		f.keepPeriod = kp
	}
}

func WithCleanPeriod(cleanPeriod string) func(*FileManager) {
	return func(f *FileManager) {
		if cleanPeriod == "" {
			return
		}

		cp, err := time.ParseDuration(cleanPeriod)
		if err != nil {
			f.log.Warn("cannot parse clean period", zap.Error(err), zap.String("cleanPeriod", cleanPeriod))
			return
		}

		f.cleanPeriod = cp
	}
}

func WithBucket(bucket string) func(*FileManager) {
	return func(f *FileManager) {
		if bucket == "" {
			f.log.Warn("provided bucket name is empty")
			return
		}

		f.bucket = bucket
	}
}

const (
	defaultVolumeDir   = "/tmp"
	defaultBucket      = "backups"
	defaultkeepTime    = 30
	defaultCleanPeriod = time.Hour * 3
	defaultCacheDir    = "/tmp/mods"
)

func NewLocalFileManager(logger *otelzap.Logger, obs objectStorage, opts ...settings) *FileManager {
	fm := &FileManager{
		log:         logger,
		obs:         obs,
		bucket:      defaultBucket,
		VolumeDir:   defaultVolumeDir,
		keepPeriod:  defaultkeepTime,
		cleanPeriod: defaultCleanPeriod,
	}

	for _, opt := range opts {
		opt(fm)
	}

	return fm
}

func (fm *FileManager) RunTrashCleaner(ctx context.Context) error {
	timer := time.NewTimer(fm.cleanPeriod)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			go func(ctx context.Context, timer *time.Timer) {
				defer timer.Reset(fm.cleanPeriod)

				err := filepath.Walk(fm.VolumeDir, func(path string, info os.FileInfo, err error) error {
					if ctxErr := ctx.Err(); ctxErr != nil {
						return err
					}

					if err != nil {
						return err
					}

					if !strings.HasPrefix(info.Name(), ".trash-") {
						return nil
					}

					if time.Since(info.ModTime()) > fm.keepPeriod {
						if err := os.Remove(path); err != nil {
							return fmt.Errorf("could not remove trash file %s: %w", path, err)
						}
					}

					return nil
				})

				if err != nil {
					fm.log.Error("could not clean trash", zap.Error(err))
				}
			}(ctx, timer)
		}
	}
}

func (fm *FileManager) GetFileStat(path string) (fs.FileInfo, error) {
	path = filepath.Join(fm.VolumeDir, path)
	if err := fm.validatePath(path); err != nil {
		return nil, err
	}

	return os.Stat(path)
}

func (fm *FileManager) GetFileData(path string) (string, error) {
	path = filepath.Join(fm.VolumeDir, path)
	if err := fm.validatePath(path); err != nil {
		info, err := fm.GetFileStat(path)
		if err != nil {
			return "", err
		}

		if info.IsDir() {
			return "", errors.New("Given Path leads to directory not File")
		}
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// check if file is editable or not
	mimeType := http.DetectContentType(data)

	editableMIMETypes := []string{"text/plain", "text/markdown", "application/json"}

	for _, mt := range editableMIMETypes {
		if strings.Contains(mimeType, mt) {
			return string(data), nil
		}
	}

	return "", fmt.Errorf("Invalid content type")
}

func (fm *FileManager) SetFileData(path string, input []byte) error {
	path = filepath.Join(fm.VolumeDir, path)
	if err := fm.validatePath(path); err != nil {
		return err
	}

	err := os.WriteFile(path, input, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (fm *FileManager) ListFilePath(path string) ([]fs.DirEntry, error) {
	path = filepath.Join(fm.VolumeDir, path)
	if err := fm.validatePath(path); err != nil {
		return nil, err
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := make([]fs.DirEntry, 0, len(files))
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".trash-") {
			continue
		}

		result = append(result, file)
	}

	return result, nil
}

func (fm *FileManager) CreateDir(path, name string) error {
	path = filepath.Join(fm.VolumeDir, path, name)
	if err := fm.validatePath(path); err != nil {
		return err
	}

	return os.Mkdir(path, 0770)
}

func (fm *FileManager) DownloadFile(path, name string) (int64, io.ReadCloser, error) {
	path = filepath.Join(fm.VolumeDir, path, name)
	if err := fm.validatePath(path); err != nil {
		return 0, nil, err
	}

	file, err := os.Stat(path)
	if err != nil {
		return 0, nil, err
	}

	if file.IsDir() {
		return 0, nil, fmt.Errorf("path %s is a directory", path)
	}

	fd, err := os.Open(filepath.Clean(path))
	if err != nil {
		return 0, nil, err
	}

	return file.Size(), fd, nil
}

func (fm *FileManager) UploadFile(ctx context.Context, path, filename string, data io.Reader) error {

	buf := make([]byte, 1024)

	fd, err := fm.CreateFile(path, filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			n, err := data.Read(buf)
			if err != nil {
				if err == io.EOF {
					return nil
				}

				return err
			}

			if _, err := fd.Write(buf[:n]); err != nil {
				return err
			}
		}
	}
}

func (fm *FileManager) CreateFile(path, filename string) (io.WriteCloser, error) {
	path = filepath.Join(fm.VolumeDir, path, filename)
	if err := fm.validatePath(path); err != nil {
		return nil, err
	}

	file, err := os.Create(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	// Change the file permissions to 770 (rw-rw----)
	if err := os.Chmod(path, 0770); err != nil {
		file.Close()
		return nil, err
	}
	fmt.Println("Changes file permissions")

	return file, err
}

func (fm *FileManager) DeleteFile(path, filename string) error {
	fullPath := filepath.Join(fm.VolumeDir, path, filename)
	if err := fm.validatePath(fullPath); err != nil {
		return err
	}

	deletedPath := filepath.Join(fm.VolumeDir, path, ".trash-"+filename)

	err := os.Rename(fullPath, deletedPath)
	if err != nil {
		if os.IsExist(err) {

			fmt.Println("deleting file/directory :", deletedPath)
			err := os.RemoveAll(deletedPath)
			if err != nil {
				return fmt.Errorf("could not trash file: %w", err)
			}

			err = os.Rename(fullPath, deletedPath)
			if err != nil {
				log.Println(err, "Trash file existing error")
				return fmt.Errorf("Trash file existing error: %w", err)
			}
			return nil
		}
		log.Println(err, "error while moving it to trash")
		return fmt.Errorf("could not move to trash: %w", err)
	}

	return nil
}
func (fm *FileManager) UninstallServer(server_name string) error {
	fullPath := filepath.Join(fm.VolumeDir, server_name)
	if err := fm.validatePath(fullPath); err != nil {
		return err
	}

	// deletedPath := filepath.Join(fm.VolumeDir, path, ".trash-"+filename)

	err := os.RemoveAll(fullPath)
	log.Println(err, "error while moving it to trash")
	if err != nil {
		return fmt.Errorf("could not move to trash: %w", err)
	}

	return nil
}

func (fm *FileManager) RecoverFile(path, filename string) error {
	fullPath := filepath.Join(fm.VolumeDir, path, filename)
	if err := fm.validatePath(fullPath); err != nil {
		return err
	}

	deletedPath := filepath.Join(fm.VolumeDir, path, ".trash-"+filename)
	_, err := os.Stat(deletedPath)
	if err != nil {
		return fmt.Errorf("could not recover: %w", err)
	}

	err = os.Rename(deletedPath, fullPath)
	if err != nil {
		return fmt.Errorf("could not recover: %w", err)
	}

	return nil
}

func (fm *FileManager) MoveFile(path, filename, newPath, newFilename string) error {
	path = filepath.Join(fm.VolumeDir, path, filename)
	if err := fm.validatePath(path); err != nil {
		return err
	}

	newPath = filepath.Join(fm.VolumeDir, newPath, newFilename)
	if err := fm.validatePath(newPath); err != nil {
		return err
	}

	return os.Rename(path, newPath)
}

// Copy File and Diretory from path to newpath ,saved with newName
func (fm *FileManager) CopyFile(path, filename, newPath, newName string) error {

	generalPath := filepath.Join(path, filename)

	path = filepath.Join(fm.VolumeDir, path, filename)
	if err := fm.validatePath(path); err != nil {
		return err
	}

	newPath = fm.generateNewCopyName(newPath, newName)
	if err := fm.validatePath(newPath); err != nil {
		return err

	}

	fileStats, err := fm.GetFileStat(generalPath)
	if err != nil {
		return err
	}

	if fileStats.IsDir() {
		fmt.Println("Coping Directory...")
		if err := fm.copyDirectory(path, newPath); err != nil {
			return err
		}
		return nil
	}

	src, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create(filepath.Clean(newPath))
	if err != nil {
		return err
	}

	err = copyFilePermissions(path, newPath)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func (fm *FileManager) CompressFile(ctx context.Context, path, filename string) error {
	inputPath := filepath.Join(fm.VolumeDir, path, filename)
	if err := fm.validatePath(inputPath); err != nil {
		return err
	}

	outputPath := filepath.Join(fm.VolumeDir, path, filename+".zip")
	if err := fm.validatePath(outputPath); err != nil {
		return err
	}

	f, err := os.Open(filepath.Clean(inputPath))
	if err != nil {
		return err
	}

	defer f.Close()

	zf, err := os.OpenFile(filepath.Clean(outputPath), os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer zf.Close()

	fm.log.Debug("compressing file", zap.String("path", inputPath), zap.String("outputPath", outputPath))

	if err := fm.ArchiveFile(ctx, inputPath, zf); err != nil {
		return err
	}

	return nil
}

func (fm *FileManager) BulkCompressFile(ctx context.Context, path string, filenames []string) (string, error) {
	if len(filenames) == 0 {
		return "", nil
	}

	inputPath := filepath.Join(fm.VolumeDir, path)
	if err := fm.validatePath(inputPath); err != nil {
		return "", err
	}

	// generate an archive name
	outputPath := fm.generateNewCopyName(path, "archive.zip")
	if err := fm.validatePath(outputPath); err != nil {
		return "", err
	}

	f, err := os.OpenFile(filepath.Clean(outputPath), os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}

	defer f.Close()

	outputPath, err = filepath.Rel(fm.VolumeDir, outputPath)
	if err != nil {
		return "", err
	}

	if err := fm.ArchiveFiles(ctx, inputPath, filenames, f); err != nil {
		return "", err
	}

	return outputPath, nil
}

// // InstallAndExtractFile will install file from url and extract it at directory located at installation path
// func (fm *FileManager) InstallAndExtractFile(ctx context.Context, url string, installationPath string) (string, error) {
// 	// Download the file from the URL
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	installationPath=filepath.Join(fm.VolumeDir,installationPath)
// 	log.Println(installationPath)
// 	// Check if the installation path exists, if not, create it
// 	if _, err := os.Stat(installationPath); os.IsNotExist(err) {
// 		err = os.MkdirAll(installationPath, 0755)
// 		if err != nil {
// 			log.Println(err, "error while creating directory")
// 			return "", err
// 		}
// 	}

// 	// Extract the file to the installation directory
// 	var targetPath string
// 	fileName := filepath.Base(url)
// 	if filepath.Ext(fileName)==".sql"{
// 		targetPath = filepath.Join(installationPath,fileName)

// 	}else{
// 		targetPath = filepath.Join(installationPath, strings.TrimSuffix(fileName, filepath.Ext(fileName)))
// 	}
// 	// log.Println("fileName", fileName)

// 	// targetPath := filepath.Join(installationPath, strings.TrimSuffix(fileName, filepath.Ext(fileName)))
// 	log.Println("target path", targetPath)
// 	file, err := os.Create(targetPath)
// 	if err != nil {
// 		log.Println("error while creating target file path", err)
// 		return "", err
// 	}
// 	defer file.Close()

// 	_, err = io.Copy(file, resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Check if the downloaded file is a zip archive
// 	if strings.HasSuffix(fileName, ".zip") {
// 		err = fm.extractZip(targetPath, installationPath)
// 		if err != nil {
// 			log.Println("error while extracting file", err)
// 			return "", err
// 		}
// 		// Remove the zip file after extraction
// 		err = os.Remove(targetPath)
// 		if err != nil {
// 			fmt.Println("Failed to remove zip file:", err)
// 		}
// 	}

// 	return targetPath, nil
// }

// // InstallAndExtractfile with cach function
// func (fm *FileManager) InstallAndExtractFile(ctx context.Context, url string, installationPath string) (string, error){
// 	var targetPath string
// 	// Defining cache path

// 	cachePath:=filepath.Join(defaultCacheDir)
// 	installationPath = filepath.Join(fm.VolumeDir, installationPath)

// 	log.Println("Cache path:", cachePath)
// 	log.Println("Installation path:", installationPath)

// // Ensure cache directory exists
// if _, err := os.Stat(cachePath); os.IsNotExist(err) {
// 	log.Println("Cache directory does not exist. Creating cache directory.")
// 	if err := os.MkdirAll(cachePath, 0755); err != nil {
// 		log.Println("Error creating cache directory:", err)
// 		return "", err
// 	}
// }

// 	// Extract filename from URL
// 	fileName := filepath.Base(url)

// 	if filepath.Ext(fileName)==".sql"{
// 		// implement caching
// 		// Check if the file already exists in the cache
// 	targetPath = filepath.Join(installationPath,fileName)
// 	}else{
// 		targetPath = filepath.Join(installationPath, strings.TrimSuffix(fileName, filepath.Ext(fileName)))
// 	}

// 	// Check if the file already exists in the cache
// 	cachedFilePath := filepath.Join(cachePath, fileName)
// 	if _, err := os.Stat(cachedFilePath); err == nil {
// 		log.Println("File found in cache. Copying from cache to installation path.")
// 		if _,err:=os.Stat(installationPath);os.IsNotExist(err){
// 			log.Println("target path does not exist.Creating target path.")
// 			if err:=os.MkdirAll(installationPath,0755);err!=nil{
// 				log.Println("Error creating target directory:",err)
// 				return "",err
// 			}
// 		}
// 		// Create the target file
// 	file, err := os.Create(targetPath)
// 	if err != nil {
// 		log.Println("Error while creating target file path:", err)
// 		return "", err
// 	}
// 	defer file.Close()
// 		if err := fm.copyFile(cachedFilePath, targetPath); err != nil {
// 			return "", err
// 		}
// // If the copied file is a zip file, extract it
// if strings.HasSuffix(fileName, ".zip") {
// 	log.Println("Extracting zip file.")
// 	// Check if any file in the zip exists in the installation path
// 	if err := fm.checkZipContents(targetPath, installationPath); err != nil {
// 		return "", err
// 	}
// 	if err := fm.extractZip(targetPath, installationPath); err != nil {
// 		log.Println("Error while extracting file:", err)
// 		return "", err
// 	}
// 	// Remove the zip file after extraction
// 	if err := os.Remove(targetPath); err != nil {
// 		log.Println("Failed to remove zip file:", err)
// 	}
// }

// return targetPath, nil
// 	}
// 	// Download the file from the URL
// 	log.Println("Downloading file from URL:", url)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	// Check if the installation path exists, if not, create it
// 	if _, err := os.Stat(installationPath); os.IsNotExist(err) {
// 		err = os.MkdirAll(installationPath, 0755)
// 		if err != nil {
// 			log.Println("Error while creating directory:", err)
// 			return "", err
// 		}
// 	}

// 	// Create the target file
// 	file, err := os.Create(targetPath)
// 	if err != nil {
// 		log.Println("Error while creating target file path:", err)
// 		return "", err
// 	}
// 	defer file.Close()

// 	// Copy the downloaded file to the target file
// 	_, err = io.Copy(file, resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Cache the downloaded file
// 	log.Println("Caching the downloaded file.")
// 	if err := fm.copyFile(targetPath, cachedFilePath); err != nil {
// 		log.Println("Error while caching file:", err)
// 		// Don't return error, continue execution
// 	}

// 	// Check if the downloaded file is a zip archive and extract it
// 	if strings.HasSuffix(fileName, ".zip") {
// 		log.Println("Extracting zip file.")
// 		// Check if any file in the zip exists in the installation path
// 		if err := fm.checkZipContents(targetPath, installationPath); err != nil {
// 			return "", err
// 		}

// 		// Extract the zip file
// 		if err := fm.extractZip(targetPath, installationPath); err != nil {
// 			log.Println("Error while extracting file:", err)
// 			return "", err
// 		}
// 		// Remove the zip file after extraction
// 		if err := os.Remove(targetPath); err != nil {
// 			log.Println("Failed to remove zip file:", err)
// 		}
// 	}

// 	return targetPath, nil

// }

// InstallAndExtractFile downloads and extracts a file from a URL to a specified installation path,
// with caching functionality.
func (fm *FileManager) InstallAndExtractFile(ctx context.Context, url string, installationPath string, disableCache bool) (string, error) {

	// check if disableCache is true then do not store installed mod in cache
	// Define cache path
	cachePath := filepath.Join(defaultCacheDir)
	installationPath = filepath.Join(fm.VolumeDir, installationPath)

	log.Println("Cache path:", cachePath)
	log.Println("Installation path:", installationPath)

	// Ensure cache directory exists
	if err := fm.ensureCacheDirectoryExists(cachePath); err != nil {
		return "", err
	}

	// Ensure installation path exists
	if err := fm.ensureInstallationPathExists(installationPath); err != nil {
		return "", err
	}
	// Extract filename from URL
	fileName := filepath.Base(url)

	targetPath := fm.getTargetPath(fileName, installationPath)

	// Check if the file already exists in the cache
	cachedFilePath := filepath.Join(cachePath, fileName)
	if _, err := os.Stat(cachedFilePath); err == nil {
		log.Println("File found in cache. Copying from cache to installation path.")
		if err := fm.copyFile(cachedFilePath, targetPath); err != nil {
			return "", err
		}

		// If the copied file is a zip file, extract it
		if strings.HasSuffix(fileName, ".zip") {
			log.Println("Extracting zip file.")
			if err := fm.checkZipAndExtract(targetPath, installationPath); err != nil {
				log.Println("Error while extracting file:", err)
				return "", err
			}
			// Remove the zip file after extraction
			if err := os.Remove(targetPath); err != nil {
				log.Println("Failed to remove zip file:", err)
			}
		}

		return targetPath, nil
	}

	// Download the file from the URL
	log.Println("Downloading file from URL:", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if err := fm.saveDownloadedFile(resp.Body, targetPath); err != nil {
		return "", err
	}
	// if disable cache is true then don't store the file in cache
	if disableCache == false {
		// Cache the downloaded file
		log.Println("Caching the downloaded file.")
		if err := fm.cacheFile(targetPath, cachedFilePath); err != nil {
			log.Println("Error while caching file:", err)
			// Don't return error, continue execution
		}
	}

	// If the downloaded file is a zip archive, extract it
	if strings.HasSuffix(fileName, ".zip") {
		log.Println("Extracting zip file.")
		if err := fm.checkZipAndExtract(targetPath, installationPath); err != nil {
			log.Println("Error while extracting file:", err)
			return "", err
		}
		// Remove the zip file after extraction
		if err := os.Remove(targetPath); err != nil {
			log.Println("Failed to remove zip file:", err)
		}
	}

	return targetPath, nil
}

// func (fm *FileManager) ensureCacheDirectoryExists(cachePath string) error {
// 	// Ensure cache directory exists
// 	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
// 		log.Println("Cache directory does not exist. Creating cache directory.")
// 		if err := os.MkdirAll(cachePath, 0755); err != nil {
// 			log.Println("Error creating cache directory:", err)
// 			return err
// 		}
// 	}
// 	return nil
// }

// func(fm *FileManager)ensureInstallationPathExists(installationPath string)error{
// 	// Ensure cache directory exists
// 	if _, err := os.Stat(installationPath); os.IsNotExist(err) {
// 		log.Println("installationPath does not exist. Creating installationPath.")
// 		if err := os.MkdirAll(installationPath, 0755); err != nil {
// 			log.Println("Error creating installationPath:", err)
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (fm *FileManager) getTargetPath(fileName, installationPath string) string {
// 	if filepath.Ext(fileName) == ".sql" {
// 		return filepath.Join(installationPath, fileName)
// 	}
// 	return filepath.Join(installationPath, strings.TrimSuffix(fileName, filepath.Ext(fileName)))
// }

// func (fm *FileManager) saveDownloadedFile(body io.ReadCloser, targetPath string) error {
// 	// Create the target file
// 	file, err := os.Create(targetPath)
// 	if err != nil {
// 		log.Println("Error while creating target file path:", err)
// 		return err
// 	}
// 	defer file.Close()

// 	// Copy the downloaded file to the target file
// 	_, err = io.Copy(file, body)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (fm *FileManager) cacheFile(sourcePath, cachedFilePath string) error {
// 	return fm.copyFile(sourcePath, cachedFilePath)
// }

// func (fm *FileManager) checkZipAndExtract(zipPath, extractionPath string) error {
// 	// Check if any file in the zip exists in the installation path
// 	if err := fm.checkZipContents(zipPath, extractionPath); err != nil {
// 		return err
// 	}

// 	// Extract the zip file
// 	return fm.extractZip(zipPath, extractionPath)
// }

func (fm *FileManager) GetDiskSpace(path string) (int64, error) {
	var size int64
	root := filepath.Join(fm.VolumeDir, path)

	err := filepath.Walk(root, func(_ string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("error walking through directory: %w", err)
	}

	return size, nil
}
