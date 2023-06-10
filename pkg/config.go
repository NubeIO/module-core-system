package pkg

type Config struct {
	Schedule Schedule `yaml:"schedule"`
	LogLevel string   `yaml:"log_level"`
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

// ValidateAndSetConfig implements plugin.Configurer
func (m *Module) ValidateAndSetConfig(config interface{}) error {
	newConfig := config.(*Config)
	m.config = newConfig
	return nil
}
