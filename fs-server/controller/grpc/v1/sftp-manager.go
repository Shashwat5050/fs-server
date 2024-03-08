package v1

import (
	"context"

	"go.uber.org/zap"
	sfp "iceline-hosting.com/backend/proto/sftp-manager"
)


func(c *controller)ConnectToSFTP(ctx context.Context,req *sfp.SFTPConnectRequest)(*sfp.SFTPConnectResponse,error){
	c.log.Info("ConnectToSFTP",zap.String("user", req.Username))

	msg,err:=c.use.ConnectToSFTP(ctx,req.Host,req.Password,req.Port,req.Username)
	if err!=nil {
		return nil,err
	}


	return &sfp.SFTPConnectResponse{
		Success: true,
		ErrorMessage: msg,
	},nil
}