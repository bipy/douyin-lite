package platform

import (
	"douyin-lite/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetNewMySQLConn() *sqlx.DB {
	return sqlx.MustConnect("mysql", utils.GetMySQLUrl())
}
