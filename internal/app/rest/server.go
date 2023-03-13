package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Speakerkfm/social_network_otus/internal/service/user"
	"github.com/Speakerkfm/social_network_otus/internal/service/user/domain"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(svc user.Service) (*http.Server, error) {
	engine := gin.Default()

	engine.Use(cors.Default())

	swagger, err := GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("fail to get swagger def: %w", err)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.

	RegisterHandlers(engine, &Implementation{
		svc: svc,
	})

	return &http.Server{
		Handler: engine,
		Addr:    "0.0.0.0:80",
	}, nil
}

type Implementation struct {
	svc user.Service
}

// GetDialogUserIdList (GET /dialog/{user_id}/list)
func (i *Implementation) GetDialogUserIdList(c *gin.Context, userId UserId) {
	_ = c.AbortWithError(http.StatusInternalServerError, errors.New("implement me"))
}

// PostDialogUserIdSend (POST /dialog/{user_id}/send)
func (i *Implementation) PostDialogUserIdSend(c *gin.Context, userId UserId) {
	_ = c.AbortWithError(http.StatusInternalServerError, errors.New("implement me"))
}

// PutFriendDeleteUserId (PUT /friend/delete/{user_id})
func (i *Implementation) PutFriendDeleteUserId(c *gin.Context, userId UserId) {
	_ = c.AbortWithError(http.StatusInternalServerError, errors.New("implement me"))
}

// PutFriendSetUserId (PUT /friend/set/{user_id})
func (i *Implementation) PutFriendSetUserId(c *gin.Context, userId UserId) {
	_ = c.AbortWithError(http.StatusInternalServerError, errors.New("implement me"))
}

// PostLogin (POST /login)
func (i *Implementation) PostLogin(c *gin.Context) {
	req := PostLoginJSONBody{}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, N5xx{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	token, err := i.svc.Login(c.Copy(), req.Id, req.Password)
	if err != nil {
		// TODO error handler
		if errors.Is(err, domain.ErrUnauthenticated) {
			c.Status(http.StatusUnauthorized)
			return
		}
		c.JSON(http.StatusInternalServerError, N5xx{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, LoginResponse{Token: token})
}

// PostPostCreate (POST /post/create)
func (i *Implementation) PostPostCreate(c *gin.Context) {
	_ = c.AbortWithError(http.StatusInternalServerError, errors.New("implement me"))
}

// PutPostDeleteId (PUT /post/delete/{id})
func (i *Implementation) PutPostDeleteId(c *gin.Context, id PostId) {
	_ = c.AbortWithError(http.StatusInternalServerError, errors.New("implement me"))

}

// GetPostFeed (GET /post/feed)
func (i *Implementation) GetPostFeed(c *gin.Context, params GetPostFeedParams) {
	_ = c.AbortWithError(http.StatusInternalServerError, errors.New("implement me"))

}

// GetPostGetId (GET /post/get/{id})
func (i *Implementation) GetPostGetId(c *gin.Context, id PostId) {
	_ = c.AbortWithError(http.StatusInternalServerError, errors.New("implement me"))

}

// PutPostUpdate (PUT /post/update)
func (i *Implementation) PutPostUpdate(c *gin.Context) {
	_ = c.AbortWithError(http.StatusInternalServerError, errors.New("implement me"))

}

// GetUserGetId (GET /user/get/{id})
func (i *Implementation) GetUserGetId(c *gin.Context, id UserId) {
	socialUser, err := i.svc.GetUserByID(c.Copy(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, N5xx{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, User{
		Age:        socialUser.Age,
		Biography:  socialUser.Biography,
		Birthdate:  BirthDate{},
		City:       socialUser.City,
		FirstName:  socialUser.FirstName,
		Id:         socialUser.ID,
		SecondName: socialUser.SecondName,
	})
}

// PostUserRegister (POST /user/register)
func (i *Implementation) PostUserRegister(c *gin.Context) {
	req := PostUserRegisterJSONBody{}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, N5xx{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	userID, err := i.svc.Register(c.Copy(), domain.RegisterUserRequest{
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		Age:        req.Age,
		Sex:        req.Sex,
		City:       req.City,
		Biography:  req.Biography,
		Password:   req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, N5xx{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, RegisterResponse{UserId: userID})

}

// GetUserSearch (GET /user/search)
func (i *Implementation) GetUserSearch(c *gin.Context, params GetUserSearchParams) {
	_ = c.AbortWithError(http.StatusInternalServerError, errors.New("implement me"))
}
