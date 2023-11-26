package friend_pg_adapter

import (
	"context"
	"errors"

	"github.com/Speakerkfm/social_network_otus/internal/service/friend/repository"
	"github.com/Speakerkfm/social_network_otus/internal/service/post/domain"
	"github.com/Speakerkfm/social_network_otus/internal/service/post/repository"
)

type Adapter struct {
	repo *friend_repository.Implementation
}

func New(repo *friend_repository.Implementation) *Adapter {
	return &Adapter{
		repo: repo,
	}
}

func (a *Adapter) GetUserFriends(ctx context.Context, userID string) ([]string, error) {
	return a.repo.GetUserFriends(ctx, userID)
}

type PostIterator struct {
	pi *repository.PostIterator
}

func (pi *PostIterator) Next() (domain.Post, error) {
	post, err := pi.pi.Next()
	if errors.Is(err, repository.ErrEOF) {
		return domain.Post{}, domain.ErrEOF
	}
	if err != nil {
		return domain.Post{}, err
	}
	return domain.Post{
		ID:           post.ID,
		AuthorUserID: post.AuthorUserID,
		Text:         post.Text,
	}, nil
}
