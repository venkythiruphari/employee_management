package routers

import (
	"employee-management/config"
	"employee-management/internal/employee"
	"employee-management/internal/middleware"
	"employee-management/internal/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize Repositories
	userRepo := users.NewRepository(db)
	employeeRepo := employee.NewRepository(db)

	// Initialize Services
	userService := users.NewService(userRepo)
	employeeService := employee.NewService(employeeRepo)

	// Initialize Handlers
	userHandler := users.NewHandler(userService, cfg)
	employeeHandler := employee.NewHandler(employeeService)

	// Register Middleware
	authMiddleware := middleware.AuthMiddleware(cfg)

	// API Grouping
	api := router.Group("/api/v1") // Example API versioning

	// Register User Routes
	users.RegisterUserRoutes(api, userHandler)

	// Register Employee Routes (protected by JWT middleware)
	employee.RegisterEmployeeRoutes(api, employeeHandler, authMiddleware)

	return router
}
