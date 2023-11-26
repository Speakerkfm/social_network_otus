package repository

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/bxcodec/dbresolver/v2"
)

const (
	tableSocialUser = "social_user"

	fieldID             = "id"
	fieldFirstName      = "first_name"
	fieldSecondName     = "second_name"
	fieldAge            = "age"
	fieldSex            = "sex"
	fieldCity           = "city"
	fieldBiography      = "biography"
	fieldHashedPassword = "hashed_password"
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

type Implementation struct {
	db dbresolver.DB
}

func New(db dbresolver.DB) *Implementation {
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

func (i *Implementation) UserSearch(ctx context.Context, firstName, secondName string) ([]SocialUser, error) {
	query, args, err := sq.Select(allSocialUserFields...).
		From(tableSocialUser).
		Where(sq.And{
			sq.Like{fieldFirstName: firstName + "%"},
			sq.Like{fieldSecondName: secondName + "%"},
		}).
		OrderBy(fieldID).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := i.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	res := make([]SocialUser, 0)
	defer rows.Close()
	for rows.Next() {
		item := SocialUser{}
		if err = rows.Scan(
			&item.ID,
			&item.FirstName,
			&item.SecondName,
			&item.Age,
			&item.Sex,
			&item.City,
			&item.Biography,
			&item.HashedPassword); err != nil {
			log.Printf("fail to scan user: %s", err.Error())
			continue
		}
		res = append(res, item)
	}
	return res, nil
}
