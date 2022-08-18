package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	validator.FieldError
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	errors := []string{}
	for _, err := range v {
		errors = append(errors, err.Error())
	}
	return errors
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	fmt.Println("Here-1")
	fmt.Printf("interface: %#v", i)
	errs := v.validate.Struct(i)
	if errs == nil {
		return nil
	}
	errp := errs.(validator.ValidationErrors)

	fmt.Println("Here-2")
	if len(errp) == 0 {
		return nil
	}

	fmt.Println("Here")

	var returnErrs []ValidationError

	for _, err := range errp {
		// cast the FieldError into our ValidationError and append to the slice
		ve := ValidationError{err}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs

}

func validateSKU(fl validator.FieldLevel) bool {
	exp := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	match := exp.FindAllString(fl.Field().String(), -1)

	return len(match) == 1
}
