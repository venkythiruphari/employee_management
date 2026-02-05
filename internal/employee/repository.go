package employee

import (
	"employee-management/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateEmployee(employee *models.Employee) error {
	return r.DB.Create(employee).Error
}

func (r *Repository) GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *Repository) GetEmployeeByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.First(&employee, id).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *Repository) UpdateEmployee(employee *models.Employee) error {
	return r.DB.Save(employee).Error
}

func (r *Repository) DeleteEmployee(id uint) error {
	return r.DB.Delete(&models.Employee{}, id).Error
}

func (r *Repository) GetSalaryMetricsByCountry(country string) (min, max, avg float64, err error) {
	type Result struct {
		Min float64
		Max float64
		Avg float64
	}
	var result Result
	err = r.DB.Model(&models.Employee{}).
		Select("MIN(gross_salary) as min, MAX(gross_salary) as max, AVG(gross_salary) as avg").
		Where("country = ?", country).
		Scan(&result).Error

	if err != nil {
		return 0, 0, 0, err
	}
	return result.Min, result.Max, result.Avg, nil
}

func (r *Repository) GetAverageSalaryByJobTitle(jobTitle string) (avg float64, err error) {
	type Result struct {
		Avg float64
	}
	var result Result
	err = r.DB.Model(&models.Employee{}).
		Select("AVG(gross_salary) as avg").
		Where("job_title = ?", jobTitle).
		Scan(&result).Error

	if err != nil {
		return 0, err
	}
	return result.Avg, nil
}
