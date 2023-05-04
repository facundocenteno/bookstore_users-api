package services

import (
	"github.com/facundocenteno/bookstore_users-api/domain/users"
	"github.com/facundocenteno/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}

func FindUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{
		Id: userId,
	}

	if err := users.GetById(); err != nil {
		return nil, err
	}

	return result, nil

}
