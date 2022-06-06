package queries

import (
	"douyin-lite/app/models"
	"github.com/jmoiron/sqlx"
)

const (
	doFavoriteQuery      = `INSERT INTO FAVORITES(USER_ID, VIDEO_ID) VALUES (?, ?)`
	cancelFavoriteQuery  = `DELETE FROM FAVORITES WHERE USER_ID = ? AND VIDEO_ID = ?`
	getFavoriteQuery     = `SELECT VIDEO_ID FROM FAVORITES WHERE USER_ID = ? AND VIDEO_ID IN (?)`
	getFavoriteListQuery = `SELECT * FROM VIDEOS WHERE VIDEO_ID IN (SELECT VIDEO_ID FROM FAVORITES WHERE USER_ID = ?)`
)

func (db *DouyinQuery) DoFavorite(userID, videoID int) error {
	_, err := db.Exec(doFavoriteQuery, userID, videoID)
	return err
}

func (db *DouyinQuery) CancelFavorite(userID, videoID int) error {
	_, err := db.Exec(cancelFavoriteQuery, userID, videoID)
	return err
}

func (db *DouyinQuery) GetFavorite(curID int, videoIDs []int) (rt []int, err error) {
	query, args, err := sqlx.In(getFavoriteQuery, curID, videoIDs)
	if err != nil {
		return
	}
	err = db.Select(&rt, query, args...)
	return
}

func (db *DouyinQuery) GetFavoriteList(userID int) (rt []models.Video, err error) {
	err = db.Select(&rt, getFavoriteListQuery, userID)
	return
}
