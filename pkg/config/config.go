package config

type Type string

const (
	CLIENT  Type = "client"
	MANAGER Type = "manager"
)

type Config interface {
	Type() Type
}

type Client struct {
	Email           string `yaml:"email"`
	ManagerEndpoint string `yaml:"managerEndpoint"`
}

func (c Client) Type() Type {
	return CLIENT
}

type Manager struct {
	Port          string `yaml:"port"`
	StoreEndpoint string `yaml:"storeEndpoint"`
}

func (m Manager) Type() Type {
	return MANAGER
}
