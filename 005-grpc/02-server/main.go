package main

import (
	"context"
	"fmt"
	"gomodules/005-grpc/02-server/echo"
	"net"

	"google.golang.org/grpc"
)

type EchoServer struct {
	echo.UnimplementedEchoServerServer
}

func (e *EchoServer) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{
		Response: "Echo: " + req.Message,
	}, nil
}

func main() {
	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	srv := &EchoServer{}
	echo.RegisterEchoServerServer(s, srv)

	fmt.Println("Listen and serve at port 8080")
	err = s.Serve(lst)
	if err != nil {
		panic(err)
	}
}
