package utility

import (
	"fmt"
	"workspace_booking/model"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Input string
	Value string
}

var validate = validator.New()

func ValidateUserStruct(user model.User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Input = err.Field()
			fmt.Println(err.Tag())
			element.Value = msgForTag(err.Tag(), err.Field())
			errors = append(errors, &element)
		}
	}
	return errors
}

func msgForTag(tag, name string) string {
	switch tag {
	case "required":
		return name + " is required"
	case "email":
		return "Invalid email"
	case "max":
		return name + " must have 6 characters long"
	}
	return ""
}
