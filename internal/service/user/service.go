package user

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/Speakerkfm/social_network_otus/internal/service/user/domain"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	CreateUser(ctx context.Context, usr domain.SocialUser) error
	GetUserByID(ctx context.Context, id string) (domain.SocialUser, error)
	CreateSession(ctx context.Context, ses domain.UserSession) error
}

type Service interface {
	Login(ctx context.Context, id, password string) (string, error)
	Register(ctx context.Context, req domain.RegisterUserRequest) (string, error)
	GetUserByID(ctx context.Context, id string) (domain.SocialUser, error)
}

type Implementation struct {
	repo Repository
}

func New(repo Repository) *Implementation {
	return &Implementation{
		repo: repo,
	}
}

func (i *Implementation) Login(ctx context.Context, id, password string) (string, error) {
	usr, err := i.repo.GetUserByID(ctx, id)
	if err != nil {
		return "", fmt.Errorf("fail to get user by id: %w", err)
	}
	if usr.HashedPassword != hashPassword(password) {
		return "", domain.ErrUnauthenticated
	}
	token := generateToken()
	if err = i.repo.CreateSession(ctx, domain.UserSession{
		ID:     generateID(),
		UserID: usr.ID,
		Token:  token,
	}); err != nil {
		return "", fmt.Errorf("fail to create session")
	}
	return token, err
}

func (i *Implementation) Register(ctx context.Context, req domain.RegisterUserRequest) (string, error) {
	userID := generateID()
	if err := i.repo.CreateUser(ctx, domain.SocialUser{
		ID:             userID,
		FirstName:      req.FirstName,
		SecondName:     req.SecondName,
		Age:            req.Age,
		Sex:            req.Sex,
		City:           req.City,
		Biography:      req.Biography,
		HashedPassword: hashPassword(req.Password),
	}); err != nil {
		return "", fmt.Errorf("fail to register user: %w", err)
	}
	return userID, nil
}

func (i *Implementation) GetUserByID(ctx context.Context, id string) (domain.SocialUser, error) {
	return i.repo.GetUserByID(ctx, id)
}

func generateID() string {
	return uuid.NewV4().String()
}

func generateToken() string {
	return uuid.NewV4().String()
}

func hashPassword(password string) string {
	return hex.EncodeToString(sha1.New().Sum([]byte(password)))
}
