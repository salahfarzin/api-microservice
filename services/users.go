package services

import (
	"fmt"

	"github.com/salahfarzin/api-microservice/domain/errors"
	"github.com/salahfarzin/api-microservice/domain/users"
)

type usersService interface {
	Create(user users.User) (*users.User, *errors.HttpError)
	Get(userId int64) (*users.User, *errors.HttpError)
}

type usersServiceImpl struct{}

var (
	UsersService    usersService = usersServiceImpl{}
	registeredUsers              = map[int64]*users.User{}
	currentUserId   int64        = 1
)

// create a user and put into registered users map
func (service usersServiceImpl) Create(user users.User) (*users.User, *errors.HttpError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Id = currentUserId
	currentUserId++

	registeredUsers[user.Id] = &user

	return &user, nil
}

//return user by id from registered users map
func (service usersServiceImpl) Get(userId int64) (*users.User, *errors.HttpError) {
	if user := registeredUsers[userId]; user != nil {
		return user, nil
	}

	return nil, errors.NewNotFoundError(fmt.Sprintf("user %d not found", userId))
}
