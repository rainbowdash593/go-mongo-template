package v1

import (
	"errors"
	"example/template/internal/domain/dto"
	"example/template/internal/domain/exceptions"
	"example/template/internal/usecase"
	logging "example/template/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type userRoutes struct {
	u usecase.UserUseCase
	l *log.Logger
}

func newUserRoutes(handler *gin.RouterGroup, u usecase.UserUseCase) {
	l := logging.GetLogger()
	r := &userRoutes{u, l}

	h := handler.Group("/users")
	{
		h.GET("/:id", r.find)
		h.POST("/", r.create)
	}
}

func (r *userRoutes) find(c *gin.Context) {
	user, err := r.u.FindUserByID(c.Request.Context(), c.Param("id"))
	if errors.Is(err, exceptions.ErrUserNotFound) {
		errorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	if err != nil {
		r.l.Error(fmt.Errorf("find user error: %w", err))
		errorResponse(c, http.StatusInternalServerError, "unhandled exceptions")
		return
	}
	c.JSON(http.StatusOK, userResponse{user})
}

func (r *userRoutes) create(c *gin.Context) {
	var request createUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	user, err := r.u.CreateUser(c.Request.Context(), dto.CreateUser{
		Name:  request.Name,
		Email: request.Email,
	})

	if errors.Is(err, exceptions.ErrUserAlreadyExists) {
		errorResponse(c, http.StatusConflict, "user already exists")
		return
	}

	if err != nil {
		r.l.Error(fmt.Errorf("create user error: %w", err))
		errorResponse(c, http.StatusInternalServerError, "unhandled exceptions")
		return
	}

	c.JSON(http.StatusOK, userResponse{user})
}
