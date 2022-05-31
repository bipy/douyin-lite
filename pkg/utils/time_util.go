package utils

import (
	"douyin-lite/pkg/repository"
	"time"
)

func UnixTimeStampToCSTMMdd(ts int64) string {
	return time.Unix(ts+repository.CSTZoneSeconds, 0).Format("01-02")
}
