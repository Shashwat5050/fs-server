package usecase

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func (fm *FileManager) ArchiveFile(ctx context.Context, inputPath string, pw io.Writer) error {
	zipWriter := zip.NewWriter(pw)
	defer zipWriter.Close()

	totalSize := uint64(0)
	walker := func(path string, info os.FileInfo, err error) error {
		if ctx.Err() != nil {
			fm.log.Error("context error", zap.Error(ctx.Err()))
			return ctx.Err()
		}

		if err != nil {
			fm.log.Error("cannot walk", zap.Error(err))
			return err
		}

		if info.IsDir() || info.Mode()&os.ModeSymlink == os.ModeSymlink {
			return nil
		}

		nextFile, err := os.Open(filepath.Clean(path))
		if err != nil {
			fm.log.Error("cannot open file", zap.Error(err))
			return err
		}

		defer nextFile.Close()

		relPath, err := filepath.Rel(inputPath, path)
		if err != nil {
			fm.log.Error("cannot get relative path", zap.Error(err))
			return err
		}

		f, err := zipWriter.Create(relPath)
		if err != nil {
			fm.log.Error("cannot create zip file", zap.Error(err))
			return err
		}

		n, err := io.Copy(f, nextFile)
		if err != nil {
			fm.log.Error("cannot copy file", zap.Error(err))
			return err
		}

		totalSize += uint64(n)

		return nil
	}

	if err := filepath.Walk(inputPath, walker); err != nil {
		return err
	}

	return nil
}

func (fm *FileManager) ArchiveFiles(ctx context.Context, inputPath string, filenames []string, pw io.Writer) error {
	zw := zip.NewWriter(pw)

	for _, filename := range filenames {
		filePath := filepath.Join(inputPath, filename)
		if err := fm.validatePath(filePath); err != nil {
			return err
		}

		f, err := os.Open(filepath.Clean(filePath))
		if err != nil {
			return err
		}

		defer f.Close()

		fInfo, err := f.Stat()
		if err != nil {
			return err
		}

		// if file is a directory or symlink, skip it
		if fInfo.IsDir() || fInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
			continue
		}

		fw, err := zw.Create(filename)
		if err != nil {
			return err
		}

		if _, err := io.Copy(fw, f); err != nil {
			return err
		}
	}

	if err := zw.Close(); err != nil {
		return err
	}

	return nil
}

func (fm *FileManager) generateNewCopyName(newPath, newName string) string {

	fmt.Println(fm.VolumeDir, newPath, newName)
	filePath := filepath.Join(fm.VolumeDir, newPath, newName)

	var err error
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return filePath
	}

	i := 1
	for {
		fileName, ext := strings.TrimSuffix(newName, filepath.Ext(newName)), filepath.Ext(newName)
		filePath = filepath.Join(fm.VolumeDir, newPath, fileName+"("+strconv.Itoa(i)+")"+ext)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			break
		}
		i++
	}

	return filePath
}

func (fm *FileManager) validatePath(path string) error {
	vpath := filepath.ToSlash(path)
	vpath = filepath.Clean(vpath)
	if !filepath.IsAbs(vpath) {
		return fmt.Errorf("path %s is not absolute", path)
	}

	if !filepath.HasPrefix(vpath, fm.VolumeDir) {
		return fmt.Errorf("path %s exceeds permitted directory", path)
	}

	return nil
}

// extractZip extracts a zip archive to the specified directory
func (fm *FileManager) extractZip(zipFile, targetDir string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		path := filepath.Join(targetDir, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		dst, err := os.Create(path)
		if err != nil {
			return err
		}
		defer dst.Close()

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	return nil
}

func (fm *FileManager) copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create destination file
	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Copy the content from source to destination
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}

// Function to check if any file in the zip exists in the installation path
func (fm *FileManager) checkZipContents(zipFilePath string, installationPath string) error {
	// Open the zip archive for reading
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer r.Close()

	// Iterate through each file in the zip archive
	for _, f := range r.File {
		// Check if the file already exists in the installation path
		existingFilePath := filepath.Join(installationPath, f.Name)
		if _, err := os.Stat(existingFilePath); err == nil {
			if err := os.Remove(zipFilePath); err != nil {
				log.Println("Failed to remove zip file:", err)

			}

			return errors.New("target path contains file(s) that are available in the new zip")
		}
	}
	return nil
}

func (fm *FileManager) ensureCacheDirectoryExists(cachePath string) error {
	// Ensure cache directory exists
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		log.Println("Cache directory does not exist. Creating cache directory.")
		if err := os.MkdirAll(cachePath, 0755); err != nil {
			log.Println("Error creating cache directory:", err)
			return err
		}
	}
	return nil
}

func (fm *FileManager) ensureInstallationPathExists(installationPath string) error {
	// Ensure cache directory exists
	if _, err := os.Stat(installationPath); os.IsNotExist(err) {
		log.Println("installationPath does not exist. Creating installationPath.")
		if err := os.MkdirAll(installationPath, 0755); err != nil {
			log.Println("Error creating installationPath:", err)
			return err
		}
	}
	return nil
}

func (fm *FileManager) getTargetPath(fileName, installationPath string) string {
	if filepath.Ext(fileName) == ".sql" {
		return filepath.Join(installationPath, fileName)
	}
	return filepath.Join(installationPath, strings.TrimSuffix(fileName, filepath.Ext(fileName)))
}

func (fm *FileManager) saveDownloadedFile(body io.ReadCloser, targetPath string) error {
	// Create the target file
	file, err := os.Create(targetPath)
	if err != nil {
		log.Println("Error while creating target file path:", err)
		return err
	}
	defer file.Close()

	// Copy the downloaded file to the target file
	_, err = io.Copy(file, body)
	if err != nil {
		return err
	}
	return nil
}

func (fm *FileManager) cacheFile(sourcePath, cachedFilePath string) error {
	return fm.copyFile(sourcePath, cachedFilePath)
}

func (fm *FileManager) checkZipAndExtract(zipPath, extractionPath string) error {
	// Check if any file in the zip exists in the installation path
	if err := fm.checkZipContents(zipPath, extractionPath); err != nil {
		return err
	}

	// Extract the zip file
	return fm.extractZip(zipPath, extractionPath)
}

func (fm *FileManager) copyDirectory(src, dest string) error {
	// Create the destination directory
	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		return err
	}

	// Read the contents of the source directory
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcEntry := filepath.Join(src, entry.Name())
		destEntry := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			// Recursively copy subdirectory
			if err := fm.copyDirectory(srcEntry, destEntry); err != nil {
				return err
			}
		} else {
			// Copy the file
			if err := fm.copyFile(srcEntry, destEntry); err != nil {
				return err
			}
		}
	}

	return nil
}
