package configs

import (
	"os"
	"strconv"
)

var (
	MySQLHost     string
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string
	JWTSecretKey  string
	FFmpegPath    string
	EnableHttps   int
)

func init() {
	MySQLHost = os.Getenv("MYSQL_HOST")
	MySQLUser = os.Getenv("MYSQL_USER")
	MySQLPassword = os.Getenv("MYSQL_PASSWORD")
	MySQLDatabase = os.Getenv("MYSQL_DATABASE")
	JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	FFmpegPath = os.Getenv("FFMPEG_PATH")
	EnableHttps, _ = strconv.Atoi(os.Getenv("ENABLE_HTTPS"))
}
