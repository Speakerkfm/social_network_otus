package auth

import (
	"context"
	"github.com/Speakerkfm/social_network_otus/internal/service/auth/domain"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	GetSession(ctx context.Context, token string) (domain.UserSession, error)
	CreateSession(ctx context.Context, ses domain.UserSession) error
}

type Service interface {
	GetSession(ctx context.Context, token string) (domain.UserSession, error)
	CreateSession(ctx context.Context, userID string) (string, error)
}

type Implementation struct {
	repo Repository
}

func New(repo Repository) *Implementation {
	return &Implementation{
		repo: repo,
	}
}

func (a *Implementation) CreateSession(ctx context.Context, userID string) (string, error) {
	token := generateToken()
	err := a.repo.CreateSession(ctx, domain.UserSession{
		ID:     generateID(),
		UserID: userID,
		Token:  token,
	})
	return token, err
}

func (a *Implementation) GetSession(ctx context.Context, token string) (domain.UserSession, error) {
	res, err := a.repo.GetSession(ctx, token)
	if err != nil {
		return domain.UserSession{}, err
	}
	return domain.UserSession{
		ID:     res.ID,
		Token:  res.Token,
		UserID: res.UserID,
	}, nil
}

func generateID() string {
	return uuid.NewV4().String()
}

func generateToken() string {
	return uuid.NewV4().String()
}
