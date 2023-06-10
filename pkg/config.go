package pkg

import (
	"github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Schedule        Schedule      `yaml:"schedule"`
	LogLevel        string        `yaml:"log_level"`
	ReIterationTime time.Duration `yaml:"re_iteration_time"`
}

type Schedule struct {
	Frequency string `yaml:"frequency"`
}

func (m *Module) DefaultConfig() interface{} {
	schedule := Schedule{
		Frequency: "60s",
	}

	return &Config{
		Schedule: schedule,
		LogLevel: "ERROR",
	}
}

func (m *Module) GetConfig() interface{} {
	return m.config
}

func (m *Module) ValidateAndSetConfig(config []byte) ([]byte, error) {
	newConfig := &Config{}
	if err := yaml.Unmarshal(config, newConfig); err != nil {
		return nil, err
	}
	if newConfig.ReIterationTime == 0 {
		newConfig.ReIterationTime = time.Duration(5) * time.Second
	}
	newConfValid, err := yaml.Marshal(newConfig)
	if err != nil {
		return nil, err
	}
	m.config = newConfig
	log.Info("config is set")
	return newConfValid, nil
}
