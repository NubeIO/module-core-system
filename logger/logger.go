package logger

import (
	"github.com/NubeIO/module-core-system/pkg"
	log "github.com/sirupsen/logrus"
)

func SetLogger(config *pkg.Config) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(config.LogLevel)
}
