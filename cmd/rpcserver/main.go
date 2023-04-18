package main

import (
	"errors"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	log "rpc"
)

type Calculator int

type Args struct {
	A, B int
}

func (c *Calculator) Add(args *Args, result *int) error {
	*result = args.A + args.B
	log.Info(`Add operation`)
	return nil
}

func (c *Calculator) Subtract(args *Args, result *int) error {
	*result = args.A - args.B
	log.Info(`Subtract operation`)
	return nil
}

func (c *Calculator) Multiply(args *Args, result *int) error {
	*result = args.A * args.B
	log.Info(`Multiply operation`)
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

	log.Info(`Server started`)
	for {
		conn, err := listener.Accept()
		log.Info(`Connection accepted`)
		if err != nil {
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

	// exercises
	// 1. Створити клієнт і сервер, використовуючи json-rpc:
	// Сервер виконує роль кошика інтернет-магазина. Кошик очікує id і name товара. Додає до кошику, оновлює і видаляє.
	// Клієнт виконує роль клієнта і працює із кошиком: додате, видаляє і редагує товари
}
