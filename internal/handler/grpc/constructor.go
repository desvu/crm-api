package grpc

import (
	"net"

	"github.com/qilin/crm-api/internal/handler/grpc/game"
	"github.com/qilin/crm-api/pkg/grpc/proto"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	listener *net.Listener
}

type Params struct {
	fx.In

	GameHandler *game.Handler
}

func New(p Params) (*Server, error) {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()
	proto.RegisterGameServer(server, p.GameHandler)

	return &Server{
		Server:   server,
		listener: &listener,
	}, nil
}

func (s *Server) Start() {
	go func() {
		if err := s.Serve(*s.listener); err != nil {
			panic("")
		}
	}()
}
