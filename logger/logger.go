package logger

import "github.com/hashicorp/go-hclog"

func SetLogger(logLevel string) {
	hclog.SetDefault(hclog.New(&hclog.LoggerOptions{
		Name:  "module-core-system",
		Level: hclog.LevelFromString(logLevel),
	}))
}
