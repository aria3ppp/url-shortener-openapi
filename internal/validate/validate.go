package validate

import (
	"github.com/aria3ppp/url-shortener-openapi/internal/oapi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func CreateLinkRequestBody(r oapi.CreateLinkRequestBody) error {
	return validation.ValidateStruct(
		&r,
		validation.Field(
			&r.Url,
			validation.Required,
			is.URL,
		),
		validation.Field(
			&r.ShortenedString,
			validation.When(
				r.ShortenedString != nil,
				validation.Required,
				is.Alphanumeric,
				validation.Length(6, 32),
			),
		),
	)
}

func CreateUserRequestBody(r oapi.CreateUserRequestBody) error {
	return validation.ValidateStruct(
		&r,
		validation.Field(
			&r.Username,
			validation.Required,
			validation.Length(8, 40),
		),
		validation.Field(
			&r.Password,
			validation.Required,
			validation.Length(8, 40),
		),
	)
}
