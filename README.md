# 简介
基于 `protobuf` 的API工程化demo，其中包含两个自定义插件用于 http 代码转化生成

protobuffer 插件
- protoc-gen-go-gors 生成 gin 代码
- protoc-go-inject-tag 修改 pb 结构体中的标签


# quick start

1. 安装
```
# 生成 pb.go
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
# 生成 grpc.pb.go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
# 生成 gin http
go install ./cmd/protoc-gen-go-gorsp
# 补充 字段 tag
go install ./cmd/protoc-go-inject-tagp
# 生成 validate 代码
go install github.com/envoyproxy/protoc-gen-validate@v0.6.7
# 生成 swagger
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

```

1. 运行exmaple
```
cd ./example/gors/api
make api

```
