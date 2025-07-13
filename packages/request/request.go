package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/usmaarn/locstique_api/internal/config"
	"net/http"
	"reflect"
	"strings"
)

func ParseBody(r *http.Request, target any) any {
	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		fmt.Println("Error parsing json: ", err)
		return errors.New("invalid request body")
	}
	return validateBody(target)
}

func validateBody(body any) any {
	errMap := map[string]string{}

	err := config.Validate.Struct(body)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, fieldError := range validationErrors {
				fieldName := getFieldJsonName(body, fieldError.Field())
				errMap[fieldName] = getErrorMessage(fieldError)
			}
			return errMap
		} else {
			fmt.Println("Error validating body: ", err)
			return errors.New("invalid request body")
		}
	}
	return nil
}

func getErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "field is required"
	case "min":
		return fmt.Sprintf("field must be at least %s characters", fieldError.Param())
	default:
		return "field is not valid"
	}
}

func getFieldJsonName(data any, structField string) string {
	t := reflect.TypeOf(data).Elem()

	field, ok := t.FieldByName(structField)
	if !ok {
		return structField
	}

	tag := field.Tag.Get("json")
	return strings.Split(tag, ",")[0]
}
