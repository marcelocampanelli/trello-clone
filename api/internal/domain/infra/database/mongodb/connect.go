package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectMongoDB() (*mongo.Client, error) {
	//mongodb://<username>:<password>@<host>:<port>
	connectionString := "mongodb://root:root@localhost:27017"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("error on ping mongodb")
		return nil, err
	}

	fmt.Println("mongodb connected")

	return client, nil
}
