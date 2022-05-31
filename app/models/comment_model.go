package models

type Comment struct {
	Id         int    `json:"id" db:"COMMENT_ID"`
	AuthorId   int    `json:"user_id" db:"USER_ID"`
	Author     *User  `json:"user"`
	Content    string `json:"content" db:"COMMENT_TEXT"`
	CreateDate string `json:"create_date"`
	CreateTime int64  `json:"create_time" db:"CREATE_TIME"`
}
