package test

import (
	mocks "github.com/corentings/stouuf-bot/mocks/interfaces"
	"github.com/corentings/stouuf-bot/models"
	"github.com/corentings/stouuf-bot/services"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestCoursService_AddCours(t *testing.T) {
	t.Run("TestCoursService_AddCours_normal", func(t *testing.T) {
		repository := &mocks.ICoursRepository{}
		repository.On("GetCours", "123").Return(&models.CoursModel{
			UserID: "123",
			Value:  5,
		}, nil)
		repository.On("UpdateCours", models.CoursModel{
			UserID: "123",
			Value:  15,
		}).Return(nil)
		service := &services.CoursService{ICoursRepository: repository}
		got, err := service.AddCours("123", 10)
		if err != nil {
			t.Errorf("CoursService.AddCours() error = %v", err)
			return
		}
		if got.Value != 15 {
			t.Errorf("CoursService.AddCours() = %v, want %v", got, 15)
		}

	})
	t.Run("TestCoursService_AddCours_errorNoDocument", func(t *testing.T) {
		repository := &mocks.ICoursRepository{}
		repository.On("GetCours", "123").Return(nil, mongo.ErrNoDocuments)
		repository.On("CreateCours", models.CoursModel{
			UserID: "123",
			Value:  10,
		}).Return(nil)
		service := &services.CoursService{ICoursRepository: repository}
		got, err := service.AddCours("123", 10)
		if err != nil {
			t.Errorf("CoursService.AddCours() error = %v", err)
			return
		}
		if got.Value != 10 {
			t.Errorf("CoursService.AddCours() = %v, want %v", got, 10)
		}
	})
}

func TestCoursService_GetCours(t *testing.T) {
	t.Run("TestCoursService_GetCours_normal", func(t *testing.T) {
		repository := &mocks.ICoursRepository{}
		repository.On("GetCours", "123").Return(&models.CoursModel{
			UserID: "123",
			Value:  10,
		}, nil)
		service := &services.CoursService{ICoursRepository: repository}
		got, err := service.GetCours("123")
		if err != nil {
			t.Errorf("CoursService.GetCours() error = %v", err)
			return
		}
		if got.Value != 10 {
			t.Errorf("CoursService.GetCours() = %v, want %v", got, 10)
		}
	})
	t.Run("TestCoursService_GetCours_errorNoDocument", func(t *testing.T) {
		repository := &mocks.ICoursRepository{}
		repository.On("GetCours", "123").Return(nil, mongo.ErrNoDocuments)
		repository.On("CreateCours", models.CoursModel{
			UserID: "123",
			Value:  0,
		}).Return(nil)
		service := &services.CoursService{ICoursRepository: repository}
		got, err := service.GetCours("123")
		if err != nil {
			t.Errorf("CoursService.GetCours() error = %v", err)
			return
		}
		if got.Value != 0 {
			t.Errorf("CoursService.GetCours() = %v, want %v", got, 0)
		}
	})
}

func TestCoursService_RemoveCours(t *testing.T) {
	t.Run("TestCoursService_RemoveCours_normal", func(t *testing.T) {
		repository := &mocks.ICoursRepository{}
		repository.On("GetCours", "123").Return(&models.CoursModel{
			UserID: "123",
			Value:  10,
		}, nil)
		repository.On("UpdateCours", models.CoursModel{
			UserID: "123",
			Value:  5,
		}).Return(nil)
		service := &services.CoursService{ICoursRepository: repository}
		got, err := service.RemoveCours("123", 5)
		if err != nil {
			t.Errorf("CoursService.RemoveCours() error = %v", err)
			return
		}
		if got.Value != 5 {
			t.Errorf("CoursService.RemoveCours() = %v, want %v", got, 5)
		}
	})
	t.Run("TestCoursService_RemoveCours_errorNoDocument", func(t *testing.T) {
		repository := &mocks.ICoursRepository{}
		repository.On("GetCours", "123").Return(nil, mongo.ErrNoDocuments)
		repository.On("CreateCours", models.CoursModel{
			UserID: "123",
			Value:  0,
		}).Return(nil)
		service := &services.CoursService{ICoursRepository: repository}
		got, err := service.RemoveCours("123", 5)
		if err != nil {
			t.Errorf("CoursService.RemoveCours() error = %v", err)
			return
		}
		if got.Value != 0 {
			t.Errorf("CoursService.RemoveCours() = %v, want %v", got, 0)
		}
	})
	t.Run("TestCoursService_RemoveCours_moreThanValue", func(t *testing.T) {
		repository := &mocks.ICoursRepository{}
		repository.On("GetCours", "123").Return(&models.CoursModel{
			UserID: "123",
			Value:  10,
		}, nil)
		repository.On("UpdateCours", models.CoursModel{
			UserID: "123",
			Value:  0,
		}).Return(nil)
		service := &services.CoursService{ICoursRepository: repository}
		got, err := service.RemoveCours("123", 15)
		if err != nil {
			t.Errorf("CoursService.RemoveCours() error = %v", err)
			return
		}
		if got.Value != 0 {
			t.Errorf("CoursService.RemoveCours() = %v, want %v", got, 0)
		}
	})
}
