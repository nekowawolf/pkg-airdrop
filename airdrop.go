package airdrop

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AirdropFree struct {
    ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name  string             `json:"name,omitempty" bson:"name,omitempty"`
    Task  string             `json:"task,omitempty" bson:"task,omitempty"`
    Link  string             `json:"link,omitempty" bson:"link,omitempty"`
}

type AirdropPaid struct {
    ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name  string             `json:"name,omitempty" bson:"name,omitempty"`
    Task  string             `json:"task,omitempty" bson:"task,omitempty"`
    Link  string             `json:"link,omitempty" bson:"link,omitempty"`
}
