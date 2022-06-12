package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	"rpc-hello/step3"
)

type HelloService struct{}

func (i *HelloService) Hello(request string, reply *string) error {
	*reply = fmt.Sprintf("[%s] Hello %s! \nby http serve.", time.Now().Format("2006-01-02 15:04:05"), request)
	return nil
}

var _ step3.HelloServiceInterface = (*HelloService)(nil)

func main() {

	var err error
	err = rpc.RegisterName(step3.HelloServiceName, new(HelloService))
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				panic(err)
			}
			return
		}
	})

	fmt.Println("http serve at: 127.0.0.1:1234")
	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}
