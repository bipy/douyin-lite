package models

type Video struct {
	Id            int    `json:"id" db:"VIDEO_ID"`
	AuthorId      int    `json:"-" db:"AUTHOR_ID"`
	PlayUrl       string `json:"play_url" db:"PLAY_URL"`
	CoverUrl      string `json:"cover_url" db:"COVER_URL"`
	Title         string `json:"title" db:"TITLE"`
	FavoriteCount int    `json:"favorite_count" db:"FAVORITE_COUNT"`
	CommentCount  int    `json:"comment_count" db:"COMMENT_COUNT"`
	CreateTime    int64  `json:"create_time" db:"CREATE_TIME"`
	Author        *User  `json:"author"`
	IsFavorite    bool   `json:"is_favorite"`
}
