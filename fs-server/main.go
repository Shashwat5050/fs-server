package main

import (
	"context"
	"log"
	"os"

	"go.uber.org/zap"
	"iceline-hosting.com/core/backblaze"
	coreLogger "iceline-hosting.com/core/logger"
	v1 "iceline-hosting.com/fs-server/controller/grpc/v1"
	"iceline-hosting.com/fs-server/usecase"
)

func main() {
	// initialize zap logger
	logger, err := coreLogger.NewDefaultLogger()
	if err != nil {
		log.Fatalf("cannot initialize otel logger: %v", err)
	}
	logger.Debug("initialized logger")

	// initialize object storage client'
	obs, err := backblaze.NewClient()
	if err != nil {
		logger.Error("cannot initialize object storage client", zap.Error(err))
		panic(err)
	}
	logger.Debug("initialized object storage client")

	fusecase := usecase.NewLocalFileManager(
		logger,
		obs,
		usecase.WithBucket(os.Getenv("B2_BUCKET")),
		usecase.WithVolumeDir(os.Getenv("FS_VOLUME_DIR")),
		usecase.WithCleanPeriod(os.Getenv("FS_CLEAN_PERIOD")),
		usecase.WithKeepPeriod(os.Getenv("FS_KEEP_PERIOD")),
	)
	logger.Debug("initialized usecase layer")

	go fusecase.RunTrashCleaner(context.Background())

	c := v1.NewController(
		logger,
		fusecase,
		v1.WithPort(os.Getenv("FS_GRPC_PORT")),
	)
	defer c.Stop()

	logger.Debug("initialized grpc controller")

	if err := c.Run(); err != nil {
		logger.Error("cannot start grpc server", zap.Error(err))
	}
}
