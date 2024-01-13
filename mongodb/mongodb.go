package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBConnect() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://igbalmar:RYSfWClaOvDJLG6H@stocks-helper-db.crz6pyh.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func MondoDBInsertData() {
	var collectionName = "recipes"
	collection := client.Database(dbName).Collection(collectionName)

	/*      *** INSERT DOCUMENTS ***
	 *
	 * You can insert individual documents using collection.Insert().
	 * In this example, we're going to create 4 documents and then
	 * insert them all in one call with InsertMany().
	 */

	eloteRecipe := Recipe{
		Name:              "elote",
		Ingredients:       []string{"corn", "mayonnaise", "cotija cheese", "sour cream", "lime"},
		PrepTimeInMinutes: 35,
	}

	locoMocoRecipe := Recipe{
		Name:              "loco moco",
		Ingredients:       []string{"ground beef", "butter", "onion", "egg", "bread bun", "mushrooms"},
		PrepTimeInMinutes: 54,
	}

	patatasBravasRecipe := Recipe{
		Name:              "patatas bravas",
		Ingredients:       []string{"potato", "tomato", "olive oil", "onion", "garlic", "paprika"},
		PrepTimeInMinutes: 80,
	}

	friedRiceRecipe := Recipe{
		Name:              "fried rice",
		Ingredients:       []string{"rice", "soy sauce", "egg", "onion", "pea", "carrot", "sesame oil"},
		PrepTimeInMinutes: 40,
	}

	// Create an interface of all the created recipes
	recipes := []interface{}{eloteRecipe, locoMocoRecipe, patatasBravasRecipe, friedRiceRecipe}
	insertManyResult, err := collection.InsertMany(context.TODO(), recipes)
	if err != nil {
		fmt.Println("Something went wrong trying to insert the new documents:")
		panic(err)
	}

	fmt.Println(len(insertManyResult.InsertedIDs), "documents successfully inserted.\n")
}
