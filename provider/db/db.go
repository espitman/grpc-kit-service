package db

import (
	"context"
	"grpc-kit-service/provider/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Collection *mongo.Collection

var ctx, _ = context.WithTimeout(context.Background(), 15*time.Second)

func Connect() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.GetString("db.uri")))
	if err != nil {
		panic(err)
	}
	Collection = Client.Database(config.GetString("db.name")).Collection(config.GetString("db.collection"))
	makeIndexes()
}

func makeIndexes() {
}

// func makeUserUniqueEmailIndex() {
// 	mod := mongo.IndexModel{
// 		Keys: bson.M{
// 			"email": 1,
// 		}, Options: options.Index().SetUnique(true),
// 	}
// 	_, err := Collection.Indexes().CreateOne(ctx, mod)
// 	if err != nil {
// 		fmt.Println("Indexes().CreateOne() ERROR:", err)
// 	}
// }

// indexName, err := mgm.Collection.Indexes().CreateOne(
// 	context.Background(),
// 	mongo.IndexModel{
// 		Keys:    bson.D{{Key: "hostname", Value: 1}},
// 		Options: options.Index().SetUnique(true),
// 	},
// )
