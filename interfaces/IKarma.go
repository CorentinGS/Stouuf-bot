package interfaces

import "github.com/corentings/stouuf-bot/models"

type IKarmaRepository interface {
	CreateKarma(karma models.Karma) error
	UpdateKarma(karma models.Karma) error
	GetKarma(userID, guildID string) (*models.Karma, error)
}

type IKarmaService interface {
	CreateKarma(karma models.Karma) error
	UpdateKarma(karma models.Karma) error
	GetKarma(userID, guildID string) (*models.Karma, error)
	AddKarma(userID, guildID string, amount uint) (*models.Karma, error)
}
