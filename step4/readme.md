## 基于protobuf的的rpc


### 使用命令辅助测试

```bash
```

### protobuf


```bash

# 生成仅有message的proto
protoc --go_out=. hello.proto

# 生成带有rpc的proto
# 教程使用protoc --go_out=plugins=grpc:. hello.proto但无效
protoc --go_out=. --go-grpc_out=. hello.proto

# 参考 plugin目录
protoc --go-netrpc_out=plugins=netrpc:. --plugin=/Users/qingliu/study/gorpcdemo/plugin/protoc-gen-go-netrpc hello.proto
```