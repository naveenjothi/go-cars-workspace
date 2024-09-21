package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MigrationModel struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Version       string             `bson:"version" json:"version"`
	Description   string             `bson:"description" json:"description"`
	AppliedAt     time.Time          `bson:"applied_at" json:"applied_at"`
	Status        string             `bson:"status" json:"status"`
	ExecutionTime string             `bson:"execution_time,omitempty" json:"execution_time"`
}
