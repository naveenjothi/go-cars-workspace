package utils

import "go.mongodb.org/mongo-driver/bson"

var CountStage = bson.D{{Key: "$count", Value: "total_documents"}}

var LimitStage = bson.D{{Key: "$limit", Value: 10}}
