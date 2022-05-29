package schedules

import (
	"douyin-lite/pkg/repository"
	"time"
)

func StartHelloSchedule() {
	go updateId()
}

func updateId() {
	for {
		repository.ID++
		time.Sleep(100 * time.Second)
	}
}
