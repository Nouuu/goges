package conf

import (
	"github.com/joho/godotenv"
)

const USERNAME_ENV = "username"
const PASSWORD_ENV = "password"
const CalendarIdEnv = "CALENDAR_ID"
const PlanningDaysSyncEnv = "PLANNING_DAYS_SYNC"

func LoadEnv() {
	godotenv.Load(".env")

}
