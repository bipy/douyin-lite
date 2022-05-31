package queries

import (
	"douyin-lite/app/models"
	"douyin-lite/pkg/repository"
)

const (
	getFeedQuery        = `SELECT * FROM VIDEOS WHERE CREATE_TIME > ? ORDER BY CREATE_TIME DESC LIMIT ?`
	createVideoQuery    = `INSERT INTO VIDEOS(AUTHOR_ID, PLAY_URL, COVER_URL, TITLE) VALUES (?, ?, ?, ?)`
	getPublishListQuery = `SELECT * FROM VIDEOS WHERE AUTHOR_ID = ?`
)

func (db *DouyinQuery) GetFeed(latestTime int64) (rt []models.Video, err error) {
	err = db.Select(&rt, getFeedQuery, latestTime, repository.MaxFeedLength)
	return
}

func (db *DouyinQuery) CreateVideo(authorID int, playUrl, coverUrl, title string) (int, error) {
	r, err := db.Exec(createVideoQuery, authorID, playUrl, coverUrl, title)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (db *DouyinQuery) GetPublishList(userID int) (rt []models.Video, err error) {
	err = db.Select(&rt, getPublishListQuery, userID)
	return
}
