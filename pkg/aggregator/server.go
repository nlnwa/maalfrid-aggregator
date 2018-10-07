package aggregator

import (
	"fmt"
	"net"
	"os"

	"github.com/nlnwa/pkg/log"
	"google.golang.org/grpc"

	api "github.com/nlnwa/maalfrid-aggregator/maalfrid/aggregator"
)

type Server struct {
	logger log.Logger
	api    *AggregatorApi
	grpcOpts []grpc.ServerOption
}

func WithLogger(logger *log.Logger) func (a *Server) error {
	return func (a *Server) error {
		a.logger = *logger
		return nil
	}
}

func WithApi(api *AggregatorApi) func (a *Server) error {
	return func (a *Server) error {
		a.api = api
		return nil
	}
}

func (s *Server) SetOption(options ...func(*Server) error) error {
	for _, opt := range options {
		if err := opt(s); err != nil {
			return err
		}
	}
	return nil
}

func ServeApi(port int, options ...func(*Server) error) error {
	s := Server{}
	s.SetOption(options...)
	log := s.logger

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Error(err.Error(), "port", port)
		os.Exit(1)
	} else {
		log.Info("API server listening", "port", port)
	}

	server := grpc.NewServer(s.grpcOpts...)
	api.RegisterAggregatorServer(server, s.api)

	return server.Serve(listener)
}
