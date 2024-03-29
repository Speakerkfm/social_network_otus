package pg_adapter

import (
	"context"

	"github.com/Speakerkfm/social_network_otus/internal/service/user/domain"
	"github.com/Speakerkfm/social_network_otus/internal/service/user/repository"
)

type Adapter struct {
	repo *repository.Implementation
}

func New(repo *repository.Implementation) *Adapter {
	return &Adapter{
		repo: repo,
	}
}

func (a *Adapter) CreateUser(ctx context.Context, usr domain.SocialUser) error {
	return a.repo.CreateUser(ctx, repository.SocialUser{
		ID:             usr.ID,
		FirstName:      usr.FirstName,
		SecondName:     usr.SecondName,
		Age:            usr.Age,
		Sex:            convertSexToPg(usr.Sex),
		City:           usr.City,
		Biography:      usr.Biography,
		HashedPassword: usr.HashedPassword,
	})
}

func (a *Adapter) GetUserByID(ctx context.Context, id string) (domain.SocialUser, error) {
	usr, err := a.repo.GetUserByID(ctx, id)
	if err != nil {
		return domain.SocialUser{}, err
	}
	return domain.SocialUser{
		ID:             usr.ID,
		FirstName:      usr.FirstName,
		SecondName:     usr.SecondName,
		Age:            usr.Age,
		Sex:            convertSexFromPg(usr.Sex),
		City:           usr.City,
		Biography:      usr.Biography,
		HashedPassword: usr.HashedPassword,
	}, nil
}

func (a *Adapter) UserSearch(ctx context.Context, firstName, secondName string) ([]domain.SocialUser, error) {
	users, err := a.repo.UserSearch(ctx, firstName, secondName)
	if err != nil {
		return nil, err
	}
	res := make([]domain.SocialUser, 0, len(users))
	for _, usr := range users {
		res = append(res, domain.SocialUser{
			ID:             usr.ID,
			FirstName:      usr.FirstName,
			SecondName:     usr.SecondName,
			Age:            usr.Age,
			Sex:            convertSexFromPg(usr.Sex),
			City:           usr.City,
			Biography:      usr.Biography,
			HashedPassword: usr.HashedPassword,
		})
	}
	return res, nil
}

func convertSexToPg(sex string) int {
	switch sex {
	case "male":
		return 1
	case "female":
		return 0
	}
	return 0
}

func convertSexFromPg(sex int) string {
	switch sex {
	case 1:
		return "male"
	case 0:
		return "female"
	}
	return "female"
}
