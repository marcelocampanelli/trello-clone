package main

import (
	"context"
	"github.com/marcelocampanelli/trello-clone/internal/infra/database/mongodb"
	webServer "github.com/marcelocampanelli/trello-clone/internal/infra/web/server"
	"net/http"
)

func main() {
	client, err := mongodb.ConnectMongoDB()
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())

	server := webServer.NewServer(client).Start()
	err = http.ListenAndServe(":8080", server)
	if err != nil {
		panic(err)
	}
}
