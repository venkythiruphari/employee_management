package employee

import (
	"github.com/gin-gonic/gin"
)

func RegisterEmployeeRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	// Protected routes for employee management
	employeeRoutes := router.Group("/employees")
	employeeRoutes.Use(authMiddleware)
	{
		employeeRoutes.POST("/", handler.CreateEmployee)
		employeeRoutes.GET("/", handler.GetAllEmployees)
		employeeRoutes.GET("/:id", handler.GetEmployeeByID)
		employeeRoutes.PUT("/:id", handler.UpdateEmployee)
		employeeRoutes.DELETE("/:id", handler.DeleteEmployee)
		employeeRoutes.GET("/:id/salary/net", handler.CalculateNetSalary)
	}

	// Metrics routes (also protected)
	metricsRoutes := router.Group("/metrics")
	metricsRoutes.Use(authMiddleware)
	{
		metricsRoutes.GET("/salary/country", handler.GetSalaryMetricsByCountry)
		metricsRoutes.GET("/salary/job-title", handler.GetAverageSalaryByJobTitle)
	}
}