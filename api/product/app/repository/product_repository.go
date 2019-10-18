package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/cbellee/goShop-productService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	conf, err = config.LoadConfig()
	//dbName       = "goShopDb"
	//dbCollection = "products"
)

// ProductRepository used to get product data from mongodDB
type ProductRepository interface {
	Get(id int64) (*Product, error)
	List() ([]*Product, error)
	Insert(product Product) (lastInsertID interface{}, err error)
	Delete(id int64) error
	Update(product Product, id int64) (upsertedCount int64, err error)
}

type productRepository struct {
	client *mongo.Client
}

// NewProductRepository returns a new instance of ProductRepository
func NewProductRepository(client *mongo.Client) ProductRepository {
	return &productRepository{
		client: client,
	}
}

// Get
func (r *productRepository) Get(id int64) (product *Product, err error) {
	var result *Product
	filter := bson.D{{"id", id}}
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	return result, err
}

// List
func (r *productRepository) List() (products []*Product, err error) {
	var results []*Product
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var element Product
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
func (r *productRepository) Delete(id int64) (err error) {
	filter := bson.D{{"id", id}}
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v document in 'productColl' collection\n", result.DeletedCount)
	return err
}

// Insert
func (r *productRepository) Insert(product Product) (lastInsertID interface{}, err error) {
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	insertResult, err := collection.InsertOne(context.Background(), product)
	if err != nil {
		log.Fatal(err)
	}

	return insertResult.InsertedID, err
}

// Update
func (r *productRepository) Update(product Product, id int64) (upsertedCount int64, err error) {
	filter := bson.D{{"id", id}}
	collection := r.client.Database(conf.DbName).Collection(conf.DbCollection)
	updateResult, err := collection.UpdateOne(context.Background(), filter, product)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v documents in 'productColl' collection\n", updateResult.UpsertedCount)
	return updateResult.UpsertedCount, err
}
