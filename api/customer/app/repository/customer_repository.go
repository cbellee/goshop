package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/cbellee/goShop-customerService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	conf, err = config.LoadConfig()
	//dbName       = "goShopDb"
	//dbCollection = "customers"
)

// CustomerRepository used to get customer data from mongodDB
type CustomerRepository interface {
	Get(id int64) (*Customer, error)
	List() ([]*Customer, error)
	Insert(customer Customer) (lastInsertID interface{}, err error)
	Delete(id int64) error
	Update(customer Customer, id int64) (upsertedCount int64, err error)
}

type customerRepository struct {
	client *mongo.Client
}

// NewCustomerRepository returns a new instance of CustomerRepository
func NewCustomerRepository(client *mongo.Client) CustomerRepository {
	return &customerRepository{
		client: client,
	}
}

// Get
func (r *customerRepository) Get(id int64) (customer *Customer, err error) {
	var result *Customer
	filter := bson.D{{"id", id}}
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	return result, err
}

// List
func (r *customerRepository) List() (customers []*Customer, err error) {
	var results []*Customer
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	//defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var element Customer
		err := cur.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &element)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	fmt.Printf("Found %d documents\n", len(results))
	return results, nil
}

// Delete
func (r *customerRepository) Delete(id int64) (err error) {
	filter := bson.D{{"id", id}}
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v document in '%s' collection\n", result.DeletedCount, conf.DbCollection)
	return err
}

// Insert
func (r *customerRepository) Insert(customer Customer) (lastInsertID interface{}, err error) {
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	insertResult, err := collection.InsertOne(context.Background(), customer)
	if err != nil {
		log.Fatal(err)
	}

	return insertResult.InsertedID, err
}

// Update
func (r *customerRepository) Update(customer Customer, id int64) (upsertedCount int64, err error) {
	filter := bson.D{{"id", id}}
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	updateResult, err := collection.UpdateOne(context.Background(), filter, customer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v documents in '%s' collection\n", updateResult.UpsertedCount, conf.DbCollection)
	return updateResult.UpsertedCount, err
}
