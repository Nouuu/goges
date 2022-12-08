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
	return godotenv.Load(".env")
}
