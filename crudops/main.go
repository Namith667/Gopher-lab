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

	if err != nil{ 
		log.Fatal("Error Connecting to Database ",err)
	}
	defer db.Disconnect()

	for {
		displayMenu()	
		var ch int
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			fmt.Println("-------creating----------")
			docs := []interface{}{
				bson.D{{"firstName", "Erik"}, {"age", 27}},
				bson.D{{"firstName", "Mohammad"}, {"lastName", "Ahmad"}, {"age", 10}},
				bson.D{{"firstName", "Todd"},{"age",19}},
				bson.D{{"firstName", "Juan"}, {"lastName", "Pablo"}},
			}
			result, err:= db.create("users",docs)	
			if err != nil{
				log.Fatalf("Error Creating Docs %v",err)
			}
			jsonData,_:=json.MarshalIndent(result,"","    ")	
			fmt.Printf("%s\n ",jsonData)			

		case 2:				
			fmt.Println("------Read-------########")
			users,err:= db.Read("users")
			if err != nil{
				log.Println("error Reading data ",err)
			}
			for _,user := range users{
				jsonData,_ :=json.MarshalIndent(user,"","   ")
				fmt.Printf("%s\n ",jsonData)
			}

		case 3:
			fmt.Println("----------UPDATING------------")
			filter := bson.D{{"firstName", "Erik"}}
			update := bson.D{{"$set", bson.D{{"lastName", "UPDATED!!!"}}}}
			result, err := db.Update("users",filter,update)
			if err != nil{
				log.Println("Error Updatingg doc! ",err)
			}
			jsonData,_:=json.MarshalIndent(result,"","   ")
			fmt.Printf("UPDATED! %s\n ",jsonData)

		case 4: 
			fmt.Println("------------Deleting----------")				
			filter := bson.D{{"age", bson.D{{"$lt", 20}}}} // Example filter
			result, err := db.Delete("users", filter)
			if err != nil {
				log.Fatalf("Error deleting documents: %v", err)
			}
			fmt.Printf("Documents deleted: %d\n", result.DeletedCount)

		case 5: 
			fmt.Println("Exiting!!")
			os.Exit(0)
			
		default:
			fmt.Println("Invalid choice")
		}
	}
}

//for db connection
func  NewDatabase(uri string)(*Database ,error){
	client, err := mongo.Connect(context.TODO(),options.Client().ApplyURI(uri))
	if err != nil{
		return nil,err																	//review
	}
	return &Database{
		Client: client,
	},nil
}

//db disconnect
func (db *Database) Disconnect() error{
	return db.Client.Disconnect(context.TODO())
}

//displayMenu
func displayMenu(){
	fmt.Println("1. Create")
    fmt.Println("2. Read")
    fmt.Println("3. Update")
    fmt.Println("4. Delete")
    fmt.Println("5. Exit")
    fmt.Print("Enter your choice: ")
}

//create docs
func (db *Database) create (collectionName string,docs []interface{}) (*mongo.InsertManyResult,error){

	coll := db.Client.Database("sample_db").Collection(collectionName)
	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		return nil, err
	}
	return result,nil
}

//read
func (db *Database) Read(collectionName string) ([]bson.M,error){	
	// access db and collection
	coll := db.Client.Database("sample_db").Collection(collectionName)
	cursor, err := coll.Find(context.TODO(), bson.M{})

	if err != nil {
	return nil, err
	}

	var results []bson.M
	if err := cursor.All(context.TODO(),&results); err != nil{
		return nil, err
	}
	return results,nil

}

//ipdate doc
func (db *Database) Update (collectionName string,filter bson.D,update bson.D ) (*mongo.UpdateResult,error){
	coll := db.Client.Database("sample_db").Collection(collectionName)
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
			log.Println("Error updating document:", err)
	} 
	return result,nil 
	
}

//delete matching docs
func (db *Database) Delete(collectionName string, filter bson.D) (*mongo.DeleteResult,error){
	collection  := db.Client.Database("sample_db").Collection(collectionName)
	result, err := collection.DeleteMany(context.TODO(),filter)
	if err != nil {
		log.Println("Error dropping database:", err)
	}
	return result,nil
}




