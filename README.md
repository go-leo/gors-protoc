# gors-protoc
gors 的 protobuffer 插件
- protoc-gen-go-gors 生成 gin 代码
- protoc-go-inject-tag 修改 pb 结构体中的标签


# quick start

1. 安装
```
go install ./cmd/protoc-gen-go-gors
go install ./cmd/protoc-go-inject-tag
```

2. 运行exmaple
```
cd ./example/gors/api
make api

```