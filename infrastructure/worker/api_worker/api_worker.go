package api_worker

import (
	"fmt"
	"github.com/InnoSoft/task/infrastructure/seeds"
	"github.com/go-co-op/gocron"
	"time"
)

func task() {
	fmt.Println("seed the database cron")
	seeds.Seed()
}
func RunCron()  {
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(60).Seconds().Do(task)
	s1.StartAsync()
}