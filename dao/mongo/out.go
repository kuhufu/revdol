package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	. "revdol/config"
	"time"
)

var mdb *mongo.Database
var perPage int64 = 10
var idolNum int64 = 6

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Config.Mongo.URL))
	if err != nil {
		log.Println(err)
		return
	}
	mdb = client.Database("revdol")
}

func GetForumCount(id int, currentPage int) interface{} {
	perPage := perPage * 2
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("idols_forum_count").Find(context.TODO(),
		bson.M{"idol_id": id},
		&options.FindOptions{
			Sort:  bson.M{"date": -1},
			Skip:  &skip,
			Limit: &perPage,
			Projection: bson.M{
				"_id":   0,
				"count": 1,
				"date":  1,
			},
		},
	)
	return many(cursor, err)
}

func GetAllIdolForumCount(currentPage int) interface{} {
	perPage := idolNum
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("idols_forum_count").Find(context.TODO(),
		bson.M{},
		&options.FindOptions{
			Sort: bson.M{
				"date":  -1,
				"count": -1,
			},
			Skip:  &skip,
			Limit: &perPage,
			Projection: bson.M{
				"_id":     0,
				"idol_id": 1,
				"count":   1,
				"date":    1,
			},
		},
	)
	return many(cursor, err)
}

func GetFansNumById(id int, currentPage int) interface{} {
	perPage := perPage * 2
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("idols_meta").Find(context.TODO(),
		bson.M{"id": id},
		&options.FindOptions{
			Sort:  bson.M{"date": -1},
			Skip:  &skip,
			Limit: &perPage,
			Projection: bson.M{
				"_id":      0,
				"fans_num": 1,
				"date":     1,
			},
		},
	)

	return many(cursor, err)
}

func GetPopularNumById(id int, currentPage int) interface{} {
	perPage := perPage * 2
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("idols_meta").Find(context.TODO(),
		bson.M{"id": id},
		&options.FindOptions{
			Sort:  bson.M{"date": -1},
			Skip:  &skip,
			Limit: &perPage,
			Projection: bson.M{
				"_id":         0,
				"id":          0,
				"popular_num": 1,
				"date":        1,
			},
		},
	)

	return many(cursor, err)
}

func GetAllIdolMeta(currentPage int) interface{} {
	perPage := int64(idolNum)
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("idols_meta").Find(context.TODO(),
		bson.M{},
		&options.FindOptions{
			Sort:  bson.M{"date": -1, "id": 1},
			Skip:  &skip,
			Limit: &perPage,
			Projection: bson.M{
				"_id": 0,
			},
		},
	)

	return many(cursor, err)
}

func GetIdolMetaById(id, currentPage int) interface{} {
	perPage := perPage * 2
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("idols_meta").Find(
		context.TODO(),
		bson.M{"id": id},
		&options.FindOptions{
			Sort:  bson.M{"date": -1},
			Skip:  &skip,
			Limit: &perPage,
			Projection: bson.M{
				"_id": 0,
				"id":  0,
			},
		},
	)

	return many(cursor, err)
}

func GetForumById(id int) interface{} {
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
		})
	if res.Err() != nil {
		log.Println(res.Err())
		return nil
	}
	return one(res)
}

func GetAllForum(currentPage int) interface{} {
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("forums").Find(
		context.TODO(),
		bson.M{},
		&options.FindOptions{
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

func GetAllUser(currentPage int) interface{} {
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("users").Find(context.TODO(), bson.M{}, &options.FindOptions{
		Sort:  bson.M{"id": -1},
		Skip:  &skip,
		Limit: &perPage,
	})

	return many(cursor, err)
}

func GetUserById(id int) interface{} {
	res := mdb.Collection("users").FindOne(
		context.TODO(),
		bson.M{"id": id})

	return one(res)
}

func GetUserContributeById(id int) interface{} {
	cursor, err := mdb.Collection("users").Find(context.TODO(),
		bson.M{"user_id": id},
		&options.FindOptions{
			Sort: bson.M{"idol_id": 1},
		})

	return many(cursor, err)
}

func GetAllUserContribute(currentPage int) interface{} {
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("contributes").Find(context.TODO(),
		bson.M{},
		&options.FindOptions{
			Sort:  bson.M{"point": -1},
			Skip:  &skip,
			Limit: &perPage,
		})

	return many(cursor, err)
}

func GetIdolById(id int) interface{} {

	res := mdb.Collection("idols").FindOne(
		context.TODO(),
		bson.M{"id": id},
	)

	return one(res)
}

func GetAllIdol() interface{} {
	cursor, err := mdb.Collection("idols").Find(context.TODO(), bson.M{})
	return many(cursor, err)
}

func many(cursor *mongo.Cursor, err error) interface{} {
	if err != nil {
		log.Println(err)
		return nil
	}
	var result []bson.M
	//ctx, _ := context.WithTimeout(context.TODO(), time.Millisecond*100)
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
