package usecase

import (
	"app/internal/features/auth/delivery/http/dto_auth"
	"app/internal/shared/constants"
	"app/internal/shared/delivery/http/middleware"
	"app/internal/shared/domain/entity"
	"app/internal/shared/domain/repository"
	"app/pkg/crypto"
	"app/pkg/jwt"
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type AuthUsecase interface {
	Register(ctx context.Context, req dto_auth.RegisterRequest) (*dto_auth.RegisterResponse, int, error)
	Login(ctx context.Context, req dto_auth.LoginRequest) (*dto_auth.LoginResponse, int, error)
}

type authUsecase struct {
	userRepo  repository.UserRepository
	logger    *logrus.Logger
	jwtSecret string
}

func NewAuthUsecase(userRepo repository.UserRepository, logger *logrus.Logger, jwtSecret string) AuthUsecase {
	return &authUsecase{
		userRepo:  userRepo,
		logger:    logger,
		jwtSecret: jwtSecret,
	}
}

func (a *authUsecase) Register(ctx context.Context, req dto_auth.RegisterRequest) (*dto_auth.RegisterResponse, int, error) {
	lang := middleware.GetLangFromContext(ctx)

	existingUser, _ := a.userRepo.GetByEmail(ctx, req.Email)
	if existingUser != nil {
		a.logger.WithContext(ctx).Error("a.userRepo.GetByEmail: user already exists")
		return nil, http.StatusBadRequest, constants.GetError(constants.UserAlreadyExists, lang)
	}

	existingUser, _ = a.userRepo.GetByUsername(ctx, req.Username)
	if existingUser != nil {
		a.logger.WithContext(ctx).Error("a.userRepo.GetByUsername: username already taken")
		return nil, http.StatusBadRequest, constants.GetError(constants.UsernameAlreadyTaken, lang)
	}

	hashedPassword, err := crypto.HashPassword(req.Password)
	if err != nil {
		a.logger.WithContext(ctx).Error("crypto.HashPassword ", err)
		return nil, http.StatusInternalServerError, constants.GetError(constants.FailedToHashPassword, lang)
	}

	user := entity.NewUser(req.Email, req.Username, hashedPassword, req.Name)

	if err := a.userRepo.Create(ctx, user); err != nil {
		a.logger.WithContext(ctx).Error("a.userRepo.Create ", err)
		return nil, http.StatusInternalServerError, constants.GetError(constants.FailedToCreateUser, lang)
	}

	return dto_auth.ToRegisterResponse(user), http.StatusCreated, nil
}

func (a *authUsecase) Login(ctx context.Context, req dto_auth.LoginRequest) (*dto_auth.LoginResponse, int, error) {
	lang := middleware.GetLangFromContext(ctx)

	user, err := a.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		a.logger.WithContext(ctx).Error("a.userRepo.GetByEmail ", err)
		return nil, http.StatusUnauthorized, constants.GetError(constants.InvalidCredentials, lang)
	}

	if err := crypto.VerifyPassword(user.Password, req.Password); err != nil {
		a.logger.WithContext(ctx).Error("crypto.VerifyPassword ", err)
		return nil, http.StatusUnauthorized, constants.GetError(constants.InvalidCredentials, lang)
	}

	// Generate token with string UUID
	token, err := jwt.GenerateToken(a.jwtSecret, jwt.UserPayload{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	})
	if err != nil {
		a.logger.WithContext(ctx).Error("jwt.GenerateToken ", err)
		return nil, http.StatusInternalServerError, constants.GetError(constants.FailedToGenerateToken, lang)
	}

	return &dto_auth.LoginResponse{
		User:  dto_auth.ToRegisterResponse(user),
		Token: token,
	}, http.StatusOK, nil
}
