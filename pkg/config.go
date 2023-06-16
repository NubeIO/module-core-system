package pkg

import (
	"github.com/NubeIO/module-core-system/logger"
	"github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Schedule Schedule  `yaml:"schedule"`
	LogLevel log.Level `yaml:"log_level"`
}

type Schedule struct {
	Frequency time.Duration `yaml:"frequency"`
}

func (m *Module) DefaultConfig() interface{} {
	schedule := Schedule{
		Frequency: time.Duration(60),
	}

	return &Config{
		Schedule: schedule,
		LogLevel: log.DebugLevel,
	}
}

func (m *Module) GetConfig() interface{} {
	return m.Config
}

func (m *Module) ValidateAndSetConfig(config []byte) ([]byte, error) {
	newConfig := &Config{}
	if err := yaml.Unmarshal(config, newConfig); err != nil {
		return nil, err
	}
	if newConfig.LogLevel == log.Level(0) {
		newConfig.LogLevel = log.DebugLevel
	}
	newConfValid, err := yaml.Marshal(newConfig)
	if err != nil {
		return nil, err
	}
	m.Config = newConfig
	logger.SetLogger(m.Config.LogLevel)
	log.Info("config is set")
	return newConfValid, nil
}
