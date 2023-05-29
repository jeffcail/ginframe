package test

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
)

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func TestRpc(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9094")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := ArithRequest{9, 2}
	var res ArithResponse

	err = client.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Println(res)
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
}
