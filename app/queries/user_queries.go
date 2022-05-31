package queries

import (
	"douyin-lite/app/models"
	"github.com/jmoiron/sqlx"
)

const (
	getUserInfoQuery  = `SELECT * FROM USERS WHERE USER_ID = ?`
	getUserInfosQuery = `SELECT * FROM USERS WHERE USER_ID IN (?)`
	getFollowQuery    = `SELECT B_ID FROM FOLLOWS WHERE A_ID = ? AND B_ID IN (?)`
)

func (db *DouyinQuery) GetUserInfos(userIDs []int) ([]models.User, error) {
	var rt []models.User
	query, args, err := sqlx.In(getUserInfosQuery, userIDs)
	if err != nil {
		return nil, err
	}
	err = db.Select(&rt, query, args...)
	if err != nil {
		return nil, err
	}
	return rt, nil
}

func (db *DouyinQuery) GetUserInfo(ID int) (user *models.User, err error) {
	err = db.Get(&user, getUserInfoQuery, ID)
	return
}

func (db *DouyinQuery) GetFollow(curID int, userIDs []int) ([]int, error) {
	var rt []int
	query, args, err := sqlx.In(getFollowQuery, curID, userIDs)
	if err != nil {
		return nil, err
	}
	err = db.Select(&rt, query, args)
	if err != nil {
		return nil, err
	}
	return rt, nil
}
