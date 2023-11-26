package friend_service

import (
	"context"
)

type Service interface {
	GetUserFriends(ctx context.Context, userID string) ([]string, error)
}

type Repository interface {
	GetUserFriends(ctx context.Context, userID string) ([]string, error)
}

type Implementation struct {
	repo Repository
}

func New(repo Repository) *Implementation {
	return &Implementation{
		repo: repo,
	}
}

func (i *Implementation) GetUserFriends(ctx context.Context, userID string) ([]string, error) {
	return i.repo.GetUserFriends(ctx, userID)
}
