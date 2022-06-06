package schedules

import (
	"douyin-lite/app/queries"
	"douyin-lite/pkg/repository"
	"log"
	"math/rand"
	"time"
)

func UpdateFavoriteCount() {
	for {
		time.Sleep(time.Duration(rand.Intn(int(repository.UpdateAllSleepTime))))
		err := queries.DouyinDB.UpdateAllFavoriteCount()
		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(repository.UpdateAllSleepTime)
	}
}

func UpdateCommentCount() {
	for {
		time.Sleep(time.Duration(rand.Intn(int(repository.UpdateAllSleepTime))))
		err := queries.DouyinDB.UpdateAllCommentCount()
		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(repository.UpdateAllSleepTime)
	}
}
