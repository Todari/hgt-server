package repository

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrPropertyNotFound = errors.New("property not found")
	ErrChatroomNotFound = errors.New("chatroom not found")
	ErrChatNotFound     = errors.New("chat not found")
)

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}
