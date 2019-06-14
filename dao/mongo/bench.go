package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)


func GetAllForum_view_sortById(currentPage int) interface{} {
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("forums_view").Find(
		context.TODO(),
		bson.M{},
		&options.FindOptions{
			Sort:  bson.M{"id": -1},
			Skip:  &skip,
			Limit: &perPage,
		})

	return many(cursor, err)
}

func GetAllForum_view(currentPage int) interface{} {
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("forums_view").Find(
		context.TODO(),
		bson.M{},
		&options.FindOptions{
			//Sort:  bson.M{"id": -1},
			Skip:  &skip,
			Limit: &perPage,
		})

	return many(cursor, err)
}

func GetAllForum_normal(currentPage int) interface{} {
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("forums").Find(context.TODO(), bson.M{}, &options.FindOptions{
		Sort:  bson.M{"id": -1},
		Skip:  &skip,
		Limit: &perPage,
		Projection: bson.M{
			"_id":           0,
			"id":            1,
			"idol_id":       1,
			"user_id":       1,
			"created_time":  1,
			"title":         1,
			"content":       1,
			"images":        1,
			"thumb":         1,
			"forum_picture": 1,
		},
	})

	return many(cursor, err)
}

func GetForumById_normal(id int) interface{} {
	res := mdb.Collection("forums").FindOne(
		context.TODO(),
		bson.M{"id": id},
		&options.FindOneOptions{
			Projection: bson.M{
				"_id":           0,
				"id":            1,
				"idol_id":       1,
				"user_id":       1,
				"created_time":  1,
				"title":         1,
				"content":       1,
				"images":        1,
				"thumb":         1,
				"forum_picture": 1,
			},
		},
	)
	if res.Err() != nil {
		log.Println(res.Err())
		return nil
	}
	return one(res)
}

func GetForumById_view(id int) interface{} {
	res := mdb.Collection("forums_view").FindOne(
		context.TODO(),
		bson.M{"id": id},
	)
	if res.Err() != nil {
		log.Println(res.Err())
		return nil
	}
	return one(res)
}
