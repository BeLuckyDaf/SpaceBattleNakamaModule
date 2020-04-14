// Copyright 2020 Vladislav Smirnov

package main

import (
	"crypto/sha1"
	"encoding/hex"
	"time"
)

// GeneratePlayerToken uses player username and current time to create a token
func GeneratePlayerToken(username string) string {
	hasher := sha1.New()
	hasher.Write([]byte(username + time.Now().String()))
	token := hex.EncodeToString(hasher.Sum(nil))
	return token
}
