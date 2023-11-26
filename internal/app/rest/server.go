package rest

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Speakerkfm/social_network_otus/internal/service/auth"
	"github.com/Speakerkfm/social_network_otus/internal/service/post"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Speakerkfm/social_network_otus/internal/service/user"
	"github.com/Speakerkfm/social_network_otus/internal/service/user/domain"
)

func New(userSvc user.Service, authSvc auth.Service, postSvc post.Service) (*http.Server, error) {
	engine := gin.Default()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AllowHeaders = []string{"*"}
	engine.Use(cors.New(corsCfg))

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
		userSvc: userSvc,
		authSvc: authSvc,
		postSvc: postSvc,
	})

	return &http.Server{
		Handler: engine,
		Addr:    "0.0.0.0:80",
	}, nil
}

type Implementation struct {
	userSvc user.Service
	authSvc auth.Service
	postSvc post.Service
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
	token, err := i.userSvc.Login(c.Copy(), req.Id, req.Password)
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
	ctx := c.Copy()
	token := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer ")
	userSession, err := i.authSvc.GetSession(ctx, token)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	log.Printf("%+v", userSession)
	posts, err := i.postSvc.Feed(ctx, userSession.UserID, 0, 0)
	if err != nil {
		log.Printf("postSvc.Feed: %s", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, posts)
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
	socialUser, err := i.userSvc.GetUserByID(c.Copy(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, N5xx{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, User{
		Age:        socialUser.Age,
		Biography:  socialUser.Biography,
		Sex:        socialUser.Sex,
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
	userID, err := i.userSvc.Register(c.Copy(), domain.RegisterUserRequest{
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
	socialUsers, err := i.userSvc.UserSearch(c.Copy(), params.FirstName, params.LastName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, N5xx{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	res := make([]User, 0, len(socialUsers))
	for _, socialUser := range socialUsers {
		res = append(res, User{
			Age:        socialUser.Age,
			Biography:  socialUser.Biography,
			Sex:        socialUser.Sex,
			City:       socialUser.City,
			FirstName:  socialUser.FirstName,
			Id:         socialUser.ID,
			SecondName: socialUser.SecondName,
		})
	}

	c.JSON(http.StatusOK, res)
}
