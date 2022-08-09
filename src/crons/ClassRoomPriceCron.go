package crons

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

type ClassRoomPriceCronAdapter struct {
	Every int
	At    string
}

func New(every int, at string) *ClassRoomPriceCronAdapter {
	return &ClassRoomPriceCronAdapter{
		Every: every,
		At:    at,
	}
}

func (srv *ClassRoomPriceCronAdapter) ProcessSchedule() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(srv.Every).Day().At(srv.At).Do(func() {
		fmt.Println("SST")
	})
	s.StartAsync()
	s.StartBlocking()
}
