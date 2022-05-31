package utils

import "errors"

type TokenMetadata struct {
	UserID      int
	Credentials map[string]bool
	Expires     int64
}

func Verify(token string) (id int, err error) {
	if len(token) == 0 {
		return 0, errors.New("empty token")
	}
	// TODO
	return
}
