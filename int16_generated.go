// Code generated by "validate -type=int16 -output=%s_generated.go"; DO NOT EDIT

package validate

import "text/template"

type Int16Validator struct {
	errorMessage string
	params       map[string]interface{}
	validate     func(int16) bool

	constParams map[string]interface{}
	validators  []func(int16) error
}

func NewInt16Validator(params map[string]interface{}) *Int16Validator {
	return &Int16Validator{constParams: params}
}

func (validator *Int16Validator) Validate(validate func(int16) bool) *Int16Validator {
	validator.validate = validate
	return validator
}

func (validator *Int16Validator) WithParams(params map[string]interface{}) *Int16Validator {
	validator.params = params
	return validator
}

func (validator *Int16Validator) WithErrorMessage(errorMessage string) *Int16Validator {
	validator.errorMessage = errorMessage
	return validator
}
func (validator *Int16Validator) Add() *Int16Validator {
	if validator.validate != nil {
		tmpl := template.Must(template.New("errorMessage").Parse(validator.errorMessage))
		params := make(map[string]interface{})
		for key, value := range validator.constParams {
			params[key] = value
		}
		for key, value := range validator.params {
			params[key] = value
		}
		isValid := validator.validate
		validator.validators = append(validator.validators,
			func(value int16) error {
				if isValid(value) {
					return nil
				}
				return FormatErrorMessage(tmpl, params, value)
			})
	}
	validator.errorMessage = ""
	validator.params = nil
	return validator
}

func (validator *Int16Validator) Build() func(int16) error {
	validator.Add()
	validators := validator.validators
	return func(value int16) error {
		for _, validate := range validators {
			if err := validate(value); err != nil {
				return err
			}
		}
		return nil
	}
}

func (validator *Int16Validator) NonZero() *Int16Validator {
	var zero int16
	return validator.Add().Validate(func(value int16) bool { return value != zero }).WithErrorMessage("{.Value} must not be a zero value")
}

func (validator *Int16Validator) Min(min int16) *Int16Validator {
	return validator.Add().Validate(func(value int16) bool { return value >= min }).
		WithParams(map[string]interface{}{"MinValue": min}).
		WithErrorMessage("{.Value} must be greater than or equal to {.MinValue}")
}

func (validator *Int16Validator) Max(max int16) *Int16Validator {
	return validator.Add().Validate(func(value int16) bool { return value <= max }).
		WithParams(map[string]interface{}{"MaxValue": max}).
		WithErrorMessage("{.Value} must be less than or equal to {.MaxValue}")
}

func (validator *Int16Validator) Between(min, max int16) *Int16Validator {
	return validator.Add().Validate(func(value int16) bool { return value >= min && value <= max }).
		WithParams(map[string]interface{}{"MinValue": min, "MaxValue": max}).
		WithErrorMessage("{.Value} must be between {.MinValue} and {.MaxValue}")
}
