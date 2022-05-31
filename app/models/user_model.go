package models

type User struct {
	Id            int    `json:"id" db:"USER_ID"`
	Name          string `json:"name" db:"USER_NAME"`
	FollowCount   int    `json:"follow_count" db:"FOLLOW_COUNT"`
	FollowerCount int    `json:"follower_count" db:"FOLLOWER_COUNT"`
	IsFollow      bool   `json:"is_follow"`
}
