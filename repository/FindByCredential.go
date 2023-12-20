package repository

import (
	"errors"

	"github.com/force445/basic-jwt-auth/models"
)

func FindByCredentials(email, password string) (*models.User, error) {
	if email == "test@mail.com" && password == "test12345" {
		return &models.User{
			ID:              1,
			Email:           "test@mail.com",
			Password:        "test12345",
			FavouritePhrase: "Hello, World",
		}, nil
	}
	return nil, errors.New("user not found")
}
