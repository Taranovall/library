package main

import (
	"jwt-service/grpc"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Start the gRPC server
	go grpc.GRPCListen()

	// Wait for termination signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	// Perform cleanup tasks here, if necessary

	// Exit
	os.Exit(0)
}
