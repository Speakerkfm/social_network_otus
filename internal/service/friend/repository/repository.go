package friend_repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/bxcodec/dbresolver/v2"
)

const (
	tableUserFriend = "user_friend"

	fieldUserID   = "user_id"
	fieldFriendID = "friend_id"
)

var allUserFriendFields = []string{
	fieldUserID,
	fieldFriendID,
}

type Implementation struct {
	db dbresolver.DB
}

func New(db dbresolver.DB) *Implementation {
	return &Implementation{
		db: db,
	}
}

func (i *Implementation) GetUserFriends(ctx context.Context, userID string) ([]string, error) {
	query, args, err := sq.Select(allUserFriendFields...).
		From(tableUserFriend).
		Where(sq.And{
			sq.Like{fieldUserID: userID},
		}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := i.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0)
	defer rows.Close()
	for rows.Next() {
		item := UserFriend{}
		if err = rows.Scan(
			&item.UserID,
			&item.FriendID); err != nil {
			continue
		}
		res = append(res, item.FriendID)
	}
	return res, nil
}
