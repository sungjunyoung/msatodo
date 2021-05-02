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

func NewClientLoader(path string) Loader {
	return &clientLoader{path: path}
}

func (cl clientLoader) Load(config Config) error {
	return load(cl.path, config)
}

type managerLoader struct {
	path string
}

func NewManagerLoader(path string) Loader {
	return &managerLoader{path: path}
}

func (ml managerLoader) Load(config Config) error {
	return load(ml.path, config)
}

func load(path string, config Config) error {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot load config from path %s", path)
	}

	if err := yaml.Unmarshal(f, config); err != nil {
		return fmt.Errorf("cannot unmarshal config: +%v", err)
	}

	return nil
}
