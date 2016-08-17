// Code generated by "validate -type=uint8 -output=%s_generated.go"; DO NOT EDIT

package validate

import "text/template"

type Uint8Validator struct {
	errorMessage string
	params       map[string]interface{}
	validate     func(uint8) bool

	constParams map[string]interface{}
	validators  []func(uint8) error
}

func NewUint8Validator(params map[string]interface{}) *Uint8Validator {
	return &Uint8Validator{constParams: params}
}

func (validator *Uint8Validator) Validate(validate func(uint8) bool) *Uint8Validator {
	validator.validate = validate
	return validator
}

func (validator *Uint8Validator) WithParams(params map[string]interface{}) *Uint8Validator {
	validator.params = params
	return validator
}

func (validator *Uint8Validator) WithErrorMessage(errorMessage string) *Uint8Validator {
	validator.errorMessage = errorMessage
	return validator
}
func (validator *Uint8Validator) Add() *Uint8Validator {
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
			func(value uint8) error {
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

func (validator *Uint8Validator) Build() func(uint8) error {
	validator.Add()
	validators := validator.validators
	return func(value uint8) error {
		for _, validate := range validators {
			if err := validate(value); err != nil {
				return err
			}
		}
		return nil
	}
}

func (validator *Uint8Validator) NonZero() *Uint8Validator {
	var zero uint8
	return validator.Add().Validate(func(value uint8) bool { return value != zero }).WithErrorMessage("{.Value} must not be a zero value")
}

func (validator *Uint8Validator) Min(min uint8) *Uint8Validator {
	return validator.Add().Validate(func(value uint8) bool { return value >= min }).
		WithParams(map[string]interface{}{"MinValue": min}).
		WithErrorMessage("{.Value} must be greater than or equal to {.MinValue}")
}

func (validator *Uint8Validator) Max(max uint8) *Uint8Validator {
	return validator.Add().Validate(func(value uint8) bool { return value <= max }).
		WithParams(map[string]interface{}{"MaxValue": max}).
		WithErrorMessage("{.Value} must be less than or equal to {.MaxValue}")
}

func (validator *Uint8Validator) Between(min, max uint8) *Uint8Validator {
	return validator.Add().Validate(func(value uint8) bool { return value >= min && value <= max }).
		WithParams(map[string]interface{}{"MinValue": min, "MaxValue": max}).
		WithErrorMessage("{.Value} must be between {.MinValue} and {.MaxValue}")
}