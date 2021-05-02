package manager

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/sungjunyoung/prototodo/protos/v1/job"
)

type jobServer struct {
	job.JobServer
}

func newJobServer() *jobServer {
	return &jobServer{}
}

func (js *jobServer) AddJob(ctx context.Context, req *job.AddJobRequest) (*job.AddJobResponse, error) {
	logrus.Infof("adding job %+v", req)
	return &job.AddJobResponse{Id: 1}, nil
}
