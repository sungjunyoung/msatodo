package client

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/sungjunyoung/prototodo/pkg/config"
	"github.com/sungjunyoung/prototodo/protos/v1/job"
	"google.golang.org/grpc"
	"time"
)

type Client interface {
	AddJob(name string, doneAt string) (*job.AddJobResponse, error)
}

type client struct {
	config *config.Client
	job    job.JobClient
}

func NewClient(loader config.Loader) (Client, error) {
	c := &config.Client{}
	if err := loader.Load(c); err != nil {
		return nil, err
	}

	return &client{config: c}, nil
}

func (c *client) AddJob(name string, doneAt string) (*job.AddJobResponse, error) {
	conn, err := c.getConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	jobClient := job.NewJobClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := jobClient.AddJob(ctx, &job.AddJobRequest{
		Job:    name,
		Email:  c.config.Email,
		DoneAt: doneAt,
	})
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *client) getConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(c.config.ManagerEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Errorf("cannot connect: %v", err)
		return nil, err
	}

	return conn, nil
}
