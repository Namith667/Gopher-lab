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

func main() {
	 if err := godotenv.Load(); err != nil {
	 	log.Println("No .env file found")
	 }

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

		//fmt.Println(client)

	for {
		displayMenu()	
		var ch int
		fmt.Scanln(&ch)
		switch ch {
		case 1:
				fmt.Println("-------creating----------")
				create(client)					
		case 2:				
			fmt.Println("------Read-------########")
			fetchAndDisplayData(client, "sample_db", "users")				
		case 3:
			fmt.Println("----------UPDATING------------")
			update(client)
		case 4: 
			fmt.Println("------------Deleting----------")				
			deleteRecords(client,"sample_db","users")
		case 5: 
			fmt.Println("Exiting!!")
			os.Exit(0)
					
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func displayMenu(){
	fmt.Println("1. Create")
    fmt.Println("2. Read")
    fmt.Println("3. Update")
    fmt.Println("4. Delete")
    fmt.Println("5. Exit")
    fmt.Print("Enter your choice: ")
}

func create(client *mongo.Client){

	coll := client.Database("sample_db").Collection("users")

	docs := []interface{}{
		bson.D{{"firstName", "Erik"}, {"age", 27}},
		bson.D{{"firstName", "Mohammad"}, {"lastName", "Ahmad"}, {"age", 10}},
		bson.D{{"firstName", "Todd"}},
		bson.D{{"firstName", "Juan"}, {"lastName", "Pablo"}},
	}
	
	result, err := coll.InsertMany(context.TODO(), docs)
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

func update(client *mongo.Client) {
	coll := client.Database("sample_db").Collection("users")
	filter := bson.D{{"firstName", "Erik"}}
	update := bson.D{{"$set", bson.D{{"lastName", "UPDATED!!!"},}},} 

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
			fmt.Println("Error updating document:", err)
	} else {
			fmt.Println("Document updated successfully:", result)
			 // Fetch the updated document
			var updatedDoc bson.M
			err = coll.FindOne(context.TODO(), filter).Decode(&updatedDoc)
			if err != nil {
				fmt.Println("Error fetching updated document:", err)
			} else {
				fmt.Println("Updated document:", updatedDoc)
			}
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	fmt.Printf("%s\n", jsonData)
	fmt.Println("!!!!!!!!!!UPDATEd")
}

func fetchAndDisplayData(client *mongo.Client,dbName,collectionName string){	
	// access db and collection
    coll := client.Database(dbName).Collection(collectionName)

    cur, err := coll.Find(context.TODO(), bson.M{})
    if err != nil {
        log.Fatalf("error fetching documents: %v", err)
    }
    defer cur.Close(context.TODO())

    // check if there are any results
    if !cur.Next(context.TODO()) {
        fmt.Println("No documents found")
        return
    }
    // Iterate over to display results
    for {
        var result bson.M
        err := cur.Decode(&result)
        if err != nil {
            log.Fatalf("error decoding document: %v", err)
        }

        if len(result) == 0 {
            fmt.Println("No data in document")
        } else {
            fmt.Println("Document:")
            for key, value := range result {
                fmt.Printf("%s: %v\n", key, value)
            }
        }
        fmt.Println()

        if !cur.Next(context.TODO()) {
            break
        }
    }

    if err := cur.Err(); err != nil {
        log.Fatalf("cursor error: %v", err)
    }
}

func deleteRecords(client *mongo.Client,dbName,collectionName string){
	collection  := client.Database(dbName).Collection(collectionName)
	filter := bson.D{{}}

	//res, err := collection.DropAll(context.TODO())
	
	res, err := collection.DeleteMany(context.TODO(),filter)
	if err != nil {
		fmt.Println("Error dropping database:", err)
	} else {
		fmt.Println("Database dropped successfully. Number of documents deleted:", res.DeletedCount)
	}
}




