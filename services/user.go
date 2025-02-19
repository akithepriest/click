package services

import (
	"context"

	"github.com/akithepriest/click/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{
		collection: collection,
	}
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*database.User, error) {
	query := &bson.D{{Key: "email", Value: email}}
	var user database.User
	
	err := s.collection.FindOne(ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, database.ErrorNotFound
		}
		return nil, err 
	}
	return &user, nil
}

func (s *UserService) InsertUser(ctx context.Context, name string, email string) (*database.User, error) {
	if _, err := s.GetUserByEmail(ctx, email); err == nil {
		return nil, database.ErrorAlreadyExists
	}
	
	user := &database.User{
		ID: primitive.NewObjectID(),
		Name: name,
		Email: email,
	}
	_, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}