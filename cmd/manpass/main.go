package main

import (
	"fmt"
	"github.com/vusalalishov/manpass/internal/server"
	"net"
)

func main() {
	srv := server.InjectGrpcServer()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5051))
	if err != nil {
		panic(err)
	}
	srv.Serve(lis)
}