package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/cbellee/goShop-orderService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	conf, err = config.LoadConfig()
	//conf.DbName       = "goShopDb"
	//conf.DbCollection = "orders"
)

// OrderRepository used to get order data from mongodDB
type OrderRepository interface {
	Get(id int64) (*Order, error)
	List() ([]*Order, error)
	Insert(order Order) (lastInsertID interface{}, err error)
	Delete(id int64) error
	Update(order Order, id int64) (upsertedCount int64, err error)
}

type orderRepository struct {
	client *mongo.Client
}

// NewOrderRepository returns a new instance of OrderRepository
func NewOrderRepository(client *mongo.Client) OrderRepository {
	return &orderRepository{
		client: client,
	}
}

// Get
func (r *orderRepository) Get(id int64) (order *Order, err error) {
	var result *Order
	filter := bson.D{{"id", id}}
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	return result, err
}

// List
func (r *orderRepository) List() (orders []*Order, err error) {
	var results []*Order
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var element Order
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
func (r *orderRepository) Delete(id int64) (err error) {
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
func (r *orderRepository) Insert(order Order) (lastInsertID interface{}, err error) {
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	insertResult, err := collection.InsertOne(context.Background(), order)
	if err != nil {
		log.Fatal(err)
	}

	return insertResult.InsertedID, err
}

// Update
func (r *orderRepository) Update(order Order, id int64) (upsertedCount int64, err error) {
	filter := bson.D{{"id", id}}
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	updateResult, err := collection.UpdateOne(context.Background(), filter, order)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v documents in '%s' collection\n", updateResult.UpsertedCount, conf.DbCollection)
	return updateResult.UpsertedCount, err
}
