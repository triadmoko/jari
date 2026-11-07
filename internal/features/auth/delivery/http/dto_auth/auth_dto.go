package dto_auth

import (
	"app/internal/shared/constants"
	"app/internal/shared/domain/entity"
	"fmt"
	"time"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (r *RegisterRequest) Validate(lang constants.Lang) map[string][]string {
	errors := make(map[string][]string)

	if r.Email == "" {
		errors["email"] = append(errors["email"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "email"))
	} else if !constants.IsValidEmail(r.Email) {
		errors["email"] = append(errors["email"], constants.GetValidationMessage(constants.InvalidEmail, lang))
	}

	if r.Username == "" {
		errors["username"] = append(errors["username"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "username"))
	} else {
		if !constants.MinLength(r.Username, 3) {
			errors["username"] = append(errors["username"], fmt.Sprintf(constants.GetValidationMessage(constants.UsernameTooShort, lang), 3))
		}
		if !constants.MaxLength(r.Username, 20) {
			errors["username"] = append(errors["username"], fmt.Sprintf(constants.GetValidationMessage(constants.UsernameTooLong, lang), 20))
		}
	}

	if r.Password == "" {
		errors["password"] = append(errors["password"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "password"))
	} else if !constants.MinLength(r.Password, 6) {
		errors["password"] = append(errors["password"], fmt.Sprintf(constants.GetValidationMessage(constants.PasswordTooShort, lang), 6))
	}

	if r.Name == "" {
		errors["name"] = append(errors["name"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "name"))
	}

	return errors
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate(lang constants.Lang) map[string][]string {
	errors := make(map[string][]string)

	if r.Email == "" {
		errors["email"] = append(errors["email"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "email"))
	} else if !constants.IsValidEmail(r.Email) {
		errors["email"] = append(errors["email"], constants.GetValidationMessage(constants.InvalidEmail, lang))
	}

	if r.Password == "" {
		errors["password"] = append(errors["password"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "password"))
	}

	return errors
}

type RegisterResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func ToRegisterResponse(user *entity.User) *RegisterResponse {
	return &RegisterResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Name:      user.Name,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
	}
}

type LoginResponse struct {
	User  *RegisterResponse `json:"user"`
	Token string            `json:"token"`
}
