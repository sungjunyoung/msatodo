package manager

import (
	"github.com/sirupsen/logrus"
	"github.com/sungjunyoung/prototodo/pkg/config"
	"net"
)

type Manager interface {
	Start() error
}

type manager struct {
	config config.Manager
	server *server
}

func NewManager(loader config.Loader) (Manager, error) {
	c := &config.Manager{}
	if err := loader.Load(c); err != nil {
		return nil, err
	}

	return &manager{
		config: *c,
		server: newServer(),
	}, nil
}

func (m *manager) Start() error {
	logrus.Infof("starting manager in port %s", m.config.Port)
	listener, err := net.Listen("tcp", ":"+m.config.Port)
	if err != nil {
		return err
	}

	return m.server.Serve(listener)
}
