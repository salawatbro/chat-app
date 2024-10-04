package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/pkg/constants"
	"reflect"
	"strings"
)

func ExtractValidationError(req interface{}) error {
	var message string
	var v = validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := v.Struct(req)
	if err != nil {
		for i, err := range err.(validator.ValidationErrors) {
			if i > 0 {
				message += " | "
			}

			if err.Tag() == "exists" {
				message += err.Field() + ": not exists"
			} else {
				message += err.Field() + ": " + err.Tag()
			}
		}

		return errors.New(message)
	}

	return nil
}

func ValidateIdParams(ctx *fiber.Ctx) (string, error) {
	id := ctx.Params("id")
	if id == "" {
		return "", JsonErrorValidation(ctx, constants.ErrIdRequired)
	}
	return id, nil
}
