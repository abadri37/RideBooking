package repository

import (
	"context"
	"ridebooking/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		Collection: collection,
	}
}

func (userRepo *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	_, err := userRepo.Collection.InsertOne(ctx, user)
	return err
}

func (userRepo *UserRepository) GetUserByUserId(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	err := userRepo.Collection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := userRepo.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo *UserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	var fetchUser model.User
	err := userRepo.Collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&fetchUser)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
			"password":   user.Password,
			"updated_at": time.Now(),
		},
	}
	_, error := userRepo.Collection.UpdateByID(ctx, fetchUser.Id, update)
	if error != nil {
		return error
	}
	return nil
}

func (userRepo *UserRepository) RemoveUserByEmail(ctx context.Context, emailId string) error {
	var fetchUser model.User
	err := userRepo.Collection.FindOne(ctx, bson.M{"email": emailId}).Decode(&fetchUser)
	if err != nil {
		return err
	}
	_, error := userRepo.Collection.DeleteOne(ctx, bson.M{"_id": fetchUser.Id})
	if error != nil {
		return error
	}
	return nil
}
