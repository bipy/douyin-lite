package queries

import (
	"douyin-lite/app/models"
	"github.com/jmoiron/sqlx"
)

const (
	getUserInfoQuery  = `SELECT * FROM USERS WHERE USER_ID = ?`
	getUserInfosQuery = `SELECT * FROM USERS WHERE USER_ID IN (?)`
)

func (db *DouyinQuery) GetUserInfos(userIDs []int) (rt []models.User, err error) {
	query, args, err := sqlx.In(getUserInfosQuery, userIDs)
	if err != nil {
		return
	}
	err = db.Select(&rt, query, args...)
	return
}

func (db *DouyinQuery) GetUserInfo(ID int) (user *models.User, err error) {
	err = db.Get(user, getUserInfoQuery, ID)
	return
}
