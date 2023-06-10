package pkg

import (
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"time"
)

var pluginName = "module-core-system"
var cron *gocron.Scheduler

func (m *Module) Enable() error {
	log.Info("plugin is enabling...")
	networks, err := m.grpcMarshaller.GetNetworksByPluginName(pluginName, "") // TODO: Check this

	if err != nil {
		log.Error(err)
	}

	if len(networks) == 0 {
		log.Warn("we don't have networks")
	}

	network := networks[0]
	m.networkUUID = network.UUID

	cron = gocron.NewScheduler(time.UTC)
	var frequency = "60s"
	if m.config.Schedule.Frequency != "" {
		frequency = m.config.Schedule.Frequency
	}
	_, _ = cron.Every(frequency).Tag("ScheduleCheck").Do(m.runSchedule)
	cron.StartAsync()

	// Not sure if this is needed or not & which network to map to
	// TODO: This is added just to update the EnableWriteable property on legacy points.
	// It should be removed in future versions.
	// networks, err := inst.db.GetNetworks(api.Args{WithDevices: true, WithPoints: true, WithPriority: true})
	// if err != nil {
	// 	log.Error("SYSTEM Enable() GetNetworks error:", err)
	// 	return nil
	// }
	// fmt.Println("SYSTEM PLUGIN ENABLE: TEMPORARY CODE TO SET EnableWriteable PROPERTY ON EVERY POINT.  REMOVE IN FUTURE RELEASES")
	// for _, network := range networks {
	// 	for _, device := range network.Devices {
	// 		for _, point := range device.Points {
	// 			fmt.Println("SYSTEM PLUGIN ENABLE point: ", point.Name)
	// 			inst.db.UpdatePointPlugin(point.UUID, point)
	// 		}
	// 	}
	// }

	log.Info("plugin is enabled")
	return nil
}

func (m *Module) Disable() error {
	log.Info("plugin is disabling...")
	cron.Clear()
	log.Info("plugin is disabled")
	return nil
}
