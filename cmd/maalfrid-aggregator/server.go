package main

import (
	"google.golang.org/grpc"
	"net"
	"fmt"
	"os"
	"github.com/nlnwa/maalfrid-aggregator/pkg/aggregator"
)

type Server struct {
	port int
}

func NewServer(api aggregator.Api, options ...func(*Server) error) error {
	var logger = logger
	var grpcOpts []grpc.ServerOption

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error(err.Error(), "port", port)
		os.Exit(1)
	} else {
		logger.Info("API server listening", "port", port)
	}
	server := grpc.NewServer(grpcOpts...)

	api.RegisterAggregatorServer(server, aggregator.NewApi(options))

	return server.Serve(listener)
}

