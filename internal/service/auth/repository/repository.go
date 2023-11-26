package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/bxcodec/dbresolver/v2"
)

const (
	tableUserSession = "user_session"

	fieldID     = "id"
	fieldToken  = "token"
	fieldUserID = "user_id"
)

var allUserSessionFields = []string{
	fieldID,
	fieldUserID,
	fieldToken,
}

type Implementation struct {
	db dbresolver.DB
}

func New(db dbresolver.DB) *Implementation {
	return &Implementation{
		db: db,
	}
}

func (i *Implementation) CreateSession(ctx context.Context, ses UserSession) error {
	query, args, err := sq.Insert(tableUserSession).
		Columns(allUserSessionFields...).
		Values(ses.ID,
			ses.UserID,
			ses.Token).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	if _, err = i.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

func (i *Implementation) GetSession(ctx context.Context, token string) (UserSession, error) {
	query, args, err := sq.Select(allUserSessionFields...).
		From(tableUserSession).
		Where(sq.Eq{fieldToken: token}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return UserSession{}, err
	}
	res := UserSession{}
	if err = i.db.QueryRowContext(ctx, query, args...).Scan(
		&res.ID,
		&res.UserID,
		&res.Token); err != nil {
		return UserSession{}, err
	}
	return res, nil
}
