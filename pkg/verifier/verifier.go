package verifier

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	englishTranslations "github.com/go-playground/validator/v10/translations/en"
)

var verifier = validator.New(validator.WithRequiredStructEnabled())
var translator ut.Translator

func init() {
	engilsh := en.New()
	uni := ut.New(engilsh, engilsh)

	var found bool
	translator, found = uni.GetTranslator("en")

	if !found {
		panic("no translator found for english language.")
	}

	err := englishTranslations.RegisterDefaultTranslations(verifier, translator)
	if err != nil {
		panic("can not register defail engilsh translations to the validator.")
	}

	// register validations
	registerValidations()
}

type ErrValidation struct {
	Msg     string
	Details map[string]string
	Raw     error
}

func (ev *ErrValidation) Error() string {
	return ev.Msg
}

func ValidateStruct(data any) *ErrValidation {
	err := verifier.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return &ErrValidation{
				Msg: "invalid request body",
				Raw: err,
			}
		}

		errs, ok := err.(validator.ValidationErrors)
		if ok {
			var validationErrors = make(map[string]string, len(errs))
			for _, e := range errs {
				fmt.Println(e.Field(), e.Error())
				validationErrors[e.Field()] = e.Translate(translator)
			}

			return &ErrValidation{
				Msg:     "invalid request body",
				Details: validationErrors,
			}
		}

		return &ErrValidation{
			Msg: "invalid request body",
			Raw: err,
		}
	}

	return nil
}
