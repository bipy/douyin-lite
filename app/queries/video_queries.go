package queries

import (
	"douyin-lite/app/models"
	"douyin-lite/pkg/repository"
	"github.com/jmoiron/sqlx"
)

const (
	getFeedQuery     = `SELECT * FROM VIDEOS WHERE CREATE_TIME > ? ORDER BY CREATE_TIME DESC LIMIT ?`
	getFavoriteQuery = `SELECT VIDEO_ID FROM FAVORITES WHERE USER_ID = ? AND VIDEO_ID IN (?)`
)

func (db *DouyinQuery) GetFeed(latestTime int64) ([]models.Video, error) {
	var rt []models.Video
	err := db.Select(&rt, getFeedQuery, latestTime, repository.MaxFeedLength)
	if err != nil {
		return nil, err
	}
	return rt, nil
}

func (db *DouyinQuery) GetFavorite(curID int, videoIDs []int) ([]int, error) {
	var rt []int
	query, args, err := sqlx.In(getFavoriteQuery, curID, videoIDs)
	if err != nil {
		return nil, err
	}
	err = db.Select(&rt, query, args)
	if err != nil {
		return nil, err
	}
	return rt, nil
}
