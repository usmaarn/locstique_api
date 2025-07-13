package helpers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func FormatValidationErrors(errs validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)

	for _, e := range errs {
		// Get the struct type and reflect field
		field, found := reflect.TypeOf(e.StructField()).FieldByName(e.StructField())
		if !found {
			// fallback to StructField name if not found
			errors[e.StructField()] = "Invalid value"
			continue
		}

		// Extract JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			jsonTag = e.StructField()
		}

		// Strip tag options like `omitempty`
		if commaIdx := len(jsonTag); commaIdx > 0 {
			if comma := len(jsonTag); comma >= 0 {
				jsonTag = jsonTag[:comma]
			}
		}

		// Customize error messages per tag
		var msg string
		switch e.Tag() {
		case "required":
			msg = "This field is required"
		case "email":
			msg = "Invalid email format"
		case "min":
			msg = fmt.Sprintf("Minimum length is %s", e.Param())
		default:
			msg = "Invalid value"
		}

		errors[jsonTag] = msg
	}

	return errors
}
