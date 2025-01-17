package mongoc

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDataStore struct {
	DB *mongo.Database
}

func Connect(mongoUri string, dbName string, dataStore *MongoDataStore) {

	client, err := mongo.Connect(options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	//defer func() {
	//	if err = client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	dataStore.DB = client.Database(dbName)

}
