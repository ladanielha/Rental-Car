package repositories

import (
	"Rental-car/internal/models"

	"gorm.io/gorm"
)

type RepoUser interface {
	Register(*models.Users) error
	GetUser(email string) (*models.Users, error)
}

type repoUser struct {
	dB *gorm.DB
}

func NewRepoUser(db *gorm.DB) RepoUser {
	return &repoUser{dB: db}
}

func (r *repoUser) Register(user *models.Users) error {
	err := r.dB.Create(user).Error
	return err
}

func (r *repoUser) GetUser(email string) (*models.Users, error) {
	user := &models.Users{}

	err := r.dB.Find(user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
