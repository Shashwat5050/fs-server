package actclient

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "iceline-hosting.com/backend/proto/activity"
	"iceline-hosting.com/core/logger"
	"iceline-hosting.com/core/models"
)

type client struct {
	log logger.Logger
	c   pb.ActivityManagerClient
}

// Create a grpc client based on the proto file /proto/activity.proto
// The client will be used to call the grpc server
func NewClient(log logger.Logger, address string) (client, error) {
	// Create a connection to the grpc server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return client{}, err
	}

	// Create a new client based on the connection
	c := pb.NewActivityManagerClient(conn)

	return client{c: c, log: log}, nil
}

func (c client) RegisterActivity(ctx context.Context, activity any) error {
	switch activity.(type) {
	case models.UserActivity:
		uc := activity.(models.UserActivity)

		_, err := c.c.RegisterUserActivity(ctx, &pb.RegisterUserActivityRequest{
			UserId:       uc.UserID,
			ActivityType: uc.ActivityType,
			LogTime:      timestamppb.New(uc.Time),
		})
		if err != nil {
			return err
		}
	case models.GsActivity:
		gc := activity.(models.GsActivity)

		_, err := c.c.RegisterGsActivity(ctx, &pb.RegisterGsActivityRequest{
			UserId:       gc.UserID,
			ActivityType: gc.ActivityType,
			ServerName:   gc.ServerName,
			LogTime:      timestamppb.New(gc.Time),
		})
		if err != nil {
			return err
		}
	default:
		return errors.New("Invalid activity type")
	}

	return nil
}
