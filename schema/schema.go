package schema

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Kit struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (t *Kit) Unmarshal(entry interface{}) Kit {
	obj, _ := json.Marshal(entry)
	response := Kit{}
	_ = json.Unmarshal(obj, &response)
	return response
}
