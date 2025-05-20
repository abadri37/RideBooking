package service

import (
	"context"
	"errors"
	"ridebooking/internal/model"
	"ridebooking/internal/repository"
	"ridebooking/internal/utils"
	"time"
)

type UserService interface {
	RegisterUser(ctx context.Context, UserRequest model.UserRequest) error
	LoginUser(ctx context.Context, email, password string) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByUserId(ctx context.Context, userId string) (*model.User, error)
	UpdateUser(ctx context.Context, UserRequest model.UserRequest) error
	RemoveUser(ctx context.Context, emailId string) error
}

type UserServiceImpl struct {
	userRepo           *repository.UserRepository
	driverLocationRepo *repository.DriverLocationRepository
}

func NewUserService(userRepo *repository.UserRepository, locationRepo *repository.DriverLocationRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: userRepo, driverLocationRepo: locationRepo}
}

func (userService *UserServiceImpl) RegisterUser(ctx context.Context, UserRequest model.UserRequest) error {
	hash := utils.HashPassword(UserRequest.Password)
	user := &model.User{
		Email:     UserRequest.Email,
		Password:  hash,
		FirstName: UserRequest.FirstName,
		LastName:  UserRequest.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Type:      UserRequest.Type,
		UserId:    utils.GetUniqueId(),
	}
	err := userService.userRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	if UserRequest.Type == model.Driver {
		driverLocation := &model.DriverLocation{
			DriverId:    user.UserId,
			IsAvailable: true,
			Location:    model.Gelocation{X: 0, Y: 0},
			LastUpdated: time.Now(),
		}
		err := userService.driverLocationRepo.CreateDriverLocation(ctx, driverLocation)
		if err != nil {
			return err
		}
	}
	return nil
}

func (userService *UserServiceImpl) LoginUser(ctx context.Context, email, password string) (string, error) {
	user, err := userService.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if !utils.CheckPassword(user.Password, password) {
		return "", errors.New("invalid login")
	}
	token, err := utils.GenerateJWT(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (userService *UserServiceImpl) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := userService.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService *UserServiceImpl) GetUserByUserId(ctx context.Context, userId string) (*model.User, error) {
	user, err := userService.userRepo.GetUserByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService *UserServiceImpl) UpdateUser(ctx context.Context, UserRequest model.UserRequest) error {
	hash := utils.HashPassword(UserRequest.Password)
	user := &model.User{
		Email:     UserRequest.Email,
		Password:  hash,
		FirstName: UserRequest.FirstName,
		LastName:  UserRequest.LastName,
		UpdatedAt: time.Now(),
	}
	err := userService.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (userService *UserServiceImpl) RemoveUser(ctx context.Context, emailId string) error {
	user, err := userService.userRepo.GetUserByEmail(ctx, emailId)
	if err != nil {
		return err
	}
	err = userService.userRepo.RemoveUserByEmail(ctx, emailId)
	if err != nil {
		return err
	}
	if user.Type == model.Driver {
		err = userService.driverLocationRepo.RemoveDriverLocationById(ctx, user.UserId)
		if err != nil {
			return err
		}
	}
	return nil
}
