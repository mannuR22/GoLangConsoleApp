package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/mannuR22/PrecizeGoLang.git/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

func getMongoClient() (*mongo.Client, error) {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb://localhost:27017/").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// InsertRecord
func InsertRecord(report models.Report) (string, error) {

	client, err := getMongoClient()
	if err != nil {
		return "dao/InsertRecord(): Error Connecting to DB", err
	}
	defer client.Disconnect(context.Background())

	db := client.Database("taskdb")

	_, err = db.Collection("Records").InsertOne(context.Background(), report)
	if err != nil {
		return "dao/InsertRecord(): Error Occured while inserting record", err
	}

	return "Successfully Inserted Doc into db.", nil
}

// GetAllRecord
func GetAllRecords() ([]models.Report, error) {

	client, err := getMongoClient()
	var reports []models.Report
	// var userList models.ListModel
	if err != nil {
		return nil, err
	}

	defer client.Disconnect(context.Background())

	db := client.Database("taskdb")

	// Get a handle to the "users" collection
	// lists := db.Collection("lists")
	// Insert a new document

	cursor, err := db.Collection("Records").Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &reports); err != nil {
		return nil, err
	}

	return reports, nil
}

func GetRecordWithName(name string) (models.Report, error) {
	var report models.Report
	client, err := getMongoClient()
	fmt.Println("hello")
	// var re
	if err != nil {
		return models.Report{}, err
	}

	defer client.Disconnect(context.Background())

	db := client.Database("taskdb")
	// Insert a new document

	filter := bson.M{"name": name}

	err = db.Collection("Records").FindOne(context.Background(), filter).Decode(&report)
	if err != nil {
		return models.Report{}, err
	}

	return report, nil
}

// UpdateRecord
func UpdateRecord(name string, score int) (string, error) {
	var report models.Report
	client, err := getMongoClient()
	fmt.Println("hello")
	// var re
	if err != nil {
		return "dao/UpdateRecord(): Error Occured while creating mongo client.", err
	}

	defer client.Disconnect(context.Background())

	db := client.Database("taskdb")
	// Insert a new document

	filter := bson.M{"name": name}

	err = db.Collection("Records").FindOne(context.Background(), filter).Decode(&report)

	if err != nil {
		return "dao/UpdateRecord(): Error Occured while fetching record from db.", err
	}

	report.SATScore = score

	update := bson.M{"$set": report}

	// Update document
	_, err = db.Collection("Records").UpdateOne(context.Background(), filter, update)
	if err != nil {

		return "dao/UPdateRecord: Error Occured while Updating record into db.", err
	}

	return "Successfully updated data in db", nil

}

// DeleteRecorc

func DeleteRecord(name string) (string, error) {
	client, err := getMongoClient()

	if err != nil {
		return "dao/DeleteRecord: Error Occured while connecting to mongodb.", err
	}

	defer client.Disconnect(context.Background())

	db := client.Database("taskdb")

	// Insert a new document

	filter := bson.M{"name": name}

	_, err = db.Collection("Records").DeleteOne(context.TODO(), filter)
	if err != nil {
		return "dao/DeleteRecord: Error Occured while deleting Record from db.", err
	}

	return "Successfully Deleted Record from db.", nil

}
