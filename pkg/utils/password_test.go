package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassword(t *testing.T) {
	password := "asbfakjasbglwige"
	hash, err := GeneratePassword(password)
	assert.NoError(t, err)
	fmt.Println(hash)
	fmt.Println(len(hash))
	assert.True(t, ComparePasswords(hash, password))
	assert.False(t, ComparePasswords(password, password))
}
