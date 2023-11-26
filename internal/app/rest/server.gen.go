// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package rest

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// BirthDate Дата рождения
type BirthDate = openapi_types.Date

// DialogMessage defines model for DialogMessage.
type DialogMessage struct {
	// From Идентификатор пользователя
	From UserId `json:"from"`

	// Text Текст сообщения
	Text DialogMessageText `json:"text"`

	// To Идентификатор пользователя
	To UserId `json:"to"`
}

// DialogMessageText Текст сообщения
type DialogMessageText = string

// Post Пост пользователя
type Post struct {
	// AuthorUserId Идентификатор пользователя
	AuthorUserId UserId `json:"author_user_id"`

	// Id Идентификатор поста
	Id PostId `json:"id"`

	// Text Текст поста
	Text PostText `json:"text"`
}

// PostId Идентификатор поста
type PostId = string

// PostText Текст поста
type PostText = string

// User defines model for User.
type User struct {
	// Age Возраст
	Age int `json:"age"`

	// Biography Интересы
	Biography string `json:"biography"`

	// City Город
	City string `json:"city"`

	// FirstName Имя
	FirstName string `json:"first_name"`

	// Id Идентификатор пользователя
	Id UserId `json:"id"`

	// SecondName Фамилия
	SecondName string `json:"second_name"`
	Sex        string `json:"sex"`
}

// UserId Идентификатор пользователя
type UserId = string

// N5xx defines model for 5xx.
type N5xx struct {
	// Code Код ошибки. Предназначен для классификации проблем и более быстрого решения проблем.
	Code int `json:"code"`

	// Message Описание ошибки
	Message string `json:"message"`

	// RequestId Идентификатор запроса. Предназначен для более быстрого поиска проблем.
	RequestId string `json:"request_id"`
}

// LoginResponse defines model for LoginResponse.
type LoginResponse struct {
	Token string `json:"token"`
}

// RegisterResponse defines model for RegisterResponse.
type RegisterResponse struct {
	UserId string `json:"user_id"`
}

// PostDialogUserIdSendJSONBody defines parameters for PostDialogUserIdSend.
type PostDialogUserIdSendJSONBody struct {
	// Text Текст сообщения
	Text DialogMessageText `json:"text"`
}

// PostLoginJSONBody defines parameters for PostLogin.
type PostLoginJSONBody struct {
	// Id Идентификатор пользователя
	Id       UserId `json:"id"`
	Password string `json:"password"`
}

// PostPostCreateJSONBody defines parameters for PostPostCreate.
type PostPostCreateJSONBody struct {
	// Text Текст поста
	Text PostText `json:"text"`
}

// GetPostFeedParams defines parameters for GetPostFeed.
type GetPostFeedParams struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

// PutPostUpdateJSONBody defines parameters for PutPostUpdate.
type PutPostUpdateJSONBody struct {
	// Id Идентификатор поста
	Id PostId `json:"id"`

	// Text Текст поста
	Text PostText `json:"text"`
}

// PostUserRegisterJSONBody defines parameters for PostUserRegister.
type PostUserRegisterJSONBody struct {
	Age        int    `json:"age"`
	Biography  string `json:"biography"`
	City       string `json:"city"`
	FirstName  string `json:"first_name"`
	Password   string `json:"password"`
	SecondName string `json:"second_name"`
	Sex        string `json:"sex"`
}

// GetUserSearchParams defines parameters for GetUserSearch.
type GetUserSearchParams struct {
	// FirstName Условие поиска по имени
	FirstName string `form:"first_name" json:"first_name"`

	// LastName Условие поиска по фамилии
	LastName string `form:"last_name" json:"last_name"`
}

// PostDialogUserIdSendJSONRequestBody defines body for PostDialogUserIdSend for application/json ContentType.
type PostDialogUserIdSendJSONRequestBody PostDialogUserIdSendJSONBody

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody PostLoginJSONBody

// PostPostCreateJSONRequestBody defines body for PostPostCreate for application/json ContentType.
type PostPostCreateJSONRequestBody PostPostCreateJSONBody

// PutPostUpdateJSONRequestBody defines body for PutPostUpdate for application/json ContentType.
type PutPostUpdateJSONRequestBody PutPostUpdateJSONBody

