package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/usmaarn/locstique_api/internal/dto"
	"github.com/usmaarn/locstique_api/packages/helpers"
	"github.com/usmaarn/locstique_api/packages/request"
	"github.com/usmaarn/locstique_api/packages/response"
	"net/http"
)

func (h *Handler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var requestDto dto.CreateUserDto

	err := request.ParseBody(r, &requestDto)
	if err != nil {
		var err validator.ValidationErrors
		if errors.As(err, &err) {
			errMap := helpers.FormatValidationErrors(err)
			response.Error(w, 400, errMap)
			return
		}
		response.Error(w, 400, err)
		return
	}

	isEmailExists := h.service.UserExistsByEmail(requestDto.Email)
	if isEmailExists {
		response.Error(w, 400, map[string]string{"email": "Email already exists"})
		return
	}

	user, err := h.service.CreateUser(requestDto)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		response.Error(w, 500, "An error occurred!")
		return
	}

	token, err := h.service.CreateToken(user.ID)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		response.Error(w, 500, "An error occurred!")
		return
	}

	response.Success(w, map[string]string{"token": token.ID})
}

// LoginHandler Authentication -> Login
func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var requestDto dto.LoginDto
	err := request.ParseBody(r, &requestDto)
	if err != nil {
		response.Error(w, 400, err)
		return
	}

	user, err := h.service.GetUserByEmail(requestDto.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(w, 400, "incorrect email or password")
			return
		}
		fmt.Println("Error fetching user: ", err)
		response.Error(w, 500, "an error occurred")
		return
	}
}
