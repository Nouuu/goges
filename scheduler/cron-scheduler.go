package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/nouuu/goges/calendar_sync"
	"github.com/nouuu/goges/conf"
	"github.com/nouuu/goges/google_api"
	"github.com/nouuu/goges/kordis"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

func Main(config *conf.Config) error {
	cronExpression := config.SchedulerCron
	timezone, err := time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		fmt.Printf("Error while loading timezone: %v, using UTC instead", err)
		timezone = time.UTC
	}

	_, err = cron.ParseStandard(cronExpression)

	if err != nil {
		return gocron.ErrCronParseFailure
	}

	scheduler := gocron.NewScheduler(timezone).Cron(cronExpression).SingletonMode()

	job, err := scheduler.Do(task, config)

	if err != nil {
		return err
	}

	go monitorJob(job, timezone)

	scheduler.StartBlocking()

	return nil
}

func task(config *conf.Config) {
	err := calendar_sync.Sync(config.PlanningDaysSync, google_api.GetCalendarService(config), kordis.GetKordisApi(config))
	if err != nil {
		log.Print(err)
	}
}

func monitorJob(job *gocron.Job, timezone *time.Location) {
	time.Sleep(1 * time.Millisecond)

	fmt.Printf("Job is scheduled to run at %v\n", job.NextRun().In(timezone).Format("2006-01-02 15:04:05"))

	job.SetEventListeners(func() {
		// Print current date timestamp format
		fmt.Println(time.Now().In(timezone).Format("2006-01-02 15:04:05"), "Job started")
	}, func() {
		// Print current date timestamp format
		time.Sleep(1 * time.Millisecond)
		fmt.Println(time.Now().In(timezone).Format("2006-01-02 15:04:05"), "Job ended")
		fmt.Println("Next run at", job.NextRun().In(timezone).Format("2006-01-02 15:04:05"))
	})
}
