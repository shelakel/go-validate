// Code generated by "validate -type=time.Time -name=Time -imports=time -output=%s_generated.go"; DO NOT EDIT

package validate

import (
	"text/template"
	"time"
)

type TimeValidator struct {
	errorMessage string
	params       map[string]interface{}
	validate     func(time.Time) bool

	constParams map[string]interface{}
	validators  []func(time.Time) error
}

func NewTimeValidator(params map[string]interface{}) *TimeValidator {
	return &TimeValidator{constParams: params}
}

func (validator *TimeValidator) Validate(validate func(time.Time) bool) *TimeValidator {
	validator.validate = validate
	return validator
}

func (validator *TimeValidator) WithParams(params map[string]interface{}) *TimeValidator {
	validator.params = params
	return validator
}

func (validator *TimeValidator) WithErrorMessage(errorMessage string) *TimeValidator {
	validator.errorMessage = errorMessage
	return validator
}
func (validator *TimeValidator) Add() *TimeValidator {
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
			func(value time.Time) error {
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

func (validator *TimeValidator) Build() func(time.Time) error {
	validator.Add()
	validators := validator.validators
	return func(value time.Time) error {
		for _, validate := range validators {
			if err := validate(value); err != nil {
				return err
			}
		}
		return nil
	}
}

func (validator *TimeValidator) NonZero() *TimeValidator {
	var zero time.Time
	return validator.Add().Validate(func(value time.Time) bool { return value != zero }).WithErrorMessage("{.Value} must not be a zero value")
}
