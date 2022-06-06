package queries

import (
	"douyin-lite/app/models"
)

const (
	createCommentQuery = `INSERT INTO COMMENTS(USER_ID, VIDEO_ID, COMMENT_TEXT) VALUES (?, ?, ?)`
	deleteCommentQuery = `UPDATE COMMENTS SET DELETED = 1 WHERE COMMENT_ID = ? AND USER_ID = ?`
	getCommentQuery    = `SELECT COMMENT_TEXT, CREATE_TIME FROM COMMENTS WHERE COMMENT_ID = ?`
	getCommentsQuery   = `SELECT COMMENT_ID, USER_ID, COMMENT_TEXT, CREATE_TIME FROM COMMENTS WHERE VIDEO_ID = ?`
)

func (db *DouyinQuery) CreateComment(userID, videoID int, commentText string) (int, error) {
	r, err := db.Exec(createCommentQuery, userID, videoID, commentText)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	return int(id), err
}

func (db *DouyinQuery) DeleteComment(commentID, opID int) (int, error) {
	r, err := db.Exec(deleteCommentQuery, commentID, opID)
	if err != nil {
		return 0, err
	}
	n, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func (db *DouyinQuery) GetComment(commentID int) (*models.Comment, error) {
	comment := &models.Comment{}
	err := db.Get(comment, getCommentQuery, commentID)
	return comment, err
}

func (db *DouyinQuery) GetComments(videoID int) (rt []models.Comment, err error) {
	err = db.Select(&rt, getCommentsQuery, videoID)
	return
}
