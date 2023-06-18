package utils

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

// getNextScheduleDay returns the next scheduled day.
func getNextScheduleDay(today DaysOfTheWeek, scheduleDays []DaysOfTheWeek) (DaysOfTheWeek, error) {
	if len(scheduleDays) == 0 {
		return 0, errors.New("NO DAYS SCHEDULED")
	}
	i := int(today)
	j := 0
	var nextDay DaysOfTheWeek = 0
	nextFound := false
	// check until the next scheduled day is found in scheduleDays
	for j < 7 {
		// check the next day (rollover at end of week [Saturday(6)]
		if i == 6 {
			i = 0
		} else {
			i++
		}
		// check each of the days in scheduleDays
		for x := 0; x < len(scheduleDays); x++ {
			if scheduleDays[x] == DaysOfTheWeek(i) {
				nextDay = DaysOfTheWeek(i)
				nextFound = true
				break
			}
		}
		if nextFound {
			break
		}
		j++
	}
	if !nextFound {
		return 0, errors.New("NO VALID DAYS IN SCHEDULE")
	} else {
		return nextDay, nil
	}
}

// getDurationTillNextScheduleDay returns the time.Duration to the next scheduled day.
func getDurationTillNextScheduleDay(today DaysOfTheWeek, nextDay DaysOfTheWeek) time.Duration {
	// Make Duration to next scheduled day
	daysTillNext := 0
	if nextDay > today {
		daysTillNext = int(nextDay) - int(today)
	} else {
		// Rest of the week, plus the next day as integer
		daysTillNext = (7 - int(today)) + int(nextDay)
	}
	tillNextDayDuration, _ := time.ParseDuration(strconv.Itoa(daysTillNext*24) + "h")
	return tillNextDayDuration
}

// CheckWeeklyScheduleEntry checks if there is a WeeklyScheduleEntry that matches
// the specified schedule Name and is currently within the scheduled period.
func CheckWeeklyScheduleEntry(entry WeeklyScheduleEntry, checkTimezone string) ScheduleCheckerResult {
	result := ScheduleCheckerResult{}
	result.Payload = entry.Value
	result.IsActive = false
	result.IsException = false

	// Get time.Location for entry timezone and check timezone
	location, err := time.LoadLocation(checkTimezone)
	if err != nil {
		result.ErrorFlag = true
		result.ErrorStrings = append(result.ErrorStrings, "Critical: Invalid Timezone")
		return result
	}

	now := time.Now().In(location)
	result.CheckTime = now.Unix()
	// Get day of week and compare with entry.DaysNums
	// nowHour, nowMinute, nowSecond := now.Clock()
	nowYear, nowMonth, nowDate := now.Date()
	// log.Println("nowYear: ", nowYear, "nowMonth: ", nowMonth, "nowDate: ", nowDate)
	nowDayOfWeek := DaysOfTheWeek(now.Weekday())
	// nowDayOfWeekString := now.String()

	// Parse start and stop times
	var entryStartHour, entryStartMins, entryStopHour, entryStopMins int
	n, err1 := fmt.Sscanf(entry.Start, "%d:%d", &entryStartHour, &entryStartMins)
	m, err2 := fmt.Sscanf(entry.End, "%d:%d", &entryStopHour, &entryStopMins)
	if n != 2 || m != 2 || err1 != nil || err2 != nil {
		result.ErrorFlag = true
		result.ErrorStrings = append(result.ErrorStrings, "Critical: Invalid Start/Stop Time")
		return result
	}

	// Parse start and end time into current day timestamps
	startTimestamp := time.Date(nowYear, nowMonth, nowDate, entryStartHour, entryStartMins, 0, 0, location)
	stopTimestamp := time.Date(nowYear, nowMonth, nowDate, entryStopHour, entryStopMins, 59, 0, location)

	// Check if the schedule is active today
	scheduleActiveToday := false
	for _, day := range entry.DaysNums {
		if day == nowDayOfWeek {
			scheduleActiveToday = true
			break
		}
	}

	// Find the next active schedule day
	nextDay, err := getNextScheduleDay(nowDayOfWeek, entry.DaysNums)
	if err != nil {
		result.ErrorFlag = true
		result.ErrorStrings = append(result.ErrorStrings, "Critical: Scheduled Days are Invalid")
		return result
	}
	nextDayDuration := getDurationTillNextScheduleDay(nowDayOfWeek, nextDay)

	if scheduleActiveToday && now.Before(stopTimestamp) { // Scheduled today and hasn't finished yet
		// Check if today's schedule is currently active
		if now.After(startTimestamp) {
			result.IsActive = true
		}
		// PeriodStartStop
		result.PeriodStart = startTimestamp.Unix()
		result.PeriodStop = stopTimestamp.Unix()
		// NextStartStop
		result.NextStart = startTimestamp.Add(nextDayDuration).Unix()
		result.NextStop = stopTimestamp.Add(nextDayDuration).Unix()
	} else { // Not scheduled today, OR scheduled today, but period already passed
		// PeriodStartStop
		result.PeriodStart = startTimestamp.Add(nextDayDuration).Unix()
		result.PeriodStop = stopTimestamp.Add(nextDayDuration).Unix()
		// NextStartStop
		// Next scheduled day is being used for PeriodStart and PeriodStop, so get the NEXT next scheduled day for NextStart and NextStop
		nextNextDay, err := getNextScheduleDay(nextDay, entry.DaysNums)
		if err != nil {
			result.ErrorFlag = true
			result.ErrorStrings = append(result.ErrorStrings, "Critical: Scheduled Days are Invalid")
			return result
		}
		nextNextDayDuration := getDurationTillNextScheduleDay(nextDay, nextNextDay)
		result.NextStart = startTimestamp.Add(nextDayDuration).Add(nextNextDayDuration).Unix()
		result.NextStop = stopTimestamp.Add(nextDayDuration).Add(nextNextDayDuration).Unix()
	}
	AddHumanReadableDatetimes(&result, checkTimezone)
	return result
}

