package utils

import (
	"douyin-lite/pkg/configs"
	"fmt"
)

func GetMySQLUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		configs.MySQLUser, configs.MySQLPassword, configs.MySQLHost, configs.MySQLDatabase)
}
