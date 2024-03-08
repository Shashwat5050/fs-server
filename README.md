
```
fs-server
├─ Dockerfile-fs-server
├─ core
│  ├─ activity-grpc-client
│  │  └─ client.go
│  ├─ authorization
│  │  ├─ authorization.go
│  │  └─ core-authorizer.go
│  ├─ backblaze
│  │  └─ client.go
│  ├─ basic_model.csv
│  ├─ config
│  │  └─ config.go
│  ├─ cronmanager
│  │  └─ cronmanager.go
│  ├─ curse-forge
│  │  └─ client.go
│  ├─ dbhelper
│  │  └─ authorization.go
│  ├─ email
│  │  └─ client.go
│  ├─ error
│  │  └─ errors.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ infrastructure
│  │  ├─ .probe-manager.go.swo
│  │  └─ probe-manager.go
│  ├─ jwthelper
│  │  └─ jwthelper.go
│  ├─ logger
│  │  └─ logger.go
│  ├─ models
│  │  ├─ activity.go
│  │  ├─ backups.go
│  │  ├─ games.go
│  │  ├─ gs-info.go
│  │  ├─ mods-models.go
│  │  ├─ nodes.go
│  │  ├─ notifications.go
│  │  ├─ schedules.go
│  │  ├─ settings-activities.go
│  │  ├─ sftp.go
│  │  ├─ tasks.go
│  │  ├─ user_role.go
│  │  └─ users.go
│  ├─ mods
│  │  └─ .modes.go.swp
│  ├─ nomad
│  │  └─ nomad.go
│  ├─ notify
│  │  └─ client.go
│  ├─ postgres
│  │  └─ postgres.go
│  ├─ redis-client
│  │  └─ rclient.go
│  ├─ types
│  │  └─ types.go
│  └─ utils
│     └─ gin.go
├─ fs-server
│  ├─ .env
│  ├─ controller
│  │  └─ grpc
│  │     └─ v1
│  │        ├─ backup-manager.go
│  │        ├─ controller.go
│  │        ├─ file-manager.go
│  │        └─ sftp-manager.go
│  ├─ example.env
│  ├─ go.mod
│  ├─ go.sum
│  ├─ main.go
│  └─ usecase
│     ├─ backup-manager.go
│     ├─ file-manager.go
│     ├─ file-manager_test.go
│     ├─ interfaces.go
│     ├─ sftp-manager.go
│     └─ utils.go
└─ proto
   ├─ activity
   │  ├─ activity.pb.go
   │  └─ activity_grpc.pb.go
   ├─ activity.proto
   ├─ authentication
   │  ├─ authentication.pb.go
   │  └─ authentication_grpc.pb.go
   ├─ authentication.proto
   ├─ authorization
   │  ├─ authorization.pb.go
   │  └─ authorization_grpc.pb.go
   ├─ authorization.proto
   ├─ backup-manager.proto
   ├─ backupmanager
   │  ├─ backup-manager.pb.go
   │  └─ backup-manager_grpc.pb.go
   ├─ fs-manager.proto
   ├─ fsmanager
   │  ├─ fs-manager.pb.go
   │  └─ fs-manager_grpc.pb.go
   ├─ go.mod
   ├─ go.sum
   ├─ jwt
   │  └─ jwt.pb.go
   ├─ jwt.proto
   ├─ proto.sh
   ├─ sftp-manager
   │  ├─ sftp-manager.pb.go
   │  └─ sftp-manager_grpc.pb.go
   └─ sftp-manager.proto

```
```
fs-server
├─ Dockerfile-fs-server
├─ README.md
├─ core
│  ├─ activity-grpc-client
│  │  └─ client.go
│  ├─ authorization
│  │  ├─ authorization.go
│  │  └─ core-authorizer.go
│  ├─ backblaze
│  │  └─ client.go
│  ├─ basic_model.csv
│  ├─ config
│  │  └─ config.go
│  ├─ cronmanager
│  │  └─ cronmanager.go
│  ├─ curse-forge
│  │  └─ client.go
│  ├─ dbhelper
│  │  └─ authorization.go
│  ├─ email
│  │  └─ client.go
│  ├─ error
│  │  └─ errors.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ infrastructure
│  │  ├─ .probe-manager.go.swo
│  │  └─ probe-manager.go
│  ├─ jwthelper
│  │  └─ jwthelper.go
│  ├─ logger
│  │  └─ logger.go
│  ├─ models
│  │  ├─ activity.go
│  │  ├─ backups.go
│  │  ├─ games.go
│  │  ├─ gs-info.go
│  │  ├─ mods-models.go
│  │  ├─ nodes.go
│  │  ├─ notifications.go
│  │  ├─ schedules.go
│  │  ├─ settings-activities.go
│  │  ├─ sftp.go
│  │  ├─ tasks.go
│  │  ├─ user_role.go
│  │  └─ users.go
│  ├─ mods
│  │  └─ .modes.go.swp
│  ├─ nomad
│  │  └─ nomad.go
│  ├─ notify
│  │  └─ client.go
│  ├─ postgres
│  │  └─ postgres.go
│  ├─ redis-client
│  │  └─ rclient.go
│  ├─ types
│  │  └─ types.go
│  └─ utils
│     └─ gin.go
├─ fs-server
│  ├─ .env
│  ├─ controller
│  │  └─ grpc
│  │     └─ v1
│  │        ├─ backup-manager.go
│  │        ├─ controller.go
│  │        ├─ file-manager.go
│  │        └─ sftp-manager.go
│  ├─ example.env
│  ├─ go.mod
│  ├─ go.sum
│  ├─ main.go
│  └─ usecase
│     ├─ backup-manager.go
│     ├─ file-manager.go
│     ├─ file-manager_test.go
│     ├─ interfaces.go
│     ├─ sftp-manager.go
│     └─ utils.go
└─ proto
   ├─ activity
   │  ├─ activity.pb.go
   │  └─ activity_grpc.pb.go
   ├─ activity.proto
   ├─ authentication
   │  ├─ authentication.pb.go
   │  └─ authentication_grpc.pb.go
   ├─ authentication.proto
   ├─ authorization
   │  ├─ authorization.pb.go
   │  └─ authorization_grpc.pb.go
   ├─ authorization.proto
   ├─ backup-manager.proto
   ├─ backupmanager
   │  ├─ backup-manager.pb.go
   │  └─ backup-manager_grpc.pb.go
   ├─ fs-manager.proto
   ├─ fsmanager
   │  ├─ fs-manager.pb.go
   │  └─ fs-manager_grpc.pb.go
   ├─ go.mod
   ├─ go.sum
   ├─ jwt
   │  └─ jwt.pb.go
   ├─ jwt.proto
   ├─ proto.sh
   ├─ sftp-manager
   │  ├─ sftp-manager.pb.go
   │  └─ sftp-manager_grpc.pb.go
   └─ sftp-manager.proto

```
```
fs-server
├─ Dockerfile-fs-server
├─ README.md
├─ core
│  ├─ activity-grpc-client
│  │  └─ client.go
│  ├─ authorization
│  │  ├─ authorization.go
│  │  └─ core-authorizer.go
│  ├─ backblaze
│  │  └─ client.go
│  ├─ basic_model.csv
│  ├─ config
│  │  └─ config.go
│  ├─ cronmanager
│  │  └─ cronmanager.go
│  ├─ curse-forge
│  │  └─ client.go
│  ├─ dbhelper
│  │  └─ authorization.go
│  ├─ email
│  │  └─ client.go
│  ├─ error
│  │  └─ errors.go
│  ├─ go.mod
│  ├─ go.sum
│  ├─ infrastructure
│  │  ├─ .probe-manager.go.swo
│  │  └─ probe-manager.go
│  ├─ jwthelper
│  │  └─ jwthelper.go
│  ├─ logger
│  │  └─ logger.go
│  ├─ models
│  │  ├─ activity.go
│  │  ├─ backups.go
│  │  ├─ games.go
│  │  ├─ gs-info.go
│  │  ├─ mods-models.go
│  │  ├─ nodes.go
│  │  ├─ notifications.go
│  │  ├─ schedules.go
│  │  ├─ settings-activities.go
│  │  ├─ sftp.go
│  │  ├─ tasks.go
│  │  ├─ user_role.go
│  │  └─ users.go
│  ├─ mods
│  │  └─ .modes.go.swp
│  ├─ nomad
│  │  └─ nomad.go
│  ├─ notify
│  │  └─ client.go
│  ├─ postgres
│  │  └─ postgres.go
│  ├─ redis-client
│  │  └─ rclient.go
│  ├─ types
│  │  └─ types.go
│  └─ utils
│     └─ gin.go
├─ docker-compose.yml
├─ fs-server
│  ├─ .env
│  ├─ controller
│  │  └─ grpc
│  │     └─ v1
│  │        ├─ backup-manager.go
│  │        ├─ controller.go
│  │        ├─ file-manager.go
│  │        └─ sftp-manager.go
│  ├─ example.env
│  ├─ go.mod
│  ├─ go.sum
│  ├─ main.go
│  └─ usecase
│     ├─ backup-manager.go
│     ├─ file-manager.go
│     ├─ file-manager_test.go
│     ├─ interfaces.go
│     ├─ sftp-manager.go
│     └─ utils.go
├─ fs-server-job.hcl
├─ fs-server.tar
└─ proto
   ├─ activity
   │  ├─ activity.pb.go
   │  └─ activity_grpc.pb.go
   ├─ activity.proto
   ├─ authentication
   │  ├─ authentication.pb.go
   │  └─ authentication_grpc.pb.go
   ├─ authentication.proto
   ├─ authorization
   │  ├─ authorization.pb.go
   │  └─ authorization_grpc.pb.go
   ├─ authorization.proto
   ├─ backup-manager.proto
   ├─ backupmanager
   │  ├─ backup-manager.pb.go
   │  └─ backup-manager_grpc.pb.go
   ├─ fs-manager.proto
   ├─ fsmanager
   │  ├─ fs-manager.pb.go
   │  └─ fs-manager_grpc.pb.go
   ├─ go.mod
   ├─ go.sum
   ├─ jwt
   │  └─ jwt.pb.go
   ├─ jwt.proto
   ├─ proto.sh
   ├─ sftp-manager
   │  ├─ sftp-manager.pb.go
   │  └─ sftp-manager_grpc.pb.go
   └─ sftp-manager.proto

```