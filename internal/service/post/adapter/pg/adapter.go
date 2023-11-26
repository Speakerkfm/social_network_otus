package pg_adapter

import (
	"context"
	"errors"

	"github.com/Speakerkfm/social_network_otus/internal/service/post/domain"
	"github.com/Speakerkfm/social_network_otus/internal/service/post/repository"
)

type Adapter struct {
	repo *repository.Implementation
}

func New(repo *repository.Implementation) *Adapter {
	return &Adapter{
		repo: repo,
	}
}

func (a *Adapter) ListIterator(ctx context.Context, userIDs []string) (domain.PostIterator, error) {
	pi, err := a.repo.ListIterator(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	return &PostIterator{
		pi: pi,
	}, nil
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
