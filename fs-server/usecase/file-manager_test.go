package usecase_test

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"iceline-hosting.com/core/logger"
	filemanager "iceline-hosting.com/fs-server/usecase"
)

func TestRunCleanNoTrash(t *testing.T) {
	t.Log("TestRunCleanNoTrash")

	log, err := logger.NewDefaultLogger()
	require.NoError(t, err)

	tmpVol, err := os.MkdirTemp("/tmp", "test")
	require.NoError(t, err)

	fm := filemanager.NewLocalFileManager(
		log,
		nil,
		filemanager.WithVolumeDir(tmpVol),
		filemanager.WithCleanPeriod("1s"),
		filemanager.WithKeepPeriod("1s"),
	)

	go fm.RunTrashCleaner(context.Background())

	// generate random file name
	randomName := strconv.Itoa(int(time.Now().UnixNano())) + ".tmp"

	fd, err := fm.CreateFile("/", randomName)
	require.NoError(t, err)

	n, err := fd.Write([]byte("test"))
	require.NoError(t, err)
	require.Equal(t, 4, n)

	err = fd.Close()
	require.NoError(t, err)

	time.Sleep(time.Second)

	_, err = fm.GetFileStat(".trash-" + randomName)
	require.ErrorIs(t, err, os.ErrNotExist)

	_, err = fm.GetFileStat(randomName)
	require.NoError(t, err)
}

func TestRunCleanInTime(t *testing.T) {
	t.Log("TestRunCleanInTime")

	log, err := logger.NewDefaultLogger()
	require.NoError(t, err)

	tmpVol, err := os.MkdirTemp("/tmp", "test")
	require.NoError(t, err)

	fm := filemanager.NewLocalFileManager(
		log,
		nil,
		filemanager.WithVolumeDir(tmpVol),
		filemanager.WithCleanPeriod("1s"),
		filemanager.WithKeepPeriod("1s"),
	)

	go fm.RunTrashCleaner(context.Background())

	// generate random file name
	randomName := strconv.Itoa(int(time.Now().UnixNano())) + ".tmp"

	fd, err := fm.CreateFile("/", randomName)
	require.NoError(t, err)

	n, err := fd.Write([]byte("test"))
	require.NoError(t, err)
	require.Equal(t, 4, n)

	err = fd.Close()
	require.NoError(t, err)

	err = fm.DeleteFile("/", randomName)
	require.NoError(t, err)

	time.Sleep(time.Second * 2)

	_, err = fm.GetFileStat(".trash-" + randomName)
	require.ErrorIs(t, err, os.ErrNotExist)
}

func TestRunCleanStillPresent(t *testing.T) {
	t.Log("TestRunCleanStillPresent")

	log, err := logger.NewDefaultLogger()
	require.NoError(t, err)

	tmpVol, err := os.MkdirTemp("/tmp", "test")
	require.NoError(t, err)

	fm := filemanager.NewLocalFileManager(
		log,
		nil,
		filemanager.WithVolumeDir(tmpVol),
		filemanager.WithCleanPeriod("1s"),
		filemanager.WithKeepPeriod("1s"),
	)

	go fm.RunTrashCleaner(context.Background())

	// generate random file name
	randomName := strconv.Itoa(int(time.Now().UnixNano())) + ".tmp"

	fd, err := fm.CreateFile("/", randomName)
	require.NoError(t, err)

	n, err := fd.Write([]byte("test"))
	require.NoError(t, err)
	require.Equal(t, 4, n)

	err = fd.Close()
	require.NoError(t, err)

	err = fm.DeleteFile("/", randomName)
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 200)

	_, err = fm.GetFileStat(".trash-" + randomName)
	require.NoError(t, err)

	time.Sleep(time.Second)

	_, err = fm.GetFileStat(".trash-" + randomName)
	require.ErrorIs(t, err, os.ErrNotExist)
}

func TestRunCleanerRecoverBeforeDelete(t *testing.T) {
	t.Log("TestRunCleanerRecoverBeforeDelete")

	log, err := logger.NewDefaultLogger()
	require.NoError(t, err)

	tmpVol, err := os.MkdirTemp("/tmp", "test")
	require.NoError(t, err)

	fm := filemanager.NewLocalFileManager(
		log,
		nil,
		filemanager.WithVolumeDir(tmpVol),
		filemanager.WithCleanPeriod("1s"),
		filemanager.WithKeepPeriod("1s"),
	)

	go fm.RunTrashCleaner(context.Background())

	// generate random file name
	randomName := strconv.Itoa(int(time.Now().UnixNano())) + ".tmp"

	fd, err := fm.CreateFile("/", randomName)
	require.NoError(t, err)

	n, err := fd.Write([]byte("test"))
	require.NoError(t, err)
	require.Equal(t, 4, n)

	err = fd.Close()
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 200)

	_, err = fm.GetFileStat(randomName)
	require.NoError(t, err)

	err = fm.DeleteFile("/", randomName)
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 200)

	_, err = fm.GetFileStat(".trash-" + randomName)
	require.NoError(t, err)

	err = fm.RecoverFile("/", randomName)
	require.NoError(t, err)

	time.Sleep(time.Second)

	_, err = fm.GetFileStat(".trash-" + randomName)
	require.ErrorIs(t, err, os.ErrNotExist)

	_, err = fm.GetFileStat(randomName)
	require.NoError(t, err)
}
