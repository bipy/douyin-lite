package queries

import (
	"douyin-lite/app/models"
	"github.com/jmoiron/sqlx"
)

const (
	getUserInfoQuery            = `SELECT USER_ID, USER_NAME, FOLLOW_COUNT, FOLLOWER_COUNT FROM USERS WHERE USER_ID = ?`
	getUserInfosQuery           = `SELECT USER_ID, USER_NAME, FOLLOW_COUNT, FOLLOWER_COUNT FROM USERS WHERE USER_ID IN (?)`
	addFollowCountQuery         = `UPDATE USERS SET FOLLOW_COUNT = FOLLOW_COUNT + ? WHERE USER_ID = ?`
	addFollowerCountQuery       = `UPDATE USERS SET FOLLOWER_COUNT = FOLLOWER_COUNT + ? WHERE USER_ID = ?`
	updateAllFollowCountQuery   = `UPDATE USERS U LEFT JOIN (SELECT F.A_ID AS A, COUNT(*) AS CNT FROM FOLLOWS F GROUP BY A) AS UF ON U.USER_ID = A SET U.FOLLOW_COUNT = IFNULL(CNT, 0)`
	updateAllFollowerCountQuery = `UPDATE USERS U LEFT JOIN (SELECT F.B_ID AS B, COUNT(*) AS CNT FROM FOLLOWS F GROUP BY B) AS UF ON U.USER_ID = B SET U.FOLLOWER_COUNT = IFNULL(CNT, 0)`
)

func (db *DouyinQuery) GetUserInfos(userIDs []int) (rt []models.User, err error) {
	query, args, err := sqlx.In(getUserInfosQuery, userIDs)
	if err != nil {
		return
	}
	err = db.Select(&rt, query, args...)
	return
}

func (db *DouyinQuery) GetUserInfo(ID int) (*models.User, error) {
	user := &models.User{}
	err := db.Get(user, getUserInfoQuery, ID)
	return user, err
}

func (db *DouyinQuery) AddFollowCount(n, userID int) error {
	_, err := db.Exec(addFollowCountQuery, n, userID)
	return err
}

func (db *DouyinQuery) AddFollowerCount(n, userID int) error {
	_, err := db.Exec(addFollowerCountQuery, n, userID)
	return err
}

func (db *DouyinQuery) UpdateAllFollowCount() error {
	_, err := db.Exec(updateAllFollowCountQuery)
	return err
}

func (db *DouyinQuery) UpdateAllFollowerCount() error {
	_, err := db.Exec(updateAllFollowerCountQuery)
	return err
}
