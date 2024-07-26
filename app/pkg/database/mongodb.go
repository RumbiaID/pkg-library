package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClientRepository struct {
	Client *mongo.Client
	DB     *mongo.Database
}

type MongoConfig struct {
	MongoURL string
	MongoDB  string
}

func NewMongoDBRepository(config *MongoConfig) (
	*MongoDBClientRepository, error,
) {
	//var dsn string
	//dsn = fmt.Sprintf(os.Getenv("MONGODB_URL"))
	clientOptions := options.Client().ApplyURI(config.MongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	template := client.Database(config.MongoDB)
	if err != nil {
		logrus.Error(fmt.Sprintf("Cannot connect to MongoDB. %v", err))
		return nil, errors.New("Cannot connect to MongoDB: " + err.Error())
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logrus.Error(fmt.Sprintf("Cannot send/receive data with MongoDB. %v", err))
		return nil, errors.New("Cannot send/receive data with MongoDB: " + err.Error())
	}
	return &MongoDBClientRepository{Client: client, DB: template}, nil
}
