package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"sync"
)

type Route struct {
	Name       string `yaml:"name"`
	Url        string `yaml:"url"`
	Method     string `yaml:"method"`
	HttpMethod string `yaml:"httpMethod"`
}
type Config struct {
	Routes []Route `yaml:"routes"`
}

var synchro sync.Once
var config *Config

const (
	CONFIG_FILE = CONFIG_DIR + "/routes.yml"
	CONFIG_DIR  = "web/config"
)

func NewConfig() *Config {
	synchro.Do(func() {
		config = &Config{
			Routes: config.GetRoutes()}

	})

	return config
}

func (config *Config) GetRoutes() []Route {
	filename, _ := filepath.Abs(CONFIG_FILE)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	configRoutes := make(map[string][]Route)
	err = yaml.Unmarshal(yamlFile, &configRoutes)
	if err != nil {
		panic(err)
	}
	routeList, errRead := configRoutes["routes"]
	if !errRead {
		panic(errRead)
	}
	return routeList
}
