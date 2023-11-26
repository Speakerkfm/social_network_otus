package post

import (
	"context"
	"errors"
	"fmt"
	"log"

	friend_service "github.com/Speakerkfm/social_network_otus/internal/service/friend"
	"github.com/Speakerkfm/social_network_otus/internal/service/post/domain"
	"github.com/Speakerkfm/social_network_otus/pkg/cache"
)

const (
	batchSize    = 10
	defaultLimit = 100
)

type Service interface {
	Feed(ctx context.Context, userID string, limit, offset int64) ([]domain.Post, error)
}

type Repository interface {
	ListIterator(ctx context.Context, userIDs []string) (domain.PostIterator, error)
}

type Implementation struct {
	cache     cache.Cache
	repo      Repository
	friendSvc friend_service.Service
}

func New(cache cache.Cache, repo Repository, friendSvc friend_service.Service) *Implementation {
	return &Implementation{
		cache:     cache,
		repo:      repo,
		friendSvc: friendSvc,
	}
}

func (i *Implementation) Feed(ctx context.Context, userID string, limit, offset int64) ([]domain.Post, error) {
	if limit == 0 {
		limit = defaultLimit
	}
	startIdx := (offset / batchSize) * batchSize
	endIdx := offset + limit
	newOffset := offset - startIdx
	log.Printf("startIdx: %d, endIdx: %d, newOffset: %d", startIdx, endIdx, newOffset)
	res := make([]domain.Post, 0, limit)
	for {
		if startIdx >= endIdx {
			break
		}
		val, err := i.cache.Get(ctx, fmt.Sprintf("feed_%s_%d_%d", userID, startIdx, startIdx+batchSize))
		if err != nil {
			return nil, err
		}
		if val == nil {
			break
		}
		posts, ok := val.([]domain.Post)
		if !ok {
			return nil, fmt.Errorf("not posts in cache")
		}
		res = append(res, posts...)
		startIdx += batchSize
	}
	if len(res) == 0 {
		go func() {
			if err := i.LoadPostsToCache(ctx, userID); err != nil {
				log.Printf("LoadPostsToCache: %s", err.Error())
			}
		}()
		return nil, nil
	}
	return res, nil
}

func (i *Implementation) LoadPostsToCache(ctx context.Context, userID string) error {
	friendIDs, err := i.friendSvc.GetUserFriends(ctx, userID)
	if err != nil {
		return fmt.Errorf("GetUserFriends: %w", err)
	}
	pi, err := i.repo.ListIterator(ctx, friendIDs)
	if err != nil {
		return fmt.Errorf("ListIterator: %w", err)
	}
	postBatch := make([]domain.Post, 0, batchSize)
	startIdx := 0
	endIdx := 0
	for {
		post, err := pi.Next()
		if errors.Is(err, domain.ErrEOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("PostIterator.Next: %w", err)
		}
		endIdx++
		postBatch = append(postBatch, post)
		if len(postBatch) == batchSize {
			i.cache.Set(ctx, fmt.Sprintf("feed_%s_%d_%d", userID, startIdx, endIdx), postBatch)
			postBatch = make([]domain.Post, 0, batchSize)
			startIdx = endIdx
		}
	}
	if len(postBatch) != 0 {
		i.cache.Set(ctx, fmt.Sprintf("feed_%s_%d_%d", userID, startIdx, endIdx), postBatch)
		postBatch = make([]domain.Post, 0, batchSize)
		startIdx = endIdx
	}
	return nil
}
