package utils

import "github.com/go-playground/validator/v10"

type Validation interface {
	Format(err error) map[string][]string
}

type validation struct{}

// Format implements [Validator].
func (v *validation) Format(err error) map[string][]string {
	errors := make(map[string][]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return errors
	}

	for _, e := range validationErrors {
		field := e.Field() // pastikan sudah pakai RegisterTagNameFunc
		tag := e.Tag()
		param := e.Param()

		var message string

		switch tag {

		case "required":
			message = "field ini wajib diisi"

		case "min":
			if e.Kind().String() == "string" {
				message = "minimal " + param + " karakter"
			} else {
				message = "minimal nilai " + param
			}

		case "max":
			if e.Kind().String() == "string" {
				message = "maksimal " + param + " karakter"
			} else {
				message = "maksimal nilai " + param
			}

		case "email":
			message = "format email tidak valid"

		case "numeric":
			message = "harus berupa angka"

		case "oneof":
			message = "harus salah satu dari: " + param

		case "gte":
			message = "minimal " + param

		case "lte":
			message = "maksimal " + param

		case "len":
			message = "harus tepat " + param + " karakter"

		case "uuid":
			message = "format UUID tidak valid"

		case "url":
			message = "format URL tidak valid"

		case "alphanum":
			message = "hanya boleh huruf dan angka"

		default:
			message = "nilai tidak valid"
		}

		errors[field] = append(errors[field], message)
	}

	return errors
}

func NewValidation() Validation {
	return &validation{}
}
