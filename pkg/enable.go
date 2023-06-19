package pkg

import (
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"time"
)

var cron *gocron.Scheduler

func (m *Module) Enable() error {
	log.Info("plugin is enabling...")
	cron = gocron.NewScheduler(time.UTC)
	_, _ = cron.Every(m.config.Schedule.Frequency).Tag("ScheduleCheck").Do(m.runSchedule)
	cron.StartAsync()
	log.Info("plugin is enabled")
	return nil
}

func (m *Module) Disable() error {
	log.Info("plugin is disabling...")
	cron.Clear()
	log.Info("plugin is disabled")
	return nil
}
