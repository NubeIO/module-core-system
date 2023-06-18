package pkg

import (
	"fmt"
	"github.com/NubeIO/module-core-system/utils"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/times/utilstime"
	"github.com/NubeIO/rubix-os/utils/boolean"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func (m *Module) runSchedule() {
	schedules, err := m.grpcMarshaller.GetSchedules()
	if err != nil {
		log.Errorf("Schedule Checks: GetSchedules %s", err.Error())
		return
	} else {
		log.Debugf("Schedule Checks: run schedule checks, schedule count: %d", len(schedules))
	}

	for _, sch := range schedules {
		scheduleJSON, err := utils.DecodeSchedule(sch.Schedule)
		if err != nil {
			log.Errorf("Schedule Checks: issue on DecodeSchedule %v\n", err)
			return
		}
		if !boolean.IsTrue(sch.Enable) {
			log.Debugf("Schedule Checks: runSchedule sch is not enabled so skip logic. name: %s", sch.Name)
			return
		}

		scheduleNameToCheck := "ALL" // TODO: we may need a way to specify the schedule name that is being checked for.

		var timezone = scheduleJSON.Config.TimeZone
		if timezone == "" {
			timezone = sch.TimeZone
		}

		_, err = time.LoadLocation(timezone)
		if timezone == "" || err != nil {
			log.Error("Schedule Checks: CheckWeeklyScheduleCollection: no timezone pass in from user")
			systemTimezone := strings.Split((*utilstime.SystemTime()).HardwareClock.Timezone, " ")[0]
			if systemTimezone == "" {
				zone, _ := utilstime.GetHardwareTZ()
				timezone = zone
			} else {
				timezone = systemTimezone
			}
			sch.TimeZone = timezone
		}

		// CHECK WEEKLY SCHEDULES
		weeklyResult, err := utils.WeeklyCheck(scheduleJSON.Schedules.Weekly, scheduleNameToCheck, timezone)
		if err != nil {
			log.Errorf("Schedule Checks: issue on WeeklyCheck %v\n", err)
		} else {
			log.Debugf("Schedule Checks: weekly schedule: %s is-active %t", weeklyResult.Name, weeklyResult.IsActive)
		}

		// CHECK EVENT SCHEDULES
		eventResult, err := utils.EventCheck(scheduleJSON.Schedules.Events, scheduleNameToCheck, timezone)
		if err != nil {
			log.Errorf("Schedule Checks: issue on eventResult %s", err.Error())
		} else {
			log.Debugf("Schedule Checks: event schedule: %s is-active: %t", eventResult.Name, eventResult.IsActive)
		}
		log.Debugf("Schedule Checks: eventResult: %+v", eventResult)

		// 	COMBINE EVENT AND WEEKLY SCHEDULE RESULTS
		weeklyAndEventResult, err := utils.CombineScheduleCheckerResults(weeklyResult, eventResult, timezone)
		if err != nil {
			log.Errorf("Schedule Checks: issue on weeklyAndEventResult %s", err.Error())
		} else {
			log.Debugf("Schedule Checks: weekly & event schedule: %s is-active: %t", weeklyAndEventResult.Name, weeklyAndEventResult.IsActive)
		}
		log.Debugf("Schedule Checks: weeklyAndEventResult: %+v", weeklyAndEventResult)

		// CHECK EXCEPTION SCHEDULES
		exceptionResult, err := utils.ExceptionCheck(scheduleJSON.Schedules.Exceptions, scheduleNameToCheck, timezone) // This will check for any active schedules with defined name.
		if err != nil {
			log.Errorf("Schedule Checks: issue on exceptionResult %s", err.Error())
		} else {
			log.Debugf("Schedule Checks: exception schedule: %s  is-active: %t", exceptionResult.Name, exceptionResult.IsActive)
		}
		if exceptionResult.CheckIfEmpty() {
			log.Debugf("Schedule Checks: exception schedule is empty: %s", exceptionResult.Name)
		}
		log.Debugf("Schedule Checks: exceptionResult: %+v", exceptionResult)

		finalResult, err := utils.ApplyExceptionSchedule(weeklyAndEventResult, exceptionResult, timezone) // This applies the exception schedule to mask the combined weekly and event schedules.
		if err != nil {
			log.Error(fmt.Sprintf("Schedule Checks: final-result: %s", err.Error()))
		}
		log.Debugf("Schedule Checks: final-result: %s  is-active: %t timezone: %s", finalResult.Name, finalResult.IsActive, timezone)
		log.Debugf("Schedule Checks: finalResult: %+v", finalResult)

		if sch != nil {
			m.store.Set(sch.Name, finalResult, -1)
			sch.IsActive = boolean.New(finalResult.IsActive)
			sch.ActiveWeekly = boolean.New(weeklyResult.IsActive)
			sch.ActiveException = boolean.New(exceptionResult.IsActive)
			sch.ActiveEvent = boolean.New(eventResult.IsActive)
			sch.Payload = finalResult.Payload

			sch.PeriodStart = finalResult.PeriodStart
			if finalResult.PeriodStart == 0 {
				sch.PeriodStartString = ""
			} else {
				sch.PeriodStartString = finalResult.PeriodStartString
			}

			sch.PeriodStop = finalResult.PeriodStop
			if finalResult.PeriodStop == 0 {
				sch.PeriodStopString = ""
			} else {
				sch.PeriodStopString = finalResult.PeriodStopString
			}

			sch.NextStart = finalResult.NextStart
			if finalResult.NextStart == 0 {
				sch.NextStartString = ""
			} else {
				sch.NextStartString = finalResult.NextStartString
			}

			sch.NextStop = finalResult.NextStop
			if finalResult.NextStop == 0 {
				sch.NextStopString = ""
			} else {
				sch.NextStopString = finalResult.NextStopString
			}

			_, err = m.grpcMarshaller.UpdateScheduleAllProps(sch.UUID, sch)
			if err != nil {
				log.Errorf("Schedule Checks: issue on UpdateSchedule %s, error: %v", sch.UUID, err)
			}
		}
	}
	return
}
