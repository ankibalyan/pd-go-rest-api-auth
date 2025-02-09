package verifier

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func registerValidations() {
	registerStrictPasswordTag()
}

func registerStrictPasswordTag() {
	strictPasswordTag := "strictPassword"

	verifier.RegisterValidation(strictPasswordTag, func(fl validator.FieldLevel) bool {
		fieldValue := fl.Field().String()

		return !regexp.MustCompile(`^([^0-9]+|[^a-z]+|[^A-Z]+|[0-9a-zA-Z]+)$`).MatchString(fieldValue)
	})

	verifier.RegisterTranslation(strictPasswordTag, translator, func(ut ut.Translator) error {
		return ut.Add(strictPasswordTag, `{0} must contain at least a digit, a small case letter, a capital letter and a special character`, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		transalatedMsg, _ := ut.T(strictPasswordTag, fe.Field())

		return transalatedMsg
	})
}
