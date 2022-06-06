package schedules

import (
	"douyin-lite/app/queries"
	"douyin-lite/pkg/repository"
	"log"
	"math/rand"
	"time"
)

func UpdateFollowCount() {
	for {
		time.Sleep(time.Duration(rand.Intn(int(repository.UpdateAllSleepTime))))
		err := queries.DouyinDB.UpdateAllFollowCount()
		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(repository.UpdateAllSleepTime)
	}
}

func UpdateFollowerCount() {
	for {
		time.Sleep(time.Duration(rand.Intn(int(repository.UpdateAllSleepTime))))
		err := queries.DouyinDB.UpdateAllFollowerCount()
		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(repository.UpdateAllSleepTime)
	}
}
