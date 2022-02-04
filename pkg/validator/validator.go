package validator

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code

		// return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
		// 	"error": err.Error(),
		// })
		var errMsg string
		for _, er := range err.(validator.ValidationErrors) {
			switch er.Tag() {
			case "required":
				errMsg += fmt.Sprintf("%s is required. ", er.Field())
			}
		}

		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": errMsg,
		})
	}
	return nil
}

func New() *Validator {
	val := validator.New()
	return &Validator{val}
}
