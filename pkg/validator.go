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

const TYPE_DEFAULT = "type"
const TYPE_BOOL = "bool"
const TYPE_COMPLEX = "complex"
const TYPE_INT = "int"
const TYPE_UINT = "uint"
const TYPE_FLOAT = "float"

type RequestValidator struct {
}

func NewRequestValidator() RequestValidator {
	return RequestValidator{}
}

type Messages map[string]string

func ValidateRequest[T any](c echo.Context, request T) (T, error) {
	err := c.Bind(request)

	// check if err is &echo.HTTPError and Internal strcon.NumError
	if _, ok := err.(*echo.HTTPError); ok {
		if e, ok := err.(*echo.HTTPError).Internal.(*strconv.NumError); ok {
			// check for ParseBool, ParseInt, ParseUint, ParseFloat, ParseComplex
			// and use it's type for msg e.g. ParseBool -> bool
			t := strings.ToLower(strings.Replace(e.Func, "Parse", "", 1))
			if t != TYPE_BOOL &&
				t != TYPE_COMPLEX &&
				t != TYPE_INT &&
				t != TYPE_UINT &&
				t != TYPE_FLOAT {
				t = TYPE_DEFAULT
			}
			msg := fmt.Sprintf("%s is not a valid %s", e.Num, t)
			return request, echo.NewHTTPError(http.StatusBadRequest, msg)
		}
	}
	if err != nil {
		return request, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	msgs, err := validatorMessages(request)
	if err != nil {
		return request, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}
	return request, nil
}

func validatorMessages[T any](request T) (Messages, error) {
	errs := validator.New().Struct(request)
	msgs := Messages{}
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(request).Elem().FieldByName(err.Field())
			q := getStructTag(field, "query")
			msg := getStructTag(field, "message")
			msgs[q] = msg
		}
		return msgs, errs
	}
	return nil, nil
}

// getStructTag returns the value of the tag with the given name
func getStructTag(f reflect.StructField, tagName string) string {
	return string(f.Tag.Get(tagName))
}
