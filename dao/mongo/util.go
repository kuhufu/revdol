package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func many(cursor *mongo.Cursor, err error) interface{} {
	if err != nil {
		log.Println(err)
		return nil
	}
	var result []bson.M
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func one(res *mongo.SingleResult) interface{} {
	if res.Err() != nil {
		log.Println(res.Err())
		return nil
	}
	result := bson.M{}
	err := res.Decode(&result)
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}
