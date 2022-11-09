package services

import (
	"errors"
	"github.com/corentings/stouuf-bot/commands/karma"
	"github.com/corentings/stouuf-bot/database"
	"github.com/corentings/stouuf-bot/interfaces"
	"github.com/corentings/stouuf-bot/models"
	"github.com/corentings/stouuf-bot/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type KarmaService struct {
	interfaces.IKarmaRepository
}

func (k *kernel) InjectKarmaCommandHandler() karma.KarmaCommand {
	DBConn := database.Mg.DB
	karmaRepo := &repositories.KarmaRepository{DBConn: DBConn}
	karmaService := &KarmaService{IKarmaRepository: karmaRepo}
	karmaCommand := karma.KarmaCommand{IKarmaService: karmaService}
	return karmaCommand
}

func (KarmaService *KarmaService) CreateKarma(karma models.Karma) error {
	err := KarmaService.IKarmaRepository.CreateKarma(karma)
	if err != nil {
		return err
	}
	return nil
}

func (KarmaService *KarmaService) UpdateKarma(karma models.Karma) error {
	err := KarmaService.IKarmaRepository.UpdateKarma(karma)
	if err != nil {
		return err
	}
	return nil
}

func (KarmaService *KarmaService) GetKarma(userID, guildID string) (*models.Karma, error) {
	getKarma, err := KarmaService.IKarmaRepository.GetKarma(userID, guildID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			result := new(models.Karma)
			result.SetKarma(userID, guildID, 0)
			err = KarmaService.CreateKarma(*result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err

	}
	return getKarma, nil
}

func (KarmaService *KarmaService) AddKarma(userID, guildID string, amount uint) (*models.Karma, error) {
	getKarma, err := KarmaService.GetKarma(userID, guildID)
	if err != nil {
		return nil, err
	}
	getKarma.AddKarma(amount)
	err = KarmaService.UpdateKarma(*getKarma)
	if err != nil {
		return nil, err
	}
	return getKarma, nil
}
