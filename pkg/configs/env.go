package configs

import (
	"os"
)

var (
	MySQLHost     string
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string
	JWTSecretKey  string
	FFmpegPath    string
	URLPrefix     string
)

func init() {
	MySQLHost = os.Getenv("MYSQL_HOST")
	MySQLUser = os.Getenv("MYSQL_USER")
	MySQLPassword = os.Getenv("MYSQL_PASSWORD")
	MySQLDatabase = os.Getenv("MYSQL_DATABASE")
	JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	FFmpegPath = os.Getenv("FFMPEG_PATH")
	URLPrefix = os.Getenv("URL_PREFIX")
	if len(URLPrefix) > 0 && URLPrefix[len(URLPrefix)-1] != '/' {
		URLPrefix += "/"
	}
}
