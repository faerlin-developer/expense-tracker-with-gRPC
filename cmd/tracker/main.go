package main

import (
	"context"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/internal/logger"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/internal/server"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/internal/store"
	"os"
)

const port = 5000

func main() {

	ctx := context.Background()

	// Store, Logger, and Server are interfaces, enabling easy switching between various implementations.
	var db store.Store
	var log logger.Logger
	var srv server.Server

	// Select and create concrete implementations of Store, Logger, and Server
	log = logger.NewStructuredLogger()
	db = store.NewInMemoryStore(log)
	srv = server.NewRpcServer(db, log)

	// Start the server
	log.Info(ctx, "Server started", "port", port)
	err := srv.ListenAndServe(ctx, 5000)
	if err != nil {
		log.Error(ctx, "failed to start server", err)
		os.Exit(1)
	}
}
