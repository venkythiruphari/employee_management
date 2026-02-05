package users

import (
	"employee-management/internal/models"
	"errors"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var ErrDuplicateUsername = errors.New("duplicate username")

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(user *models.User) error {
	err := r.DB.Create(user).Error
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			return ErrDuplicateUsername
		}
		return err
	}
	return nil
}

func (r *Repository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
