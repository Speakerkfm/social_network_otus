package repository

import (
	"errors"
	"time"
)

var ErrEOF = errors.New("eof")

type Post struct {
	ID           string
	AuthorUserID string
	Text         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
