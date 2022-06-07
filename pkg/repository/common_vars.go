package repository

import "time"

const (
	MaxUsernameLength         = 32
	MaxUserPasswordLength     = 32
	MaxFeedLength             = 30
	MaxFeedBackwardsMilliSec  = 24 * 60 * 60 * 1000
	MaxVideoTitleLength       = 32
	MaxVideoCommentTextLength = 255
	CSTZoneSeconds            = 8 * 3600
	UpdateAllSleepTime        = time.Hour
)
