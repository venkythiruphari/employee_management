package employee

import (
	"employee-management/internal/models"
	"errors"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateEmployee(employee *models.Employee) error {
	return s.Repo.CreateEmployee(employee)
}

func (s *Service) GetAllEmployees() ([]models.Employee, error) {
	return s.Repo.GetAllEmployees()
}

func (s *Service) GetEmployeeByID(id uint) (*models.Employee, error) {
	return s.Repo.GetEmployeeByID(id)
}

func (s *Service) UpdateEmployee(employee *models.Employee) error {
	_, err := s.Repo.GetEmployeeByID(employee.ID)
	if err != nil {
		return errors.New("employee not found")
	}
	return s.Repo.UpdateEmployee(employee)
}

func (s *Service) DeleteEmployee(id uint) error {
	_, err := s.Repo.GetEmployeeByID(id)
	if err != nil {
		return errors.New("employee not found")
	}
	return s.Repo.DeleteEmployee(id)
}

type NetSalary struct {
	GrossSalary    float64 `json:"gross_salary"`
	DeductionAmount float64 `json:"deduction_amount"`
	NetSalary      float64 `json:"net_salary"`
}

func (s *Service) CalculateNetSalary(employeeID uint) (*NetSalary, error) {
	employee, err := s.Repo.GetEmployeeByID(employeeID)
	if err != nil {
		return nil, errors.New("employee not found")
	}

	deductionRate := 0.0
	switch employee.Country {
	case "India":
		deductionRate = 0.10
	case "United States":
		deductionRate = 0.12
	}

	deductionAmount := employee.GrossSalary * deductionRate
	netSalary := employee.GrossSalary - deductionAmount

	return &NetSalary{
		GrossSalary:    employee.GrossSalary,
		DeductionAmount: deductionAmount,
		NetSalary:      netSalary,
	}, nil
}

type SalaryMetrics struct {
	MinSalary float64 `json:"min_salary"`
	MaxSalary float64 `json:"max_salary"`
	AvgSalary float64 `json:"avg_salary"`
}

func (s *Service) GetSalaryMetricsByCountry(country string) (*SalaryMetrics, error) {
	min, max, avg, err := s.Repo.GetSalaryMetricsByCountry(country)
	if err != nil {
		return nil, err
	}
	if min == 0 && max == 0 && avg == 0 {
		return nil, errors.New("no data found for this country")
	}
	return &SalaryMetrics{MinSalary: min, MaxSalary: max, AvgSalary: avg}, nil
}

func (s *Service) GetAverageSalaryByJobTitle(jobTitle string) (float64, error) {
	avg, err := s.Repo.GetAverageSalaryByJobTitle(jobTitle)
	if err != nil {
		return 0, err
	}
	if avg == 0 {
		return 0, errors.New("no data found for this job title")
	}
	return avg, nil
}
