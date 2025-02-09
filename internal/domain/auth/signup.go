package auth

import (
	"net/http"
	"pdauth/pkg/httputils"
	"pdauth/pkg/randomness"
	"pdauth/pkg/verifier"
)

type SingupRequest struct {
	Name            string `json:"name" validate:"required,min=3"`
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Gender          string `json:"gender"`
	DateOfBirth     string `json:"dateOfBirth" validate:"required"`
	Password        string `json:"password" validate:"required,min=7,max=20,strictPassword=medium"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type User struct {
	ID          string
	Name        string
	Username    string
	Email       string
	Gender      string
	DateOfBirth string
	Password    string
}

func SignupHanlder(w http.ResponseWriter, r *http.Request) {
	// read the request body
	var input SingupRequest
	err := httputils.ParseRequestBody(r, &input)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// validate the request body
	validationErros := verifier.ValidateStruct(input)

	if validationErros != nil {
		result := map[string]any{
			"code": "VALIDATION_ERROR",
			"data": validationErros.Details,
			"msg":  validationErros.Msg,
		}

		httputils.SendJSONResponse(w, result)

		return
	}
	// validations for signup
	// create the user

	user := User{
		ID:          randomness.NewRandomID(),
		Name:        input.Name,
		Username:    input.Username,
		Email:       input.Email,
		Gender:      input.Gender,
		DateOfBirth: input.DateOfBirth,
		Password:    input.Password,
	}

	result := map[string]any{
		"code": "SUCCESS",
		"data": user,
		"msg":  "user created",
	}

	// response
	httputils.SendJSONResponse(w, result)
}
