package test

import (
	mocks "github.com/corentings/stouuf-bot/mocks/interfaces"
	"github.com/corentings/stouuf-bot/models"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestCoursRepository_GetCours(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		want    *models.CoursModel
		arg     string
	}{
		{
			name:    "TestCoursRepository_GetCours_normal",
			wantErr: false,
			want: &models.CoursModel{
				UserID: "123",
				Value:  10,
			},
			arg: "123",
		},
		{
			name:    "TestCoursRepository_GetCours_not_found",
			wantErr: true,
			want:    nil,
			arg:     "10",
		},
	}

	repo := &mocks.ICoursRepository{}
	repo.On("GetCours", "10").Return(nil, mongo.ErrNoDocuments)
	repo.On("GetCours", "123").Return(&models.CoursModel{
		UserID: "123",
		Value:  10,
	}, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetCours(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CoursRepository.GetCours() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CoursRepository.GetCours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoursRepository_CreateCours(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		arg     models.CoursModel
	}{
		{
			name:    "TestCoursRepository_CreateCours_normal",
			wantErr: false,
			arg: models.CoursModel{
				UserID: "123",
				Value:  10,
			},
		},
	}

	repo := &mocks.ICoursRepository{}
	repo.On("CreateCours", models.CoursModel{
		UserID: "123",
		Value:  10,
	}).Return(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.CreateCours(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("CoursRepository.CreateCours() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCoursRepository_UpdateCours(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		arg     models.CoursModel
	}{
		{
			name:    "TestCoursRepository_UpdateCours_normal",
			wantErr: false,
			arg: models.CoursModel{
				UserID: "123",
				Value:  10,
			},
		},
	}

	repo := &mocks.ICoursRepository{}
	repo.On("UpdateCours", models.CoursModel{
		UserID: "123",
		Value:  10,
	}).Return(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.UpdateCours(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("CoursRepository.UpdateCours() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
