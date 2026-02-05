package employee

import (
	"employee-management/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

func (h *Handler) GetAllEmployees(c *gin.Context) {
	employees, err := h.Service.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve employees"})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (h *Handler) GetEmployeeByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	employee, err := h.Service.GetEmployeeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employee.ID = uint(id)

	if err := h.Service.UpdateEmployee(&employee); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // Service returns "employee not found"
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	if err := h.Service.DeleteEmployee(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // Service returns "employee not found"
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) CalculateNetSalary(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	netSalary, err := h.Service.CalculateNetSalary(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // Service returns "employee not found"
		return
	}

	c.JSON(http.StatusOK, netSalary)
}

func (h *Handler) GetSalaryMetricsByCountry(c *gin.Context) {
	country := c.Query("country")
	if country == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Country query parameter is required"})
		return
	}

	metrics, err := h.Service.GetSalaryMetricsByCountry(country)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // Service returns "no data found"
		return
	}

	c.JSON(http.StatusOK, metrics)
}

func (h *Handler) GetAverageSalaryByJobTitle(c *gin.Context) {
	jobTitle := c.Query("job_title")
	if jobTitle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Job title query parameter is required"})
		return
	}

	avgSalary, err := h.Service.GetAverageSalaryByJobTitle(jobTitle)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // Service returns "no data found"
		return
	}

	c.JSON(http.StatusOK, gin.H{"job_title": jobTitle, "average_salary": avgSalary})
}
