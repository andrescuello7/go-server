package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr = "conteo"
	pwd = "admin"
	db  = "goprueba"
)

func GetDatabase(collection string) *mongo.Collection {

	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.qa9gg.mongodb.net/%s?retryWrites=true&w=majority", usr, pwd, db)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	log.Println(err, "error")
	log.Println(ctx, "ctx")

	return client.Database(db).Collection(collection)
}
