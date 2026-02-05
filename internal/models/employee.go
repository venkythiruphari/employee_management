package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	FullName    string  `json:"full_name"`
	JobTitle    string  `json:"job_title"`
	Country     string  `json:"country"`
	GrossSalary float64 `json:"gross_salary"`
}
