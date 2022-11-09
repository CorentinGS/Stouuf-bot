package interfaces

import "github.com/corentings/stouuf-bot/models"

type ICoursRepository interface {
	CreateCours(cours models.CoursModel) error
	UpdateCours(cours models.CoursModel) error
	GetCours(userID string) (*models.CoursModel, error)
}

type ICoursService interface {
	AddCours(userID string, value uint) (*models.CoursModel, error)
	GetCours(userID string) (*models.CoursModel, error)
	RemoveCours(userID string, value uint) (*models.CoursModel, error)
}
