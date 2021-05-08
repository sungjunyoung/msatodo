package grpc

import (
	"context"
	"github.com/araddon/dateparse"
	"github.com/sirupsen/logrus"
	"github.com/sungjunyoung/prototodo/pkg/manager/adding"
	"google.golang.org/grpc"
)

type Handler struct {
	*grpc.Server
}

func NewHandler(addingSvc adding.Service) *Handler {
	grpcServer := grpc.NewServer()
	RegisterJobServer(grpcServer, &jobServer{adding: addingSvc})

	return &Handler{grpcServer}
}

type jobServer struct {
	JobServer
	adding adding.Service
}

func (js *jobServer) AddJob(ctx context.Context, req *AddJobRequest) (*AddJobResponse, error) {
	doneAt, err := dateparse.ParseLocal(req.DoneAt)
	if err != nil {
		logrus.Errorf("cannot parse done time: %+v", err)
		return nil, err
	}

	job := adding.Job{
		Name:   req.Name,
		Email:  req.Email,
		DoneAt: doneAt,
	}

	id, err := js.adding.AddJob(job)
	if err != nil {
		return nil, err
	}

	return &AddJobResponse{Id: id}, nil
}
