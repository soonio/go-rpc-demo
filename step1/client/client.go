package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	var err error
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "Hari", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
