package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

func ConnectMongoDB() (*mongo.Client, error) {
	//mongodb://<username>:<password>@<host>:<port>
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	connectionString := "mongodb://mongodb:27017"

	ctx, cancel := context.WithTimeout(context.Background(), 10000000*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		sugar.Error("Error on ping in MongoDB")
		return nil, err
	}

	sugar.Info("Connected in MongoDB")

	return client, nil
}