func ConvertDaysStringsToInt(weeklyScheduleEntry WeeklyScheduleEntry) WeeklyScheduleEntry {
	var lowerCaseStringDay string
	for _, v := range weeklyScheduleEntry.Days {
		lowerCaseStringDay = strings.ToLower(v)
		dayInt := DaysMap[lowerCaseStringDay]
		weeklyScheduleEntry.DaysNums = append(weeklyScheduleEntry.DaysNums, dayInt)
	}
	return weeklyScheduleEntry
}

func CheckWeeklyScheduleCollection(scheduleMap TypeWeekly, scheduleName, timezone string) ScheduleCheckerResult {
	finalResult := ScheduleCheckerResult{}
	var singleResult ScheduleCheckerResult
	count := 0
	var err error
	for _, scheduleEntry := range scheduleMap {
		if scheduleEntry.Enable == nil || (scheduleEntry.Enable != nil && *scheduleEntry.Enable) {
			if scheduleName == "ANY" || scheduleName == "ALL" || scheduleName == scheduleEntry.Name {
				scheduleEntry = ConvertDaysStringsToInt(scheduleEntry)
				singleResult = CheckWeeklyScheduleEntry(scheduleEntry, timezone)
				singleResult.Name = scheduleName
				if count == 0 {
					finalResult = singleResult
				} else {
					finalResult, err = CombineScheduleCheckerResults(finalResult, singleResult, timezone)
					if err != nil {
						log.Errorf("CheckEventScheduleEntry %v\n", err)
					}
				}
				count++
			}
		}
	}
	AddHumanReadableDatetimes(&finalResult, timezone)
	return finalResult
}

// WeeklyCheck checks all Weekly Schedules in the payload for active periods.
// It returns a combined ScheduleCheckerResult of all Weekly Schedules.
func WeeklyCheck(weekly TypeWeekly, scheduleName, timezone string) (ScheduleCheckerResult, error) {
	results := CheckWeeklyScheduleCollection(weekly, scheduleName, timezone)
	return results, nil
}
