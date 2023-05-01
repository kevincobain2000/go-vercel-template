package pkg

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RequestValidator struct {
}

func NewRequestValidator() RequestValidator {
	return RequestValidator{}
}

func ValidateRequest[T any](c echo.Context, request T) (T, error) {
	err := c.Bind(request)

	// check if err is &echo.HTTPError and Internal strcon.NumError
	if _, ok := err.(*echo.HTTPError); ok {
		if e, ok := err.(*echo.HTTPError).Internal.(*strconv.NumError); ok {
			// check for ParseBool, ParseInt, ParseUint, ParseFloat, ParseComplex
			// and use it's type for msg e.g. ParseBool -> bool
			t := strings.ToLower(strings.Replace(e.Func, "Parse", "", 1))
			if t != "bool" && t != "complex" && t != "int" && t != "uint" && t != "float" {
				t = "type"
			}
			msg := fmt.Sprintf("%s is not a valid %s", e.Num, t)
			return request, echo.NewHTTPError(http.StatusBadRequest, msg)
		}
	} else if err != nil {
		return request, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	msgs, err := validatorMessages(request)
	if err != nil {
		return request, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}
	return request, nil
}

func validatorMessages[T any](request T) (map[string]string, error) {
	errs := validator.New().Struct(request)
	msgs := make(map[string]string)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(request).Elem().FieldByName(err.Field())
			queryTag := getStructTag(field, "query")
			message := getStructTag(field, "message")
			msgs[queryTag] = message
		}
		return msgs, errs
	}
	return nil, nil
}

// getStructTag returns the value of the tag with the given name
func getStructTag(f reflect.StructField, tagName string) string {
	return string(f.Tag.Get(tagName))
}
