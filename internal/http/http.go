package http

import (
	"log"
	"net"

	piper "github.com/dreamstax/go-genproto/kaiapis/piper/v1alpha1"
	"google.golang.org/grpc"
)

type Server struct {
	*piper.UnimplementedPiperServer
	Port string
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.Port)
	if err != nil {
		return err
	}
	defer func() {
		err := lis.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	grpc := grpc.NewServer()
	piper.RegisterPiperServer(grpc, s)
	log.Printf("grpc server listening on port: %s", s.Port)
	return grpc.Serve(lis)
}
