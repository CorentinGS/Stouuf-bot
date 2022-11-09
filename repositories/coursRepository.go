package repositories

import (
	"context"
	"github.com/corentings/stouuf-bot/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CoursRepository struct {
	DBConn *mongo.Database
}

func (c *CoursRepository) GetCours(userID string) (*models.CoursModel, error) {
	collection := c.DBConn.Collection("cours")
	result := new(models.CoursModel)
	err := collection.FindOne(context.TODO(), bson.D{{"userid", userID}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CoursRepository) CreateCours(cours models.CoursModel) error {
	collection := c.DBConn.Collection("cours")

	_, err := collection.InsertOne(context.TODO(), cours)
	if err != nil {
		return err
	}
	return nil
}

func (c *CoursRepository) UpdateCours(cours models.CoursModel) error {
	collection := c.DBConn.Collection("cours")

	_, err := collection.UpdateOne(context.TODO(),
		bson.D{{"userid", cours.UserID}},
		bson.D{
			{"$set", bson.D{
				{"value", cours.Value},
			}}},
	)
	if err != nil {
		return err
	}

	return nil
}
