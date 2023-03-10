package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableSocialUser  = "social_user"
	tableUserSession = "user_session"

	fieldID             = "id"
	fieldFirstName      = "first_name"
	fieldSecondName     = "second_name"
	fieldAge            = "age"
	fieldSex            = "sex"
	fieldCity           = "city"
	fieldBiography      = "biography"
	fieldHashedPassword = "hashed_password"
	fieldToken          = "token"
	fieldUserID         = "user_id"
)

var allSocialUserFields = []string{
	fieldID,
	fieldFirstName,
	fieldSecondName,
	fieldAge,
	fieldSex,
	fieldCity,
	fieldBiography,
	fieldHashedPassword,
}

var allUserSessionFields = []string{
	fieldID,
	fieldUserID,
	fieldToken,
}

type Implementation struct {
	db *sql.DB
}

func New(db *sql.DB) *Implementation {
	return &Implementation{
		db: db,
	}
}

func (i *Implementation) CreateUser(ctx context.Context, usr SocialUser) error {
	query, args, err := sq.Insert(tableSocialUser).
		Columns(allSocialUserFields...).
		Values(usr.ID,
			usr.FirstName,
			usr.SecondName,
			usr.Age,
			usr.Sex,
			usr.City,
			usr.Biography,
			usr.HashedPassword).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	if _, err = i.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

func (i *Implementation) GetUserByID(ctx context.Context, id string) (SocialUser, error) {
	query, args, err := sq.Select(allSocialUserFields...).
		From(tableSocialUser).
		Where(sq.Eq{fieldID: id}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return SocialUser{}, err
	}
	res := SocialUser{}
	if err = i.db.QueryRowContext(ctx, query, args...).Scan(
		&res.ID,
		&res.FirstName,
		&res.SecondName,
		&res.Age,
		&res.Sex,
		&res.City,
		&res.Biography,
		&res.HashedPassword); err != nil {
		return SocialUser{}, err
	}
	return res, nil
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
