package repositories

import (
	"context"
	"github.com/corentings/stouuf-bot/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type KarmaRepository struct {
	DBConn *mongo.Database
}

func (KarmaRepository *KarmaRepository) CreateKarma(karma models.Karma) error {
	collection := KarmaRepository.DBConn.Collection("karma")

	_, err := collection.InsertOne(context.TODO(), karma)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (KarmaRepository *KarmaRepository) UpdateKarma(karma models.Karma) error {
	collection := KarmaRepository.DBConn.Collection("karma")

	_, err := collection.UpdateOne(context.TODO(), bson.D{{"userid", karma.UserID}, {"guildid", karma.GuildID}}, bson.D{{"$set", bson.D{{"value", karma.Value}}}})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (KarmaRepository *KarmaRepository) GetKarma(userID, guildID string) (*models.Karma, error) {
	collection := KarmaRepository.DBConn.Collection("karma")
	result := new(models.Karma)
	err := collection.FindOne(context.TODO(), bson.D{{"userid", userID}, {"guildid", guildID}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
