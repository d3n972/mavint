package auth

import "errors"

var ErrorWrongPassword = errors.New("Wrong password")
var ErrorWrongUsername = errors.New("Wrong password")
var ErrorUserNotFound = errors.New("Wrong password")
var ErrorUserNotInitialized = errors.New("User instance is not initialized")

var BCryptCost = 51200

type IAuthenticate interface {
	GetPassword() ([]byte, error)
	SetPassword(passwd []byte) (bool, error)
	GetUsername() (string, error)
	SetUsername() (bool, error)
	Authenticate() (bool, error)
}
