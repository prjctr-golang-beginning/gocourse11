package main

import (
	"errors"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Calculator int

type Args struct {
	A, B int
}

func (c *Calculator) Add(args *Args, result *int) error {
	*result = args.A + args.B
	return nil
}

func (c *Calculator) Subtract(args *Args, result *int) error {
	*result = args.A - args.B
	return nil
}

func (c *Calculator) Multiply(args *Args, result *int) error {
	*result = args.A * args.B
	return nil
}

func (c *Calculator) Divide(args *Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}

func main() {
	calculator := new(Calculator)
	server := rpc.NewServer()
	server.Register(calculator)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
