package utils

type TokenMetadata struct {
	UserID      int
	Credentials map[string]bool
	Expires     int64
}