// PostUserRegisterJSONRequestBody defines body for PostUserRegister for application/json ContentType.
type PostUserRegisterJSONRequestBody PostUserRegisterJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /dialog/{user_id}/list)
	GetDialogUserIdList(c *gin.Context, userId UserId)

	// (POST /dialog/{user_id}/send)
	PostDialogUserIdSend(c *gin.Context, userId UserId)

	// (PUT /friend/delete/{user_id})
	PutFriendDeleteUserId(c *gin.Context, userId UserId)

	// (PUT /friend/set/{user_id})
	PutFriendSetUserId(c *gin.Context, userId UserId)

	// (POST /login)
	PostLogin(c *gin.Context)

	// (POST /post/create)
	PostPostCreate(c *gin.Context)

	// (PUT /post/delete/{id})
	PutPostDeleteId(c *gin.Context, id PostId)

	// (GET /post/feed)
	GetPostFeed(c *gin.Context, params GetPostFeedParams)

	// (GET /post/get/{id})
	GetPostGetId(c *gin.Context, id PostId)

	// (PUT /post/update)
	PutPostUpdate(c *gin.Context)

	// (GET /user/get/{id})
	GetUserGetId(c *gin.Context, id UserId)

	// (POST /user/register)
	PostUserRegister(c *gin.Context)

	// (GET /user/search)
	GetUserSearch(c *gin.Context, params GetUserSearchParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetDialogUserIdList operation middleware
func (siw *ServerInterfaceWrapper) GetDialogUserIdList(c *gin.Context) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "user_id", c.Param("user_id"), &userId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetDialogUserIdList(c, userId)
}

// PostDialogUserIdSend operation middleware
func (siw *ServerInterfaceWrapper) PostDialogUserIdSend(c *gin.Context) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "user_id", c.Param("user_id"), &userId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostDialogUserIdSend(c, userId)
}

// PutFriendDeleteUserId operation middleware
func (siw *ServerInterfaceWrapper) PutFriendDeleteUserId(c *gin.Context) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "user_id", c.Param("user_id"), &userId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutFriendDeleteUserId(c, userId)
}

// PutFriendSetUserId operation middleware
func (siw *ServerInterfaceWrapper) PutFriendSetUserId(c *gin.Context) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId UserId

	err = runtime.BindStyledParameter("simple", false, "user_id", c.Param("user_id"), &userId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutFriendSetUserId(c, userId)
}

// PostLogin operation middleware
func (siw *ServerInterfaceWrapper) PostLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostLogin(c)
}

// PostPostCreate operation middleware
func (siw *ServerInterfaceWrapper) PostPostCreate(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostPostCreate(c)
}

// PutPostDeleteId operation middleware
func (siw *ServerInterfaceWrapper) PutPostDeleteId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id PostId

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutPostDeleteId(c, id)
}

// GetPostFeed operation middleware
func (siw *ServerInterfaceWrapper) GetPostFeed(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPostFeedParams

	// ------------- Required query parameter "offset" -------------

	if paramValue := c.Query("offset"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument offset is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", c.Request.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter offset: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "limit" -------------

	if paramValue := c.Query("limit"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument limit is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter limit: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPostFeed(c, params)
}

// GetPostGetId operation middleware
func (siw *ServerInterfaceWrapper) GetPostGetId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id PostId

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPostGetId(c, id)
}

// PutPostUpdate operation middleware
func (siw *ServerInterfaceWrapper) PutPostUpdate(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutPostUpdate(c)
}

// GetUserGetId operation middleware
func (siw *ServerInterfaceWrapper) GetUserGetId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id UserId

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserGetId(c, id)
}

// PostUserRegister operation middleware
func (siw *ServerInterfaceWrapper) PostUserRegister(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUserRegister(c)
}

