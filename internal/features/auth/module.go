package auth

import (
	"app/internal/features/auth/delivery/http/handler"
	"app/internal/features/auth/usecase"
	"app/internal/shared/domain/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Module struct {
	handler *handler.AuthHandler
}

func NewModule(userRepo repository.UserRepository, logger *logrus.Logger, jwtSecret string) *Module {
	uc := usecase.NewAuthUsecase(userRepo, logger, jwtSecret)
	h := handler.NewAuthHandler(uc)

	return &Module{handler: h}
}

func (m *Module) Name() string {
	return "auth"
}

func (m *Module) RegisterRoutes(rg *gin.RouterGroup) {
	// Auth routes - all public (no auth required)
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/register", m.handler.Register)
		authGroup.POST("/login", m.handler.Login)
	}
}
