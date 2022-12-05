package auth

import "errors"

var ErrorWrongPassword = errors.New("wrong password")
var ErrorWrongUsername = errors.New("wrong username")
var ErrorUserNotFound = errors.New("user not found")
var ErrorUserNotInitialized = errors.New("user instance is not initialized")

var BCryptCost = 51200

type IAuthenticate interface {
	GetPassword() ([]byte, error)
	SetPassword(passwd []byte) (bool, error)
	GetUsername() (string, error)
	SetUsername() (bool, error)
	Authenticate() (bool, error)
}
