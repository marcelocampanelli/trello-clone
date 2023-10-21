package main

import (
	"context"
	"github.com/marcelocampanelli/trello-clone/internal/domain/infra/database/mongodb"
)

func main() {
	client, err := mongodb.ConnectMongoDB()
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())
}
