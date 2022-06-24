# Minimal-douyin

Project documentation: https://ntp0va1sf4.feishu.cn/docx/doxcnGUuDyP5k19AWPsf9yPPe5d

## Introduction

Minimal-douyin service is divided into six main sections.

| Service Name    | Usage                    | Framework  | protocol | Path                        | IDL                               |
| --------------- | ------------------------ | ---------- | -------- | --------------------------- | --------------------------------- |
| apiservice      | http interface           | kitex/gin  | http     | minimal-douyin/cmd/api      |                                   |
| userservice     | user data management     | kitex/gorm | protobuf | minimal-douyin/cmd/user     | minimal-douyin/idl/user.proto     |
| videoservice    | video data management    | kitex/gorm | protobuf | minimal-douyin/cmd/video    | minimal-douyin/idl/video.proto    |
| favoriteservice | favorite data management | kitex/gorm | protobuf | minimal-douyin/cmd/favorite | minimal-douyin/idl/favorite.proto |
| commentservice  | comment data management  | kitex/gorm | protobuf | minimal-douyin/cmd/comment  | minimal-douyin/idl/comment.proto  |
| followservice   | follow data management   | kitex/gorm | protobuf | minimal-douyin/cmd/follow   | minimal-douyin/idl/follow.proto   |


### Use Basic Features

- Middleware、Rate Limiting、Request Retry、Connection Multiplexing
- Tracing
  - use jaeger to tracing
- Service Discovery and Register
  - use [registry-etcd](https://github.com/kitex-contrib/registry-etcd) to discovery and register service

### catalog introduce

| catalog        | introduce                |
| -------------- | ------------------------ |
| pkg/constants  | constant                 |
| pkg/bound      | customized bound handler |
| pkg/errno      | customized error number  |
| pkg/middleware | RPC middleware           |
| pkg/tracer     | init jaeger              |
| dal            | db operation             |
| pack           | data pack                |
| service        | business logic           |

## Quick Start

### 1.Setup Basic Dependence

```
docker-compose up
```

### 2.Run User RPC Server

```
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### 3.Run Video RPC Server

```
cd cmd/video
sh build.sh
sh output/bootstrap.sh
```

### 4.Run Favorite RPC Server

```
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh
```

### 5.Run Comment RPC Server

```
cd cmd/comment
sh build.sh
sh output/bootstrap.sh
```

### 6.Run Follow RPC Server

```
cd cmd/follow
sh build.sh
sh output/bootstrap.sh
```

### 7.Run API Server

```
cd cmd/api
chmod +x run.sh
./run.sh
```
