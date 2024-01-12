package demo

import (
	"time"

	"github.com/go-co-op/gocron"
)

func (d *Demo) runJobs() {
	d.scheduler = gocron.NewScheduler(time.UTC)

	// job list
	d.scheduler.Every(3).Second().SingletonMode().Do(d.polling)

	d.scheduler.StartAsync()
}

func (d *Demo) polling() {
	log.Info("polling jobs...")
}
