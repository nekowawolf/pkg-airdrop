package airdrop

import (
	"context"
	"errors"
	"time"
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

func InsertAirdropFree(name, task, link, level, status, backed, funds, supply, marketCap, vesting, linkClaim string, price float64, usdIncome int) (interface{}, error) {
	var endedAt *time.Time
	if status == "ended" {
		now := time.Now()
		endedAt = &now 
	}

	freeAirdrop := AirdropFree{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Task:      task,
		Link:      link,
		Level:     level,
		Status:    status,
		Backed:    backed,
		Funds:     funds,
		Supply:    supply,       
		MarketCap: marketCap,
		Vesting:   vesting,
		LinkClaim: linkClaim,
		Price:     price,
		USDIncome: usdIncome,
		CreatedAt: time.Now(),
		EndedAt:   endedAt,        
	}
	return InsertOneDoc("airdrop_free", freeAirdrop)
}

func InsertAirdropPaid(name, task, link, level, status, backed, funds, supply, marketCap, vesting, linkClaim string, price float64, usdIncome int) (interface{}, error) {
	var endedAt *time.Time
	if status == "ended" {
		now := time.Now()
		endedAt = &now 
	}

	paidAirdrop := AirdropPaid{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Task:      task,
		Link:      link,
		Level:     level,
		Status:    status,
		Backed:    backed,
		Funds:     funds,
		Supply:    supply,         
		MarketCap: marketCap,
		Vesting:   vesting,
		LinkClaim: linkClaim,
		Price:     price,
		USDIncome: usdIncome,
		CreatedAt: time.Now(),
		EndedAt:   endedAt,        
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

func GetAirdropFreeByName(name string) ([]AirdropFree, error) {
	collection := database.Collection("airdrop_free")
	filter := bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("GetAirdropFreeByName Find: %v", err)
	}
	var airdrops []AirdropFree
	if err = cursor.All(context.TODO(), &airdrops); err != nil {
		return nil, fmt.Errorf("GetAirdropFreeByName All: %v", err)
	}
	return airdrops, nil
}

func GetAirdropPaidByName(name string) ([]AirdropPaid, error) {
	collection := database.Collection("airdrop_paid")
	filter := bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("GetAirdropPaidByName Find: %v", err)
	}
	var airdrops []AirdropPaid
	if err = cursor.All(context.TODO(), &airdrops); err != nil {
		return nil, fmt.Errorf("GetAirdropPaidByName All: %v", err)
	}
	return airdrops, nil
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

func UpdateAirdropFreeByID(id primitive.ObjectID, name, task, link, level, status, backed, funds, supply, marketCap, vesting, linkClaim string, price float64, usdIncome int) error {
	collection := "airdrop_free"
	filter := bson.M{"_id": id}

	updateFields := bson.M{
		"name":       name,
		"task":       task,
		"link":       link,
		"level":      level,
		"status":     status,
		"backed":     backed,
		"funds":      funds,
		"supply":     supply,       
		"market_cap": marketCap,
		"vesting":    vesting,
		"link_claim": linkClaim,
		"price":      price,
		"usd_income": usdIncome,
	}

	if status == "ended" {
		now := time.Now()
		updateFields["ended_at"] = now
	}

	update := bson.M{
		"$set": updateFields,
	}

	result, err := database.Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("UpdateAirdropFreeByID: %v", err)
	}

	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}

	return nil
}

func UpdateAirdropPaidByID(id primitive.ObjectID, name, task, link, level, status, backed, funds, supply, marketCap, vesting, linkClaim string, price float64, usdIncome int) error {
	collection := "airdrop_paid"
	filter := bson.M{"_id": id}

	updateFields := bson.M{
		"name":       name,
		"task":       task,
		"link":       link,
		"level":      level,
		"status":     status,
		"backed":     backed,
		"funds":      funds,
		"supply":     supply,       
		"market_cap": marketCap,
		"vesting":    vesting,
		"link_claim": linkClaim,
		"price":      price,
		"usd_income": usdIncome,
	}

	if status == "ended" {
		now := time.Now()
		updateFields["ended_at"] = now
	}

	update := bson.M{
		"$set": updateFields,
	}

	result, err := database.Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("UpdateAirdropPaidByID: %v", err)
	}

	if result.ModifiedCount == 0 {
		return errors.New("no data has been changed with the specified ID")
	}

	return nil
}

func DeleteAirdropFreeByID(id primitive.ObjectID) error {
	collection := database.Collection("airdrop_free")
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s in airdrop_free: %s", id.Hex(), err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found in airdrop_free", id.Hex())
	}

	return nil
}

func DeleteAirdropPaidByID(id primitive.ObjectID) error {
	collection := database.Collection("airdrop_paid")
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s in airdrop_paid: %s", id.Hex(), err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found in airdrop_paid", id.Hex())
	}

	return nil
}