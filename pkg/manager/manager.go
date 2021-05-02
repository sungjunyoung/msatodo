package manager

import (
	"github.com/sirupsen/logrus"
	"net"
)

type Manager interface {
	Start() error
}

type manager struct {
	server *server
	port   string
}

func NewManager() Manager {
	return &manager{
		server: newServer(),
		port:   "2222",
	}
}

func (m *manager) Start() error {
	logrus.Infof("starting manager in port %s", m.port)
	listener, err := net.Listen("tcp", ":"+m.port)
	if err != nil {
		return err
	}

	return m.server.Serve(listener)
}
