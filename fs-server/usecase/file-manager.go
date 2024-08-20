package usecase

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type FileManager struct {
	log         *otelzap.Logger
	obs         objectStorage
	bucket      string
	VolumeDir   string
	keepPeriod  time.Duration
	cleanPeriod time.Duration
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
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	// fmt.Println(string(data))
	return string(data), nil
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

	return os.Mkdir(path, 0750)
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
	path = filepath.Join(fm.VolumeDir, path)
	if err := fm.validatePath(path); err != nil {
		return err
	}

	path = filepath.Join(path, filename)

	buf := make([]byte, 1024)
	fd, err := os.Create(filepath.Clean(path))
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

	return os.Create(filepath.Clean(path))
}

func (fm *FileManager) DeleteFile(path, filename string) error {
	fullPath := filepath.Join(fm.VolumeDir, path, filename)
	if err := fm.validatePath(fullPath); err != nil {
		return err
	}

	deletedPath := filepath.Join(fm.VolumeDir, path, ".trash-"+filename)

	err := os.Rename(fullPath, deletedPath)
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

func (fm *FileManager) CopyFile(path, filename, newPath, newName string) error {
	path = filepath.Join(fm.VolumeDir, path, filename)
	if err := fm.validatePath(path); err != nil {
		return err
	}

	newPath = fm.generateNewCopyName(newPath, newName)
	if err := fm.validatePath(newPath); err != nil {
		return err
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
