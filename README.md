### 自动化部署平台

#### 目录结构
```
- api/
- bin/
- main/
- utils
- pkg
- protobuf
- db
- model
```

#### 生成pb文件
```shell
protoc -I./protobuf --go_out=./protobuf/pb --go_opt=paths=source_relative --go-grpc_out=./protobuf/pb/ --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false protobuf/*.proto 
```

#### 打包
```shell
go build -o bin/cicd main.go
```