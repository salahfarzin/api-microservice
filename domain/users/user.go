package users

import "github.com/salahfarzin/api-microservice/domain/errors"

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (user User) Validate() *errors.HttpError {
	if user.FirstName == "" {
		return errors.NewBadRequestError("invalid first name")
	}

	if user.LastName == "" {
		return errors.NewBadRequestError("invalid last name")
	}

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}
