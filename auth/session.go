package auth

import (
	secRnd "crypto/rand"
	"encoding/base32"
)

type Session struct {
	Properties []struct {
		Key   string `json:"key"`
		Value any    `json:"value"`
	} `json:"props"`
}

func StartSession() (*Session, error) {
	return &Session{}, nil
}
func (s Session) GenerateSessionID() string {
	id := make([]byte, 15)
	secRnd.Read(id)
	return base32.HexEncoding.EncodeToString(id)
}
