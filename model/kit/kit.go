package model_kit

import (
	"context"
	"grpc-kit-service/provider/db"
	"grpc-kit-service/schema"

	"github.com/espitman/protos-kit/kit"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.Background()

func CreateNewkit(req *kit.RequestCreate) (*kit.ResponseDetails, error) {
	newkit := (&schema.Kit{}).Unmarshal(req)
	res, err := db.Collection.InsertOne(ctx, newkit)
	if err != nil {
		return nil, err
	}
	var response = (&kit.ResponseDetails{}).Unmarshal(newkit)
	response.Id = res.InsertedID.(primitive.ObjectID).Hex()
	return &response, nil
}
