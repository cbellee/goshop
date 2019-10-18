package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cbellee/goShop-orderService/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

// Connect to database by creating a client
func Connect(conf config.Config) (*mongo.Client, error) {
	connStr := connectionString(conf)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client(), options.Client().ApplyURI(connStr))
	err = client.Connect(ctx)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// check connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Print("Connection to database with connection string: ", connStr, "\n")		
		log.Fatal("Unable to connect to the database with error ", err)
	} else {
		fmt.Print("connected to MongoDB!\n")
		fmt.Print("connStr: ", connStr, "\n")
	}

	return client, nil
}