// GetUserSearch operation middleware
func (siw *ServerInterfaceWrapper) GetUserSearch(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserSearchParams

	// ------------- Required query parameter "first_name" -------------

	if paramValue := c.Query("first_name"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument first_name is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "first_name", c.Request.URL.Query(), &params.FirstName)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter first_name: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "last_name" -------------

	if paramValue := c.Query("last_name"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument last_name is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "last_name", c.Request.URL.Query(), &params.LastName)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter last_name: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserSearch(c, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/dialog/:user_id/list", wrapper.GetDialogUserIdList)
	router.POST(options.BaseURL+"/dialog/:user_id/send", wrapper.PostDialogUserIdSend)
	router.PUT(options.BaseURL+"/friend/delete/:user_id", wrapper.PutFriendDeleteUserId)
	router.PUT(options.BaseURL+"/friend/set/:user_id", wrapper.PutFriendSetUserId)
	router.POST(options.BaseURL+"/login", wrapper.PostLogin)
	router.POST(options.BaseURL+"/post/create", wrapper.PostPostCreate)
	router.PUT(options.BaseURL+"/post/delete/:id", wrapper.PutPostDeleteId)
	router.GET(options.BaseURL+"/post/feed", wrapper.GetPostFeed)
	router.GET(options.BaseURL+"/post/get/:id", wrapper.GetPostGetId)
	router.PUT(options.BaseURL+"/post/update", wrapper.PutPostUpdate)
	router.GET(options.BaseURL+"/user/get/:id", wrapper.GetUserGetId)
	router.POST(options.BaseURL+"/user/register", wrapper.PostUserRegister)
	router.GET(options.BaseURL+"/user/search", wrapper.GetUserSearch)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RaW28bxxX+K4NtH3lZ6mIHeimSGkkLuGjhy0sNw1jtDslJSe56djaVYAgQRceOIaMO",
	"0jwUaZu0TYu+UrQZrSmJ+gtn/lFxziwvSy5vspTUNmCLt9kz536+c2YeWa5fD/wGb6jQ2npkSR4GfiPk",
	"9GHDtvHF46ErRaCE37C2LPg7dKEDbTiBGF7BmT6ELoNX0IazwYcOdKCPX1l7OWvDLs0g0oaOPoC+3ocY",
	"jqFPRA2N10iwr5v6QLfgHKls7uwgFddvKN5Q+NYJgppwHSRY/DREqo+s0K3yuoPvAukHXCphBHF9j2cw",
	"8Q1yyaCvv4AYjqAHcYHBd3ofuigYtOEY/+qn0IUz5OhEv2DQgxNo66ZuQqwfQww9aOsnEEPM4FzvQx+O",
	"4AS6cMrwmyPo06cugyN9SALhkpfQZ7iN/gJJQ4x0Uw8XrJyldgNubVmioXiFS1RCnYehU8mS5Fs4h1g3",
	"SYEx7jYm04hUqKRoVJCS5A8jHqoHwssg9hd4hWzpgzEJyU4MjqFtGMW9llDWHPnhHPrEcw/as6UfsJzw",
	"LCT3rK17Q02kJMkZO98fPuxvf8pdZe3h05MKG+qnzXQTunofOvS3beWsKnc8LslzbnEld/MflhWXGZr6",
	"isQ/1S9yDAWnT8foIv3EsfsoeVc/gy7au40/nukW/ABn6AFNUjS604F+nlKulbP4TlAjty07tZDnxnx7",
	"2jFCtVsjbYl6UONG4Jt+RTRuJeH8BrGj/D9wWsB3HKK+ZfENb41f27bzrsfX8htr7mbecdz1vL39wXrp",
	"WnmtxDc/WGhCQ3cpY32vm3BO0YIe9oJBW7f0wZSPPsFAQnXc4hURKi4vQfoo5DKJkkuWf0D5QhogT3uJ",
	"4UMRNZR9b+AnxPxHQqrqDUdlZYyvKajbjJz0BxPxRCM3JuiaXbqet9fydumObW/Rv4Jt27+3clbZl3VH",
	"WVuWh/QzUswN4dT8ym9GKSut2LL06/j6c8nL1pb1s+KoDhUTEYp3Qy5/7SExxXfUotWpDe/gA/igv+wm",
	"E+Yh/pKNicy0nSZkvJPwOKHof0EXemgnjPc+JjlMBxnKpmQaUxo6yGESaUOPDTLEL7JU/Ds/zNrxO1M6",
	"TYY90c8HxZVC5oR2TdvCiVTVlw/GfH05qyxeixwub0FcbQw3YQxK7okpJpjNMkuy60qV7dwojfL/yCYl",
	"b3N9s+xdy1/fXCvlN9ztUt5xrnl5e720za+X1tbdDW+WZRb6w4w9b/qS15kIwqjOPL/mSxYKxZw6Vznm",
	"YjpzFVeRZI4nAhG6olFhvCZUjoXcY57PuIjCuu8xxeuBL5louMITXtRQLFKs5mz7kjOuDGnO6k6l4TCn",
	"Jh5GToHd5K6KQlZ3IilCFtWUFC4PGZd+yESDuZEMo5CpSAYCV4WhU8gSH11kOuSzoctX0IdjymFNfTCu",
	"iNIHWRhoW/gV6QTV3UzzommpDOumPkxH138IXxxBnGMQpxciVNMHBTjPFMYVKmuzPycF/lV6m78RNuph",
	"vGURKwsZqgcNp84z+T+dzAmDr6YIrRKnIXf9hjdr139DG04hRjw/ufvkT1NchHwnXRrrTo0vLH0UzmOa",
	"SHOYswyyQ9rj9k4skRXwiairBvyM5DjNfMjdSAq1extVa5x5mzuSyw8jVc3Y9stUZ5NUZ9qS4ffQI6Za",
	"uTGoaLqeI30IJwPeWgMw3WFwilWBmqouK9YQ2llJqUdWDTMj1qtKBQZHiEaZKqASiuzz2zt3b7NfiUq1",
	"5jse+1C6VaFQiTnrMy5Dw36psFawUa1+wBtOIKwta71gF2ysG46qkvhFj0pf8VGSh/eKNWGKUYXTCwY+",
	"QSy0i/UJV6ZWGkvdxLVITWJaI6h979EU4hXIDO5o5Szju0PQNO5PSkYpeLxkTExA5vu5dO+7ZnrfpWGj",
	"ULwergRRqCoaezlSOruZwO9riKnV7sNLcgKEarqFsKCjW5gbZjoyBq7pv+1ZbA0FLuKiUa++aG2JOvJl",
	"6GLbTmvXl1w7FmzkE+Nhdu/+3n1cMO17IW9Q9AcJIEo7Hxbjce+7jav/D72PGtmPfG/3Tbq1iwHlyd4M",
	"v8xqTPayo2ROt9JnlOHOqf3tUH+f9L4pLAzdd9RXy1Lwhlf0eI0rPnJZslyU5ayR+pgeuUFPJC7zVuTK",
	"6TYkIy89Z7qVcg/douHhCcIMRGbHDF7pfd2CY+hiSZxVp99pdwm5WsVXbnP1XjhKj8aMbTjB/NGBPg1B",
	"+gOHeTkYOL9zXmEA31iBm0y4ydzQJNNkgm6+ekItTnPOzMxMrs2Pp4wUboa6bf0Uf4tXxtA0+k5hWMK/",
	"I+gL7eGEmGJfP4czjHaah5vhMDH/eXo2xSaODBL2rVxGwb+ZYOTLKaurdFuBE4Z/9OXEzBD+SX3/PuH4",
	"ZIg3GIf3snrFrJ5pSHqF2jzfB9NT4lH0rHjmY57cWCG40eCMzg1eGwNfYYhRDGHoFF3Jk3HobKiI/39p",
	"1v2YqGzm8GtVMLY0j8sM7hbMohMgd2x8YTjRekfTMLnQAMktKMzUdNDSC1blNyjIo6nrxQry99kArfue",
	"WLjMuTdvkoHq/RjXzDarsRZZ9WHE5e7IrH65HHI117QeLztRTVlbdsax5WP9mM4rDxgW9fGTRqqb5gg2",
	"plc6VOzoQ1PKdSs1XLXtnFUXDVGP6rTR7EPFsi/r1l5uKelqoi6WFK40Ld1fIYZTiOn4ow8vqWXEwv+U",
	"jkXa+k/6GcQIbXpYVPKIABnBwGPo0AR5iH/051hdW/oZejDN2LvwOi3/mPilxeL/KOMhOs9ZZio00WCn",
	"oJY+HAYpvh3rpd7lmK1gw5Sk5Hlx+wlXb0E+vrQSvrL3vM0JfuQNUeANYN6c8nw38C4T5P0Ep6GXNqLD",
	"buqUkud7UOSjkMushJHZuQwbWGxXsBz1sPgmSXbGAdJU5sHmcEbmuYwDq6vNXld3ZrJo18XZC41yfkE7",
	"XUGz++VgX2hPN7hmOnWVuY8cWyaXn+aMi/6RdXmIkT4747fzlvFuSqQhl4M7V5eWTpPbAqtcB7jM0/6L",
	"Hukvc3Z/SROiqYP9qz+5v8Ch/baQqpoUWlLuVYyxpm78vVFwX3mMhtyRbnV+6aGrsWOpbFZduW1oLSos",
	"mDZPKJbjQdIcXb41EMAAgHi8oMzu9FKeMK/dS3HxX3PNBxvT4X7DUfAYT+lrKN9AH86S60HTjjzZpq4o",
	"uH48CpVlha85byR7esvF8n+rm9nXuH+CHtXU5dV6VBNk5yOnvpKSfGVRS9lSfjYIrEjWkis2W8Viae16",
	"wS7YhVLR2ru/978AAAD//6A8FUVOMQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
