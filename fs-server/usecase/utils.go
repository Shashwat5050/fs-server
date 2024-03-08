package usecase

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
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
	filePath := filepath.Join(fm.VolumeDir, newPath, newName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
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
