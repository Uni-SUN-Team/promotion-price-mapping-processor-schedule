package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(func() {
		fmt.Println("SST")
	})
	s.StartAsync()
	s.StartBlocking()
}
