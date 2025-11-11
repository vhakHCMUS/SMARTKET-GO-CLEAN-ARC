package routes

import (
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/api/controllers"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/api/middlewares"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/lib"
	handlers "github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/presentation/http"
)

// AuthRoutes struct
type AuthRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	authController controllers.JWTAuthController
	authHandler    *handlers.AuthHandler
}

// Setup user routes
func (s AuthRoutes) Setup() {
	s.logger.Info("Setting up auth routes")

	// Old auth endpoints (backward compatibility)
	auth := s.handler.Gin.Group("/auth")
	{
		auth.POST("/login", s.authController.SignIn)
		auth.POST("/register", s.authController.Register)
	}

	// New MVP auth endpoints
	api := s.handler.Gin.Group("/api/auth")
	{
		api.POST("/register", s.authHandler.Register)
		api.POST("/login", s.authHandler.Login)

		// Protected routes
		protected := api.Group("")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.POST("/logout", s.authHandler.Logout)
			protected.GET("/profile", s.authHandler.GetProfile)
			protected.PUT("/profile", s.authHandler.UpdateProfile)
		}
	}
}

// NewAuthRoutes creates new user controller
func NewAuthRoutes(
	handler lib.RequestHandler,
	authController controllers.JWTAuthController,
	authHandler *handlers.AuthHandler,
	logger lib.Logger,
) AuthRoutes {
	return AuthRoutes{
		handler:        handler,
		logger:         logger,
		authController: authController,
		authHandler:    authHandler,
	}
}
