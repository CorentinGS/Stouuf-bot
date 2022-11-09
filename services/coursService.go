package services

import (
	"errors"
	"github.com/corentings/stouuf-bot/commands/cours"
	"github.com/corentings/stouuf-bot/database"
	"github.com/corentings/stouuf-bot/interfaces"
	"github.com/corentings/stouuf-bot/models"
	"github.com/corentings/stouuf-bot/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type CoursService struct {
	interfaces.ICoursRepository
}

func (k *kernel) InjectCoursCommandHandler() cours.CoursCommand {
	DBConn := database.Mg.DB

	coursRepo := &repositories.CoursRepository{DBConn: DBConn}
	coursService := &CoursService{ICoursRepository: coursRepo}
	coursCommand := cours.CoursCommand{ICoursService: coursService}
	return coursCommand
}

func (c *CoursService) AddCours(userID string, value uint) (*models.CoursModel, error) {
	getCours, err := c.ICoursRepository.GetCours(userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			result := &models.CoursModel{
				UserID: userID,
				Value:  value,
			}
			err = c.ICoursRepository.CreateCours(*result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	getCours.Value += value
	err = c.UpdateCours(*getCours)
	if err != nil {
		return nil, err
	}

	return getCours, nil
}

func (c *CoursService) GetCours(userID string) (*models.CoursModel, error) {
	getCours, err := c.ICoursRepository.GetCours(userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			result := &models.CoursModel{
				UserID: userID,
				Value:  0,
			}
			err = c.ICoursRepository.CreateCours(*result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}
	return getCours, nil
}

func (c *CoursService) RemoveCours(userID string, value uint) (*models.CoursModel, error) {
	getCours, err := c.ICoursRepository.GetCours(userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			result := &models.CoursModel{
				UserID: userID,
				Value:  0,
			}
			err = c.ICoursRepository.CreateCours(*result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	if getCours.Value > value {
		getCours.Value -= value
	} else {
		getCours.Value = 0
	}

	err = c.ICoursRepository.UpdateCours(*getCours)
	if err != nil {
		return nil, err
	}

	return getCours, nil
}
