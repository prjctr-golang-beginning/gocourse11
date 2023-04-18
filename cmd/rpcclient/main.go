package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	args := &Args{2, 3}
	var result int

	err = client.Call("Calculator.Add", args, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d + %d = %d\n", args.A, args.B, result)

	args = &Args{4, 2}

	err = client.Call("Calculator.Divide", args, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d / %d = %d\n", args.A, args.B, result)
}
