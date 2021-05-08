package manager

import (
	"github.com/sirupsen/logrus"
	"github.com/sungjunyoung/prototodo/pkg/config"
	"github.com/sungjunyoung/prototodo/pkg/manager/adding"
	"github.com/sungjunyoung/prototodo/pkg/manager/http/grpc"
	"net"
)

type Manager struct {
	config *config.Manager
	grpc   *grpc.Handler
}

func NewManager(loader config.Loader, addingSvc adding.Service) (Manager, error) {
	mgr := Manager{}

	cfg := &config.Manager{}
	if err := loader.Load(cfg); err != nil {
		return mgr, err
	}

	return Manager{
		config: cfg,
		grpc:   grpc.NewHandler(addingSvc),
	}, nil
}

func (m Manager) ServeGrpc(errCh chan error) {
	logrus.Infof("starting manager in port %s for grpc", m.config.Port)
	listener, err := net.Listen("tcp", ":"+m.config.Port)
	if err != nil {
		errCh <- err
	}

	err = m.grpc.Serve(listener)
	if err != nil {
		errCh <- err
	}
}
