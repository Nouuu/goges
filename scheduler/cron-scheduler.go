package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/nouuu/goges/conf"
	"github.com/robfig/cron/v3"
	"os"
	"time"
)

func Main() error {
	cronExpression := os.Getenv(conf.SchedulerCronEnv)
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

	job, err := scheduler.Do(task)

	if err != nil {
		return err
	}

	go monitorJob(job, timezone)

	scheduler.StartBlocking()

	return nil
}

func task() {
	fmt.Println("I am running")
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
