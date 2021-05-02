package config

type Type string

const (
	CLIENT Type = "client"
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
