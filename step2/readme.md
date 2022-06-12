## gob | json


### 使用命令辅助测试

```bash
# 使用nc命令建立临时tcp服务监听1234端口
nc -l 1234
# 使用`nc 主机名称 端口` 发送消息
echo -e '{"method":"HelloService.Hello","params":["Hari"],"id":0}' | nc localhost 1234
```