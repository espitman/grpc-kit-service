package schema

import (
	"encoding/json"

	"github.com/Kamva/mgm"
)

type Kit struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
}

func (t *Kit) Unmarshal(entry interface{}) Kit {
	obj, _ := json.Marshal(entry)
	response := Kit{}
	_ = json.Unmarshal(obj, &response)
	return response
}
