package validate

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Error       bool        `json:"error"`
	FailedField string      `json:"failed_field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
}

func Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func ParseAndValidate(body []byte, out any) []ErrorResponse {
	err := json.Unmarshal(body, out)
	if err != nil {
		return []ErrorResponse{
			{
				Error:       true,
				FailedField: "body",
				Tag:         "json",
				Value:       err.Error(),
			},
		}
	}

	return Validate(out)
}
