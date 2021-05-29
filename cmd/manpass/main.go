package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/vusalalishov/manpass/internal/db"
	"github.com/vusalalishov/manpass/internal/server"
	"net"
)

func main() {

	migrator, err := db.InjectMigrator()
	if err != nil {
		panic(err)
	}

	err = migrator.Migrate()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	srv, err := server.InjectGrpcServer()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5051))
	if err != nil {
		panic(err)
	}
	err = srv.Serve(lis)
	if err != nil {
		panic(err)
	}
}