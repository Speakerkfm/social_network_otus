package domain

import (
	"errors"
)

var (
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrUserNotFound    = errors.New("user not found")
)

type SocialUser struct {
	ID             string
	FirstName      string
	SecondName     string
	Age            int
	Sex            string
	City           string
	Biography      string
	HashedPassword string
}

type UserSession struct {
	ID     string
	UserID string
	Token  string
}

type RegisterUserRequest struct {
	FirstName  string
	SecondName string
	Age        int
	Sex        string
	City       string
	Biography  string
	Password   string
}
