package validate

func (validator *BoolValidator) True() *BoolValidator {
	return validator.Add().Validate(func(value bool) bool { return value }).WithErrorMessage("must be true")
}

func (validator *BoolValidator) False() *BoolValidator {
	return validator.Add().Validate(func(value bool) bool { return !value }).WithErrorMessage("must be false")
}
