package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Todari/go-mongo-crud-rest-api/model"
	"github.com/Todari/go-mongo-crud-rest-api/repository"
)

type Server struct {
	repository repository.Repository
}

func NewServer(repository repository.Repository) *Server {
	return &Server{repository: repository}
}

func (s Server) GetUser(ctx *gin.Context) {
	studentId := ctx.Param("studentId")
	if studentId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument studentId"})
		return
	}
	user, err := s.repository.GetUser(ctx, studentId)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	user, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) UpdateUser(ctx *gin.Context) {
	studentId := ctx.Param("studentId")
	if studentId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument studentId"})
		return
	}
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	user.StudentId = studentId
	user, err := s.repository.UpdateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) DeleteUser(ctx *gin.Context) {
	studentId := ctx.Param("studentId")
	if studentId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument studentId"})
		return
	}
	if err := s.repository.DeleteUser(ctx, studentId); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
