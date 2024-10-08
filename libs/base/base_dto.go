package base

import "go.mongodb.org/mongo-driver/bson"

type BaseDto struct {
	Count int32    `json:"count"`
	Items []bson.D `json:"items"`
}
