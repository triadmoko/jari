package user

import (
	"app/internal/features/user/delivery/http/handler"
	"app/internal/features/user/usecase"
	"app/internal/shared/delivery/http/middleware"
	"app/internal/shared/domain/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Module is the user feature module that combines DI and route registration
type Module struct {
	handler   *handler.UserHandler
	jwtSecret string
}

// NewModule creates and wires all user feature dependencies
func NewModule(userRepo repository.UserRepository, logger *logrus.Logger, jwtSecret string) *Module {
	// Wire dependencies
	uc := usecase.NewUserUsecase(userRepo, logger)
	h := handler.NewUserHandler(uc)

	return &Module{handler: h, jwtSecret: jwtSecret}
}

// Name returns the feature name
func (m *Module) Name() string {
	return "user"
}

// RegisterRoutes registers all user routes
func (m *Module) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		// Protected routes - auth middleware applied inline
		authMW := middleware.AuthMiddleware(m.jwtSecret)
		users.GET("/profile", authMW, m.handler.GetProfile)
		users.PUT("/profile", authMW, m.handler.UpdateProfile)
		users.GET("", authMW, m.handler.GetUsers)
	}
}
