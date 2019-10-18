package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cbellee/goShop-customerService/app/repository"
	"github.com/cbellee/goShop-customerService/config"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create database
func Create(conf config.Config) (res *mongo.InsertManyResult, err error) {
	client, err := Connect(conf)
	if err != nil {
		log.Fatal(err)
	}

	// create or get existing db & collection
	collection := client.Database(conf.DbName).Collection(conf.DbCollection)
	numDocs, err := collection.EstimatedDocumentCount(context.Background(), nil, nil)
	if numDocs > 0 {
		fmt.Printf("found %d documents in collection '%s'\n", numDocs, collection.Name())
		return
	}

	// Load values from JSON file to model
	var customers []repository.Customer
	byteValues, err := ioutil.ReadFile("seed_data.json")
	if err != nil {
		log.Fatal(err)
	}

	// read customers from JSON file
	json.Unmarshal(byteValues, &customers)
	var cst []interface{}
	for _, c := range customers {
		cst = append(cst, c)
	}

	// insert customers into DB
	res, err = collection.InsertMany(context.Background(), cst)
	if err != nil {
		return nil, err
	}
	numInserts := len(res.InsertedIDs)
	fmt.Printf("inserted %d customers into collection '%s' in database '%s'\n", numInserts, conf.DbCollection, conf.DbName)
	return res, nil
}

// Drop database
func Drop(conf config.Config) (err error) {
	client, err := Connect(conf)
	if err != nil {
		return err
	}

	client.Database(conf.DbName).Drop(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("dropped database %s\n", conf.DbName)
	return nil
}

// Migrate database
/* func Migrate(conf config.Config, args cli.Args) error {
	source := migrationSource(args)
	log.Printf("Migrate database from source '%s'\n", source)

	migration, err := migrate.New(source, connectionString(conf))
	if err != nil {
		return err
	}
	defer migration.Close()
	return migration.Up()
} */

// Rollback database
/* func Rollback(conf config.Config, args cli.Args) error {
	source := migrationSource(args)
	log.Printf("Migrate database from source '%s'\n", source)

	migration, err := migrate.New(source, connectionString(conf))
	if err != nil {
		return err
	}
	defer migration.Close()
	return migration.Down()
} */

// ResetTestDB reset test database
/* func ResetTestDB(conf config.Config, source string) (err error) {
	conn, err := sql.Open("postgres", connectionStringWithDBName(conf, "template1"))
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = conn.Exec(fmt.Sprintf(`DROP DATABASE IF EXISTS "%s"`, conf.DbName))
	if err != nil {
		return
	}
	_, err = conn.Exec(fmt.Sprintf(`CREATE DATABASE "%s"`, conf.DbName))
	if err != nil {
		return
	}

	migration, err := migrate.New(source, connectionString(conf))
	if err != nil {
		return err
	}
	defer migration.Close()
	return migration.Up()
} */

/* func migrationSource(args cli.Args) string {
	dir := config.DefaultMigrationDirectory
	if len(args) > 0 {
		dir = args.First()
	}
	return fmt.Sprintf("file://%s", dir)
} */
