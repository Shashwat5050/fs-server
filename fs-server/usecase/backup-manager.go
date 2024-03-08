package usecase

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func (fm *FileManager) BackupFile(ctx context.Context, inputPath string) (uint64, string, error) {
	inputPath = filepath.Join(fm.VolumeDir, inputPath)
	if err := fm.validatePath(inputPath); err != nil {
		fm.log.Error("invalid path", zap.Error(err))
		return 0, "", err
	}

	// opening a game server folder to compress
	file, err := os.Open(filepath.Clean(inputPath))
	if err != nil {
		fm.log.Error("cannot open file", zap.Error(err))
		return 0, "", err
	}

	defer file.Close()

	zipFileName := filepath.Base(inputPath) + time.Now().Format(time.RFC3339) + ".zip"

	pr, pw := io.Pipe()
	errgp, errctx := errgroup.WithContext(ctx)

	errgp.Go(func() error {
		err := fm.obs.WriteObject(errctx, fm.bucket, zipFileName, pr)
		if err != nil {
			fm.log.Error("cannot write object", zap.Error(err))
			return err
		}

		fm.log.Info("object created", zap.String("bucket", fm.bucket), zap.String("key", zipFileName))

		return nil
	})

	errgp.Go(func() error {
		defer pw.Close()

		// creating zip file to write compressed data
		if err := fm.ArchiveFile(errctx, inputPath, pw); err != nil {
			fm.log.Error("cannot archive file", zap.Error(err))
			return err
		}

		return nil
	})

	obsErr := errgp.Wait()
	if obsErr != nil {
		return 0, "", obsErr
	}

	size, err := fm.obs.GetObjectSize(ctx, fm.bucket, zipFileName)
	if err != nil {
		fm.log.Error("cannot get object size", zap.Error(err))

		return 0, "", err
	}

	return uint64(size), zipFileName, obsErr
}

func (fm *FileManager) RestoreFile(ctx context.Context, path, filename string) error {
	fm.log.Info("restoring file", zap.String("path", path), zap.String("filename", filename))

	object, size, err := fm.obs.GetObject(ctx, fm.bucket, filename)
	if err != nil {
		fm.log.Error("cannot get object", zap.Error(err))

		return err
	}

	defer object.Close()

	buff := bytes.NewBuffer(make([]byte, 0, size))
	n, err := io.Copy(buff, object)
	if err != nil {
		fm.log.Error("cannot copy object", zap.Error(err))

		return err
	}

	reader := bytes.NewReader(buff.Bytes())

	zipReader, err := zip.NewReader(reader, n)
	if err != nil {
		fm.log.Error("cannot read zip file", zap.Error(err))

		return err
	}

	destination, err := filepath.Abs(filepath.Join(fm.VolumeDir, path))
	if err != nil {
		fm.log.Error("cannot get absolute path", zap.Error(err))

		return err
	}

	for _, f := range zipReader.File {
		if err := unzipFile(f, destination); err != nil {
			fm.log.Error("cannot unzip file", zap.Error(err))

			return err
		}
	}

	fm.log.Info("file restored", zap.String("path", path), zap.String("filename", filename))

	return nil
}

func sanitizeArchivePath(d, t string) (v string, err error) {
	v = filepath.Join(d, t)
	if strings.HasPrefix(v, filepath.Clean(d)) {
		return v, nil
	}

	return "", fmt.Errorf("%s: %s", "content filepath is tainted", t)
}

func unzipFile(f *zip.File, destination string) error {
	filePath, err := sanitizeArchivePath(destination, f.Name)
	if err != nil {
		return fmt.Errorf("invalid file path: %s", err)
	}

	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filepath.Clean(filePath), os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory: %s", err)
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Errorf("cannot create directory: %s", err)
	}

	destinationFile, err := os.OpenFile(filepath.Clean(filePath), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return fmt.Errorf("cannot open file: %s", err)
	}

	defer destinationFile.Close()

	zippedFile, err := f.Open()
	if err != nil {
		return fmt.Errorf("cannot open zipped file: %s", err)
	}

	defer zippedFile.Close()

	for {
		_, err := io.CopyN(destinationFile, zippedFile, 1024)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	return nil
}
