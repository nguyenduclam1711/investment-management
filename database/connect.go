package database

import (
	"context"
	"fmt"

	"github.com/nguyenduclam1711/investment-management/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MgClient *mongo.Client
	DB       *mongo.Database
)

func ConnectDB() {
	uri := config.Config("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database(config.Config("MONGODB_DATABASE"))

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connect to database successfully")
	MgClient = client
	DB = db
}
