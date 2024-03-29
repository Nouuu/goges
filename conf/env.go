package conf

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"os"
)

type Config struct {
	Username         string `env:"KORDIS_USERNAME"`
	Password         string `env:"KORDIS_PASSWORD"`
	SchedulerCron    string `env:"SCHEDULER_CRON"`
	CalendarID       string `env:"CALENDAR_ID"`
	PlanningDaysSync int    `env:"PLANNING_DAYS_SYNC"`
	Mode             string `env:"MODE"`
}

func LoadEnv() (*Config, error) {
	var cfg = &Config{}

	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Validate() error {

	// Validate the PlanningDaysSync field
	if c.PlanningDaysSync < 0 {
		return fmt.Errorf("invalid PlanningDaysSync value: must be a positive integer")
	}

	if len(c.SchedulerCron) > 0 {
		// Validate the SchedulerCron field
		if _, err := cron.ParseStandard(c.SchedulerCron); err != nil {
			return fmt.Errorf("invalid SchedulerCron value: %v", err)
		}
	}

	return nil
}
