package pg

import (
	"context"
	"github.com/Speakerkfm/social_network_otus/internal/service/auth/domain"
	"github.com/Speakerkfm/social_network_otus/internal/service/auth/repository"
)

type Adapter struct {
	repo *repository.Implementation
}

func New(repo *repository.Implementation) *Adapter {
	return &Adapter{
		repo: repo,
	}
}

func (a *Adapter) CreateSession(ctx context.Context, ses domain.UserSession) error {
	return a.repo.CreateSession(ctx, repository.UserSession{
		ID:     ses.ID,
		UserID: ses.UserID,
		Token:  ses.Token,
	})
}

func (a *Adapter) GetSession(ctx context.Context, token string) (domain.UserSession, error) {
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
