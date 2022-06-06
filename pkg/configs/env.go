package configs

import "os"

var (
	MySQLHost     string
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string
	RedisHost     string
	RedisPassword string
	RedisDBNum    string
	JWTSecretKey  string
	FFmpegPath    string
)

func init() {
	MySQLHost = os.Getenv("MYSQL_HOST")
	MySQLUser = os.Getenv("MYSQL_USER")
	MySQLPassword = os.Getenv("MYSQL_PASSWORD")
	MySQLDatabase = os.Getenv("MYSQL_DATABASE")
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	RedisDBNum = os.Getenv("REDIS_DB_NUMBER")
	JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	FFmpegPath = os.Getenv("FFMPEG_PATH")
}
