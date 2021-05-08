package grpc

import (
	"context"
	"github.com/araddon/dateparse"
	"github.com/sirupsen/logrus"
	"github.com/sungjunyoung/prototodo/pkg/manager/adding"
	"github.com/sungjunyoung/prototodo/pkg/manager/listing"
	"google.golang.org/grpc"
)

type Handler struct {
	*grpc.Server
}

func NewHandler(addingSvc adding.Service, listingSvc listing.Service) *Handler {
	grpcServer := grpc.NewServer()
	RegisterJobServer(
		grpcServer,
		&jobServer{adding: addingSvc, listing: listingSvc},
	)

	return &Handler{grpcServer}
}

type jobServer struct {
	JobServer
	adding  adding.Service
	listing listing.Service
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

func (js *jobServer) ListJobs(ctx context.Context, req *ListJobsRequest) (*ListJobsResponse, error) {
	jobs, err := js.listing.ListJobs()
	if err != nil {
		return nil, err
	}

	jobMessages := make([]*JobMessage, len(jobs))
	for i, job := range jobs {
		jobMessages[i] = &JobMessage{
			Id:        job.ID,
			Name:      job.Name,
			Email:     job.Email,
			CreatedAt: job.CreatedAt.String(),
			DoneAt:    job.DoneAt.String(),
		}
	}

	return &ListJobsResponse{
		Jobs: jobMessages,
	}, nil
}
