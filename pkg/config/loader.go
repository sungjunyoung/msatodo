package config

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
)

type Loader interface {
	Load(config Config) error
}

type clientLoader struct {
	path string
}

func NewClientLoader(path string) *clientLoader {
	return &clientLoader{path: path}
}

func (cl clientLoader) Load(config Config) error {
	f, err := ioutil.ReadFile(cl.path)
	if err != nil {
		return fmt.Errorf("cannot load config from path %s", cl.path)
	}

	if err := yaml.Unmarshal(f, config); err != nil {
		return fmt.Errorf("cannot unmarshal config: +%v", err)
	}

	return nil
}
