package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/bxcodec/dbresolver/v2"
)

const (
	tablePost = "post"

	fieldID           = "id"
	fieldAuthorUserID = "author_user_id"
	fieldText         = "text"
	fieldCreatedAt    = "created_at"
	fieldUpdatedAt    = "updated_at"

	defaultPostFeedLimit = 10000
)

var allPostFields = []string{
	fieldID,
	fieldAuthorUserID,
	fieldText,
	fieldCreatedAt,
	fieldUpdatedAt,
}

type Implementation struct {
	db dbresolver.DB
}

func New(db dbresolver.DB) *Implementation {
	return &Implementation{
		db: db,
	}
}

func (i *Implementation) ListIterator(ctx context.Context, userIDs []string) (*PostIterator, error) {
	query, args, err := sq.Select(allPostFields...).
		From(tablePost).
		Where(sq.And{
			sq.Eq{fieldAuthorUserID: userIDs},
		}).
		OrderBy(fieldCreatedAt).
		Limit(defaultPostFeedLimit).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := i.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &PostIterator{
		rows: rows,
	}, nil
}

type PostIterator struct {
	rows *sql.Rows
}

func (pi *PostIterator) Next() (*Post, error) {
	if !pi.rows.Next() {
		if err := pi.rows.Close(); err != nil {
			return nil, err
		}
		return nil, ErrEOF
	}
	item := &Post{}
	if err := pi.rows.Scan(
		&item.ID,
		&item.AuthorUserID,
		&item.Text,
		&item.CreatedAt,
		&item.UpdatedAt); err != nil {
		return nil, err
	}
	return item, nil
}
