package conf

import (
	"github.com/joho/godotenv"
)

const UsernameEnv = "USERNAME"
const PasswordEnv = "PASSWORD"
const SchedulerCronEnv = "SCHEDULER_CRON"
const CalendarIdEnv = "CALENDAR_ID"
const PlanningDaysSyncEnv = "PLANNING_DAYS_SYNC"

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}
