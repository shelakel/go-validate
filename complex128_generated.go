// Code generated by "validate -type=complex128 -output=%s_generated.go"; DO NOT EDIT

package validate

import "text/template"

type Complex128Validator struct {
	errorMessage string
	params       map[string]interface{}
	validate     func(complex128) bool

	constParams map[string]interface{}
	validators  []func(complex128) error
}

func NewComplex128Validator(params map[string]interface{}) *Complex128Validator {
	return &Complex128Validator{constParams: params}
}

func (validator *Complex128Validator) Validate(validate func(complex128) bool) *Complex128Validator {
	validator.validate = validate
	return validator
}

func (validator *Complex128Validator) WithParams(params map[string]interface{}) *Complex128Validator {
	validator.params = params
	return validator
}

func (validator *Complex128Validator) WithErrorMessage(errorMessage string) *Complex128Validator {
	validator.errorMessage = errorMessage
	return validator
}
func (validator *Complex128Validator) Add() *Complex128Validator {
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
			func(value complex128) error {
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

func (validator *Complex128Validator) Build() func(complex128) error {
	validator.Add()
	validators := validator.validators
	return func(value complex128) error {
		for _, validate := range validators {
			if err := validate(value); err != nil {
				return err
			}
		}
		return nil
	}
}

func (validator *Complex128Validator) NonZero() *Complex128Validator {
	var zero complex128
	return validator.Add().Validate(func(value complex128) bool { return value != zero }).WithErrorMessage("{.Value} must not be a zero value")
}
