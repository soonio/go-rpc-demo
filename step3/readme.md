## 支持http的rpc

### 使用命令辅助测试

```bash
curl localhost:1234/hello -X POST --data '{"method":"HelloService.Hello","params":["Harri"],"id":10086}'
```