package queries

import (
	"douyin-lite/app/models"
	"douyin-lite/pkg/repository"
)

const (
	getFeedQuery                = `SELECT * FROM VIDEOS WHERE CREATE_TIME > ? ORDER BY CREATE_TIME DESC LIMIT ?`
	createVideoQuery            = `INSERT INTO VIDEOS(AUTHOR_ID, PLAY_URL, COVER_URL, TITLE) VALUES (?, ?, ?, ?)`
	getPublishListQuery         = `SELECT * FROM VIDEOS WHERE AUTHOR_ID = ?`
	addFavoriteCountQuery       = `UPDATE VIDEOS SET FAVORITE_COUNT = FAVORITE_COUNT + ? WHERE VIDEO_ID = ?`
	addCommentCountQuery        = `UPDATE VIDEOS SET COMMENT_COUNT = COMMENT_COUNT + ? WHERE VIDEO_ID = ?`
	updateAllCommentCountQuery  = `UPDATE VIDEOS V LEFT JOIN (SELECT CO.VIDEO_ID AS VID, COUNT(*) AS CNT FROM COMMENTS CO GROUP BY VID) AS VC ON V.VIDEO_ID = VID SET V.COMMENT_COUNT = IFNULL(CNT, 0)`
	updateAllFavoriteCountQuery = `UPDATE VIDEOS V LEFT JOIN (SELECT F.VIDEO_ID AS VID, COUNT(*) AS CNT FROM FAVORITES F GROUP BY VID)  AS VF ON V.VIDEO_ID = VID SET V.FAVORITE_COUNT = IFNULL(CNT, 0)`
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

func (db *DouyinQuery) AddFavoriteCount(n, videoID int) error {
	_, err := db.Exec(addFavoriteCountQuery, n, videoID)
	return err
}

func (db *DouyinQuery) AddCommentCount(n, videoID int) error {
	_, err := db.Exec(addCommentCountQuery, n, videoID)
	return err
}

func (db *DouyinQuery) UpdateAllCommentCount() error {
	_, err := db.Exec(updateAllCommentCountQuery)
	return err
}

func (db *DouyinQuery) UpdateAllFavoriteCount() error {
	_, err := db.Exec(updateAllFavoriteCountQuery)
	return err
}
