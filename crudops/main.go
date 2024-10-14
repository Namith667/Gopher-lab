package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

func main() {

	//error logging implemented instead of fmt
	if err := godotenv.Load(); err != nil {
		log.Fatal("error in connecting database : ", err)
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set the 'MONGODB_URI' environment variable. ")
	}

	db, err := NewDatabase(uri)

	if err != nil {
		log.Fatal("Error Connecting to Database ", err)
	}
	defer db.Disconnect()

	for {
		displayMenu()
		var ch int
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			fmt.Println("-------Creating Document(s)----------")

			var numDocs int
			fmt.Print("Enter the number of documents you want to create: ")
			fmt.Scanln(&numDocs)

			docs := make([]interface{}, 0, numDocs)

			for i := 0; i < numDocs; i++ {
				fmt.Printf("\n--- Document %d ---\n", i+1)

				// Create a BSON document to hold key-value pairs
				doc := bson.D{}

				var firstName string
				fmt.Print("Enter firstName: ")
				fmt.Scanln(&firstName)
				doc = append(doc, bson.E{Key: "firstName", Value: firstName})

				var lastName string
				fmt.Print("Enter lastName (optional): ")
				fmt.Scanln(&lastName)
				if lastName != "" {
					doc = append(doc, bson.E{Key: "lastName", Value: lastName})
				}

				var age int
				fmt.Print("Enter age (optional, enter 0 to skip): ")
				fmt.Scanln(&age)
				if age != 0 {
					doc = append(doc, bson.E{Key: "age", Value: age})
				}

				docs = append(docs, doc)
			}

			result, err := db.Create("users", docs)
			if err != nil {
				log.Fatalf("Error creating documents: %v", err)
			}

			// Pretty print the result
			jsonData, _ := json.MarshalIndent(result, "", "    ")
			fmt.Printf("%s\n", jsonData)

		case 2:
			fmt.Println("------Read-------########")
			users, err := db.Read("users")
			if err != nil {
				log.Println("error Reading data ", err)
			}
			for _, user := range users {
				jsonData, _ := json.MarshalIndent(user, "", "   ")
				fmt.Printf("%s\n ", jsonData)
			}

		case 3:
			fmt.Println("----------UPDATING------------")

			// Input filter field and value from user
			var filterField, filterValue, updateField, updateValue string
			fmt.Print("Enter the field to be changed: ")
			fmt.Scanln(&filterField)
			fmt.Print("Enter the value to be changes: ")
			fmt.Scanln(&filterValue)

			// Input update field and value from user
			fmt.Print("Enter the field to update: ")
			fmt.Scanln(&updateField)
			fmt.Print("Enter the new value: ")
			fmt.Scanln(&updateValue)

			// Create the filter and update objects dynamically
			filter := bson.D{{filterField, filterValue}}
			update := bson.D{{"$set", bson.D{{updateField, updateValue}}}}

			result, err := db.Update("users", filter, update)
			if err != nil {
				log.Println("Error updating document: ", err)
			} else {
				jsonData, _ := json.MarshalIndent(result, "", "   ")
				fmt.Printf("UPDATED! %s\n", jsonData)
			}

		case 4:
			fmt.Println("------------Deleting----------")

			// Input filter field and value from user
			var filterField, filterValue string
			fmt.Print("Enter the field to filter by: ")
			fmt.Scanln(&filterField)
			fmt.Print("Enter the value to filter by: ")
			fmt.Scanln(&filterValue)

			// Create the filter dynamically
			filter := bson.D{{filterField, filterValue}}

			result, err := db.Delete("users", filter)
			if err != nil {
				log.Fatalf("Error deleting documents: %v", err)
			} else {
				fmt.Printf("Documents deleted: %d\n", result.DeletedCount)
			}

		case 5:
			fmt.Println("Exiting!!")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice")
		}
	}
}

// for db connection
func NewDatabase(uri string) (*Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err //review
	}
	return &Database{
		Client: client,
	}, nil
}

// db disconnect
func (db *Database) Disconnect() error {
	return db.Client.Disconnect(context.TODO())
}

// displayMenu
func displayMenu() {
	fmt.Println("1. Create")
	fmt.Println("2. Read")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
}

// create docs
func (db *Database) Create(collectionName string, docs []interface{}) (*mongo.InsertManyResult, error) {
	coll := db.Client.Database("sample_db").Collection(collectionName)
	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// read
func (db *Database) Read(collectionName string) ([]bson.M, error) {
	// access db and collection
	coll := db.Client.Database("sample_db").Collection(collectionName)
	cursor, err := coll.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil

}

// ipdate doc
func (db *Database) Update(collectionName string, filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	coll := db.Client.Database("sample_db").Collection(collectionName)
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error updating document:", err)
	}
	return result, nil

}

// delete matching docs
func (db *Database) Delete(collectionName string, filter bson.D) (*mongo.DeleteResult, error) {
	collection := db.Client.Database("sample_db").Collection(collectionName)
	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Println("Error dropping database:", err)
	}
	return result, nil
}

func errHandler(e error) {
	if e != nil {
		log.Print(e)
	}
}

