package queries

import (
	"douyin-lite/app/models"
	"github.com/jmoiron/sqlx"
)

const (
	isFollowQuery        = `SELECT COUNT(*) FROM FOLLOWS WHERE A_ID = ? AND B_ID = ?`
	getFollowsQuery      = `SELECT B_ID FROM FOLLOWS WHERE A_ID = ? AND B_ID IN (?)`
	doFollowQuery        = `INSERT INTO FOLLOWS(A_ID, B_ID) VALUES (?, ?)`
	cancelFollowQuery    = `DELETE FROM FOLLOWS WHERE A_ID = ? AND B_ID = ?`
	getFollowListQuery   = `SELECT * FROM USERS WHERE USER_ID IN (SELECT B_ID FROM FOLLOWS WHERE A_ID = ?)`
	getFollowerListQuery = `SELECT * FROM USERS WHERE USER_ID IN (SELECT A_ID FROM FOLLOWS WHERE B_ID = ?)`
)

func (db *DouyinQuery) GetFollows(curID int, userIDs []int) (rt []int, err error) {
	query, args, err := sqlx.In(getFollowsQuery, curID, userIDs)
	if err != nil {
		return
	}
	err = db.Select(&rt, query, args)
	return
}

func (db *DouyinQuery) IsFollow(AID, BID int) (bool, error) {
	var n int
	err := db.Get(&n, isFollowQuery, AID, BID)
	if err != nil {
		return false, err
	}
	return n == 1, nil
}

func (db *DouyinQuery) DoFollow(AID, BID int) error {
	_, err := db.Exec(doFollowQuery, AID, BID)
	return err
}

func (db *DouyinQuery) CancelFollow(AID, BID int) error {
	_, err := db.Exec(cancelFollowQuery, AID, BID)
	return err
}

func (db *DouyinQuery) GetFollowList(userID int) (rt []models.User, err error) {
	err = db.Select(&rt, getFollowListQuery, userID)
	return
}

func (db *DouyinQuery) GetFollowerList(userID int) (rt []models.User, err error) {
	err = db.Select(&rt, getFollowerListQuery, userID)
	return
}
