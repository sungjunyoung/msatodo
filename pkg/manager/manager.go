package manager

import (
	"github.com/sirupsen/logrus"
	"github.com/sungjunyoung/msatodo/pkg/config"
	"github.com/sungjunyoung/msatodo/pkg/manager/adding"
	"github.com/sungjunyoung/msatodo/pkg/manager/http/grpc"
	"github.com/sungjunyoung/msatodo/pkg/manager/listing"
	"net"
)

type Manager struct {
	config *config.Manager
	grpc   *grpc.Handler
}

func NewManager(
	loader config.Loader,
	addingSvc adding.Service,
	listingSvc listing.Service) (Manager, error) {
	mgr := Manager{}

	cfg := &config.Manager{}
	if err := loader.Load(cfg); err != nil {
		return mgr, err
	}

	return Manager{
		config: cfg,
		grpc:   grpc.NewHandler(addingSvc, listingSvc),
	}, nil
}

func (m Manager) ServeGrpc(errCh chan error) {
	logrus.Infof("starting manager in port %s for grpc", m.config.GrpcPort)
	listener, err := net.Listen("tcp", ":"+m.config.GrpcPort)
	if err != nil {
		errCh <- err
	}

	err = m.grpc.Serve(listener)
	if err != nil {
		errCh <- err
	}
}
