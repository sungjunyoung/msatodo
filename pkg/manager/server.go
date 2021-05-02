package manager

import (
	"github.com/sungjunyoung/prototodo/protos/v1/job"
	"google.golang.org/grpc"
)

type server struct {
	*grpc.Server
}

func newServer() *server {
	grpcServer := grpc.NewServer()
	job.RegisterJobServer(grpcServer, newJobServer())

	return &server{grpcServer}
}
