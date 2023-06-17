package pkg

import (
	"github.com/NubeIO/module-core-system/logger"
	"github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Schedule Schedule `yaml:"schedule"`
	LogLevel string   `yaml:"log_level"`
}

type Schedule struct {
	Frequency time.Duration `yaml:"frequency"`
}

var LogLevelMap = []string{
	"fatal",
	"error",
	"panic",
	"info",
	"debug",
	"warn",
	"trace",
}

func (m *Module) DefaultConfig() interface{} {
	schedule := Schedule{
		Frequency: time.Duration(60),
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

	configHasMatchingLogLevel := false
	for _, locLogLevel := range LogLevelMap {
		if newConfig.LogLevel == locLogLevel {
			configHasMatchingLogLevel = true
			break
		}
	}
	if !configHasMatchingLogLevel {
		newConfig.LogLevel = "ERROR"
	}
	logLevel, err := log.ParseLevel(newConfig.LogLevel)
	if err != nil {
		return nil, err
	}
	logger.SetLogger(logLevel) // TODO: Check this

	newConfValid, err := yaml.Marshal(newConfig)
	if err != nil {
		return nil, err
	}
	m.config = newConfig
	log.Info("config is set")
	return newConfValid, nil
}
