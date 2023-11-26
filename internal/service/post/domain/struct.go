package domain

import "errors"

var ErrEOF = errors.New("eof")

type Post struct {
	ID           string
	AuthorUserID string
	Text         string
}

type PostIterator interface {
	Next() (Post, error)
}
