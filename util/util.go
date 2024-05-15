package util

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

func ValidateData[T any](b []byte) (T, error) {
	var resource T

	json.Unmarshal(b, &resource)

	validate := validator.New()

	if err := validate.Struct(resource); err != nil {

		errs := err.(validator.ValidationErrors)
		for _, fieldErr := range errs {
			fmt.Printf("field %s: %s\n", fieldErr.Field(), fieldErr.Tag())
		}

		return resource, errors.New("bad request")
	}

	return resource, nil
}
