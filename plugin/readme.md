## 构建插件


### 放到 go/bin目录中
```bash
go build -o $GOPATH/bin/protoc-gen-go-netrpc .
ls -al $GOPATH/bin

# 使用时无需指定plugin位置
protoc --go-netrpc_out=plugins=netrpc:. hello.proto
```

### 放到当前目录中
```bash
go build -o protoc-gen-go-netrpc .
ls -al .

# 使用时需要指定plugin的位置
protoc --go-netrpc_out=plugins=netrpc:. --plugin=../plugin/protoc-gen-go-netrpc hello.proto
```


## 注意事项

	"github.com/golang/protobuf/protoc-gen-go/generator" 会提示一下问题
	//WARNING: Package "github.com/golang/protobuf/protoc-gen-go/generator" is deprecated.
	//A future release of golang/protobuf will delete this package,
	//which has long been excluded from the compatibility promise.