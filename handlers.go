package airdrop

import (
	"context"
	"errors"
	"fmt"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoString string = os.Getenv("MONGOSTRING")
	database    *mongo.Database
)

func init() {
	db, err := MongoConnect("airdrop")
	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v\n", err)
		os.Exit(1)
	}
	database = db
}

// MongoConnect initializes a connection to the MongoDB server.
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

func InsertOneDoc(collection string, doc interface{}) (interface{}, error) {
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
	return InsertOneDoc("airdrop_free", freeAirdrop)
}

func InsertAirdropPaid(name string, task string, link string) (interface{}, error) {
	paidAirdrop := AirdropPaid{
		ID:   primitive.NewObjectID(),
		Name: name,
		Task: task,
		Link: link,
	}
	return InsertOneDoc("airdrop_paid", paidAirdrop)
}

func GetAllAirdropFree() ([]AirdropFree, error) {
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

func GetAirdropFreeByID(id primitive.ObjectID) (AirdropFree, error) {
	collection := database.Collection("airdrop_free")
	var airdrop AirdropFree
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&airdrop)
	if err != nil {
		return AirdropFree{}, err
	}
	return airdrop, nil
}

func GetAirdropPaidByID(id primitive.ObjectID) (AirdropPaid, error) {
	collection := database.Collection("airdrop_paid")
	var airdrop AirdropPaid
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&airdrop)
	if err != nil {
		return AirdropPaid{}, err
	}
	return airdrop, nil
}

func GetAllAirdrop() ([]interface{}, error) {
	var allAirdrops []interface{}

	freeAirdrops, err := GetAllAirdropFree()
	if err != nil {
		return nil, err
	}
	for _, free := range freeAirdrops {
		allAirdrops = append(allAirdrops, free)
	}

	paidAirdrops, err := GetAllAirdropPaid()
	if err != nil {
		return nil, err
	}
	for _, paid := range paidAirdrops {
		allAirdrops = append(allAirdrops, paid)
	}

	return allAirdrops, nil
}

func UpdateAirdropByID(col string, id primitive.ObjectID, name string, task string, link string) error {
	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"name": name,
			"task": task,
			"link": link,
		},
	}

	result, err := database.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateAirdrop: %v\n", err)
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}

	return nil
}

func DeleteAirdropByID(col string, id primitive.ObjectID) error {
	collection := database.Collection(col)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", id.Hex(), err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", id.Hex())
	}

	return nil
}
