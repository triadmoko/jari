package user

import (
	"app/internal/features/user/delivery/http/handler"
	"app/internal/features/user/usecase"
	"app/internal/shared/delivery/http/middleware"
	"app/internal/shared/domain/entity"
	"app/internal/shared/domain/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Module struct {
	handler   *handler.UserHandler
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewModule(userRepo repository.UserRepository, logger *logrus.Logger, jwtSecret string) *Module {
	uc := usecase.NewUserUsecase(userRepo, logger)
	h := handler.NewUserHandler(uc)

	return &Module{handler: h, userRepo: userRepo, jwtSecret: jwtSecret}
}

func (m *Module) Name() string {
	return "user"
}

func (m *Module) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		// Protected routes - auth middleware applied inline
		authMW := middleware.AuthMiddleware(m.jwtSecret)
		users.GET("/profile", authMW, m.handler.GetProfile)
		users.PUT("/profile", authMW, m.handler.UpdateProfile)
		// Admin-only: temporary bridge until the permissions & roles module (PRD 03).
		adminMW := middleware.RequireRole(m.userRepo, entity.RoleAdmin)
		users.GET("", authMW, adminMW, m.handler.GetUsers)
	}
}
