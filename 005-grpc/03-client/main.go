package main

import (
	"context"
	"fmt"
	"gomodules/005-grpc/03-client/echo"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	e := echo.NewEchoServerClient(conn)
	resp, err := e.Echo(ctx, &echo.EchoRequest{
		Message: "Hello Fu!",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Got from server:", resp.Response)

}
