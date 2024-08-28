package airdrop

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		return nil, fmt.Errorf("MongoConnect: %v", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("MongoConnect Ping: %v", err)
	}
	return client.Database(dbname), nil
}

func InsertOneDoc(db string, collection string, doc interface{}) (interface{}, error) {
	database, err := MongoConnect(db)
	if err != nil {
		return nil, err
	}
	collectionRef := database.Collection(collection)
	insertResult, err := collectionRef.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, fmt.Errorf("InsertOneDoc: %v", err)
	}
	return insertResult.InsertedID, nil
}

func InsertAirdropFree(name string, task string, link string) (interface{}, error) {
	freeAirdrop := AirdropFree{
		ID:   primitive.NewObjectID(), 
		Name: name,
		Task: task,
		Link: link,
	}
	return InsertOneDoc("airdrop", "airdrop_free", freeAirdrop)
}

func InsertAirdropPaid(name string, task string, link string) (interface{}, error) {
	paidAirdrop := AirdropPaid{
		ID:   primitive.NewObjectID(), 
		Name: name,
		Task: task,
		Link: link,
	}
	return InsertOneDoc("airdrop", "airdrop_paid", paidAirdrop)
}

func GetAllAirdropFree() ([]AirdropFree, error) {
	database, err := MongoConnect("airdrop")
	if err != nil {
		return nil, err
	}
	collection := database.Collection("airdrop_free")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetAllAirdropFree Find: %v", err)
	}
	var airdrops []AirdropFree
	if err = cursor.All(context.TODO(), &airdrops); err != nil {
		return nil, fmt.Errorf("GetAllAirdropFree All: %v", err)
	}
	return airdrops, nil
}

func GetAllAirdropPaid() ([]AirdropPaid, error) {
	database, err := MongoConnect("airdrop")
	if err != nil {
		return nil, err
	}
	collection := database.Collection("airdrop_paid")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetAllAirdropPaid Find: %v", err)
	}
	var airdrops []AirdropPaid
	if err = cursor.All(context.TODO(), &airdrops); err != nil {
		return nil, fmt.Errorf("GetAllAirdropPaid All: %v", err)
	}
	return airdrops, nil
}
