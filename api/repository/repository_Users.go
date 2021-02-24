package repository

import "blog-api/api/models"

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll(page int,size int) ([]models.User, error)
	FindByID(string) (models.User, error)
	UpdateByID(string, models.User) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string) ([]models.User, error)
}

