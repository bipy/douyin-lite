package queries

import (
	"douyin-lite/platform"
	"github.com/jmoiron/sqlx"
)

type DouyinQuery struct {
	*sqlx.DB
}

var (
	DouyinDB *DouyinQuery
)

func init() {
	DouyinDB = &DouyinQuery{DB: platform.GetNewMySQLConn()}
}
