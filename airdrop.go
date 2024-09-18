package airdrop

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AirdropFree struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Task      string             `json:"task,omitempty" bson:"task,omitempty"`
	Link      string             `json:"link,omitempty" bson:"link,omitempty"`
	Level     string             `json:"level,omitempty" bson:"level,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type AirdropPaid struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Task      string             `json:"task,omitempty" bson:"task,omitempty"`
	Link      string             `json:"link,omitempty" bson:"link,omitempty"`
	Level     string             `json:"level,omitempty" bson:"level,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}