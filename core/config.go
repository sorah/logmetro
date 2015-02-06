package core

import (
	"gopkg.in/yaml.v2"
	"ioutil"
)

type Config struct {
	Inputs          []*ConfigInput
	Outputs         []*ConfigOutput
	Web             ConfigWeb
	ServiceGrouping []*ConfigServiceGrouping
	Services        map[string]*ConfigService
}

type ConfigInput struct {
	Type    string
	Options interface{}
}

type ConfigWeb struct {
	Listen string
	Cors   string
	Ui     string
}

type ConfigServiceGrouping struct {
	Key             string
	DefaultTemplate string `yaml:"default"`
}

type ConfigService struct {
	Items     map[string]*ConfigServiceItem
	Time      string
	Realtime  bool
	Dashboard ConfigServiceDashboard
	Window    ConfigServiceWindow
	Groups    []*ConfigServiceGroup
}

type ConfigServiceItem struct {
	Type    string
	Options interface{}
}

type ConfigServiceWindow struct {
	Length    string
	Retention string
}

type ConfigServiceGroup struct {
	Key               string
	Filter            string
	ExcludeFromParent bool
}

type ConfigOutput struct {
	Type    string
	Options interface{}
}

func LoadConfigFromFile(path string) (config *Config, err error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}
	return ParseConfig(bs)
}

func ParseConfig(configByte []byte) (config *Config, err error) {
	config := make(Config)
	err := yaml.Unmarshal(configByte, &config)
}
