package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/priyanshu360/urlShortnerApp.git/config"
	"github.com/priyanshu360/urlShortnerApp.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

func InitMongo() (*MongoDB, error) {
	log.Println("init Mongo ... ")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	credential := options.Credential{
		Username:   config.DB_USER,
		Password:   config.DB_PASSWORD,
		AuthSource: config.DB_NAME,
	}
	clientOpts := options.Client().ApplyURI(config.DATABASE_URL).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Println("Error connecting to MongoDB:", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Println("Error pinging MongoDB server:", err)
		return nil, err
	}

	return &MongoDB{Client: client}, nil
}

func (m MongoDB) GetLongUrl(hash string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database(config.DB_NAME).Collection("urlrecords")
	filter := bson.M{"hash": hash}

	var result models.URLRecord
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return "",  fmt.Errorf("URL record not found for hash: %s", hash)
		}
		return "", err
	}

	return result.LongURL, nil
}

func (m MongoDB) GetHashValue(longURL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database(config.DB_NAME).Collection("urlrecords")

	filter := bson.M{"long_url": longURL}

	var result models.URLRecord

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("Hash not found for long URL: %s", longURL)
		}
		return "", err
	}

	return result.Hash, nil
}

func (m MongoDB) DeleteURLRecord(hash string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database(config.DB_NAME).Collection("urlrecords")
	filter := bson.M{"hash": hash}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("URL record not found for hash: %s", hash)
	}

	return nil
}

func (m MongoDB) CreateURLRecord(record *models.URLRecord) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database(config.DB_NAME).Collection("urlrecords")
	if _, err := collection.InsertOne(ctx, record); err != nil {
		return err
	}

	return nil
}
