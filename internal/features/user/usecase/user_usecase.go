package usecase

import (
	"app/internal/features/user/delivery/http/dto_user"
	"app/internal/shared/constants"
	"app/internal/shared/delivery/http/middleware"
	"app/internal/shared/domain/entity"
	"app/internal/shared/domain/repository"
	"app/pkg"
	"context"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type UserUsecase interface {
	GetProfile(ctx context.Context, userID string) (*dto_user.UserResponse, int, error)
	UpdateProfile(ctx context.Context, userID string, req *dto_user.UpdateProfileRequest) (*dto_user.UserResponse, int, error)
	GetUsers(ctx context.Context, queries map[string]string) ([]*dto_user.UserResponse, pkg.PaginationResponse, int, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
	logger   *logrus.Logger
}

func NewUserUsecase(userRepo repository.UserRepository, logger *logrus.Logger) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (u *userUsecase) GetProfile(ctx context.Context, userID string) (*dto_user.UserResponse, int, error) {
	lang := middleware.GetLangFromContext(ctx)

	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		u.logger.WithContext(ctx).Error("u.userRepo.GetByID ", err)
		return nil, http.StatusNotFound, constants.GetError(constants.UserNotFound, lang)
	}

	return dto_user.ToUserResponse(user), http.StatusOK, nil
}

func (u *userUsecase) UpdateProfile(ctx context.Context, userID string, req *dto_user.UpdateProfileRequest) (*dto_user.UserResponse, int, error) {
	lang := middleware.GetLangFromContext(ctx)

	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		u.logger.WithContext(ctx).Error("u.userRepo.GetByID ", err)
		return nil, http.StatusNotFound, constants.GetError(constants.UserNotFound, lang)
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	filter := entity.FilterUser{
		ID: userID,
	}

	if err := u.userRepo.Update(ctx, filter, user); err != nil {
		u.logger.WithContext(ctx).Error("u.userRepo.Update ", err)
		return nil, http.StatusInternalServerError, constants.GetError(constants.FailedToUpdateUser, lang)
	}

	return dto_user.ToUserResponse(user), http.StatusOK, nil
}

func (u *userUsecase) GetUsers(ctx context.Context, queries map[string]string) ([]*dto_user.UserResponse, pkg.PaginationResponse, int, error) {
	lang := middleware.GetLangFromContext(ctx)

	pagination := pkg.PaginationBuilder(queries["per_page"], queries["page"])

	var roles []string
	if queries["roles"] != "" {
		roles = strings.Split(queries["roles"], ",")
	}

	filter := entity.FilterUser{
		ID:       queries["id"],
		Email:    queries["email"],
		Username: queries["username"],
		Name:     queries["name"],
		Role:     queries["role"],
		Roles:    roles,
		PerPage:  pagination.PerPage,
		Offset:   pagination.Offset,
	}

	users, total, err := u.userRepo.List(ctx, filter)
	if err != nil {
		u.logger.WithContext(ctx).Error("u.userRepo.List ", err)
		return nil, pkg.PaginationResponse{}, http.StatusInternalServerError, constants.GetError(constants.FailedToGetUsers, lang)
	}

	userResponses := make([]*dto_user.UserResponse, 0, len(users))
	for _, user := range users {
		userResponses = append(userResponses, dto_user.ToUserResponse(user))
	}

	totalPage := pkg.TotalPage(total, pagination.PerPage)
	paginationResponse := pkg.PaginationResponse{
		PerPage:   pagination.PerPage,
		TotalPage: totalPage,
		TotalData: total,
		Page:      pagination.Page,
	}

	return userResponses, paginationResponse, http.StatusOK, nil
}
