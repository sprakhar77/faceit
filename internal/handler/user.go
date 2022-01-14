package handler

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/sprakhar77/faceit/internal/api"
	"github.com/sprakhar77/faceit/internal/core/domain"
	"github.com/sprakhar77/faceit/internal/core/port"
	"github.com/sprakhar77/faceit/internal/errors"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService port.UserService
}

func NewUserHandler(UserService port.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func (hdl *UserHandler) GetByID(request *gin.Context) {
	user, err := hdl.UserService.GetByID(request, request.Param("user_id"))
	if err != nil {
		logAndAbort(request, errors.NewApiError(errors.InvalidInput, err))
		return
	}

	request.JSON(http.StatusOK, user)
}

func (hdl *UserHandler) GetAll(request *gin.Context) {
	params := decodeQueryParams(request)
	users, err := hdl.UserService.GetAll(request, port.GetUsersFilter{
		Country: params.Country,
		Limit:   params.Limit,
		Offset:  params.Offset,
	})
	if err != nil {
		logAndAbort(request, errors.NewApiError(errors.InvalidInput, err))
		return
	}

	request.JSON(http.StatusOK, users)
}

func (hdl *UserHandler) Create(request *gin.Context) {
	body := api.CreateUserRequest{}
	if err := request.BindJSON(&body); err != nil {
		logAndAbort(request, errors.NewApiError(errors.InvalidInput, err))
		return
	}

	if err := body.Validate(); err != nil {
		logAndAbort(request, errors.NewApiError(errors.InvalidInput, err))
		return
	}

	user, err := hdl.UserService.Create(request, createRequestToUser(body))
	if err != nil {
		logAndAbort(request, errors.NewApiError(errors.Internal, err))
		return
	}

	request.JSON(http.StatusCreated, user)
}

func (hdl *UserHandler) Update(request *gin.Context) {
	userID := request.Param("user_id")
	body := api.UpdateUserRequest{}
	if err := request.BindJSON(&body); err != nil {
		logAndAbort(request, errors.NewApiError(errors.InvalidInput, err))
		return
	}

	if err := body.Validate(); err != nil {
		logAndAbort(request, errors.NewApiError(errors.InvalidInput, err))
		return
	}

	err := hdl.UserService.Update(request, userID, updateRequestToUser(body))
	if err != nil {
		logAndAbort(request, errors.NewApiError(errors.Internal, err))
		return
	}

	request.JSON(http.StatusOK, nil)
}

func (hdl *UserHandler) Delete(request *gin.Context) {
	err := hdl.UserService.Delete(request, request.Param("user_id"))
	if err != nil {
		logAndAbort(request, errors.NewApiError(errors.Internal, err))
		return
	}

	request.JSON(http.StatusOK, nil)
}

func logAndAbort(request *gin.Context, err errors.ApiError) {
	logger.Error(err.Message)
	request.AbortWithStatusJSON(err.Code, err)
}

func createRequestToUser(req api.CreateUserRequest) domain.User {
	return domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		NickName:  req.NickName,
		Password:  req.Password,
		Email:     req.Email,
		Country:   req.Country,
	}
}

func updateRequestToUser(req api.UpdateUserRequest) domain.User {
	return domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		NickName:  req.NickName,
		Password:  req.Password,
		Email:     req.Email,
		Country:   req.Country,
	}
}

func decodeQueryParams(request *gin.Context) api.GetAllUsersRequest {
	req := api.GetAllUsersRequest{}
	if country, ok := request.GetQuery("country"); ok {
		req.Country = country
	}

	if limit, ok := request.GetQuery("limit"); ok {
		req.Limit = toUint64Ptr(limit)
	}

	if offset, ok := request.GetQuery("offset"); ok {
		req.Offset = toUint64Ptr(offset)
	}

	return req
}

func toUint64Ptr(s string) *uint64 {
	u, _ := strconv.ParseUint(s, 10, 64)
	return &u
}
