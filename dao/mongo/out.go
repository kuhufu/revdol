package mongo

import (
	"context"
	. "github.com/kuhufu/revdol/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math"
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

func GetIdolForumCount(id int, currentPage int) interface{} {
	limit := perPage * 2
	skip := skipNum(currentPage, limit)
	cursor, err := mdb.Collection("idols_forum_count").Find(context.TODO(),
		bson.M{"idol_id": id},
		&options.FindOptions{
			Sort:  bson.M{"date": -1},
			Skip:  &skip,
			Limit: &limit,
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
	skip := skipNum(currentPage, perPage)
	cursor, err := mdb.Collection("idols_forum_count").Find(context.TODO(),
		bson.M{},
		&options.FindOptions{
			Sort: bson.D{
				{"date", -1},
				{"count", -1},
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
	skip := skipNum(currentPage, perPage)
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
	skip := skipNum(currentPage, perPage)
	cursor, err := mdb.Collection("idols_meta").Find(context.TODO(),
		bson.M{"id": id},
		&options.FindOptions{
			Sort:  bson.M{"date": -1},
			Skip:  &skip,
			Limit: &perPage,
			Projection: bson.M{
				"_id":         0,
				"popular_num": 1,
				"date":        1,
			},
		},
	)

	return many(cursor, err)
}

func GetAllIdolMeta(currentPage int) interface{} {
	perPage := int64(idolNum)
	skip := skipNum(currentPage, perPage)
	cursor, err := mdb.Collection("idols_meta").Find(context.TODO(),
		bson.M{},
		&options.FindOptions{
			Sort: bson.D{
				{"date", -1},
				{"id", 1},
			},
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

func GetAllForum(params map[string]interface{}) interface{} {
	currentPage := 1
	if page, ok := params["page"]; ok {
		currentPage, ok = page.(int)
	}
	delete(params, "page")
	return queryForum(params, currentPage)
}

func queryForum(filter bson.M, currentPage int) interface{} {
	perPage := perPage
	skip := skipNum(currentPage, perPage)
	cursor, err := mdb.Collection("forums").Find(
		context.TODO(),
		filter,
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

	res := map[string]int{}
	err = mdb.RunCommand(context.TODO(), bson.D{
		{"count", "forums"},
		{"query", filter},
	}).Decode(&res)
	if err != nil {
		log.Println(err)
	}
	count := res["n"]
	forums := many(cursor, err)
	return map[string]interface{}{
		"_meta": map[string]interface{}{
			"count":       count,
			"totalPage":   int(math.Ceil(float64(count) / float64(perPage))),
			"currentPage": currentPage,
			"perPage":     perPage,
		},
		"items": forums,
	}
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
	res := mdb.Collection("users").FindOne(context.TODO(),
		bson.M{"uid": id},
		&options.FindOneOptions{
			Projection: bson.M{
				"_id":        0,
				"address":    1,
				"birth":      1,
				"id":         1,
				"uid":        1,
				"nickname":   1,
				"sex":        1,
				"headimg":    1,
				"created_at": 1,
				"updated_at": 1,
				"tel":        1,
				"area_code":  1,
				"status":     1,
			},
		},
	)

	return one(res)
}

func GetUserContributeById(id int) interface{} {
	cursor, err := mdb.Collection("contributes").Find(context.TODO(),
		bson.M{"user_id": id},
		&options.FindOptions{
			Sort: bson.M{"point": -1},
			Projection: bson.M{
				"_id":     0,
				"idol_id": 1,
				//"user_id":    1,
				"point":      1,
				"created_at": 1,
				"updated_at": 1,
			},
		})

	return many(cursor, err)
}

func GetAllUserContribute(currentPage int) interface{} {
	skip := skipNum(currentPage, perPage)
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

func SearchUser(keyWord string, currentPage int) interface{} {
	var perPage int64 = 10
	skip := skipNum(currentPage, perPage)
	cursor, err := mdb.Collection("users").Find(context.TODO(),
		bson.M{
			"nickname": bson.M{
				"$regex":   keyWord,
				"$options": "i",
			}},
		&options.FindOptions{
			Limit: &perPage,
			Skip:  &skip,
			Projection: bson.M{
				"_id":      0,
				"uid":      1, //用 uid 代替 id，id现在有问题
				"nickname": 1,
			},
		},
	)
	count := count(bson.D{
		{"count", "users"},
		{"query", bson.M{"nickname": bson.M{
			"$regex":   keyWord,
			"$options": "i",
		}}},
	})

	return map[string]interface{}{
		"_meta": pageInfo(count, perPage, currentPage),
		"items": many(cursor, err),
	}
}

func SearchForum(field, keyWord string, currentPage int) interface{} {
	var perPage int64 = 10
	skip := skipNum(currentPage, perPage)
	cursor, err := mdb.Collection("forums").Find(context.TODO(),
		bson.M{
			field: bson.M{
				"$regex":   keyWord,
				"$options": "i",
			}},
		&options.FindOptions{
			Limit: &perPage,
			Skip:  &skip,
			Sort: bson.M{
				"id": -1,
			},
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

	count := count(bson.D{
		{"count", "forums"},
		{"query", bson.M{
			field: bson.M{
				"$regex":   keyWord,
				"$options": "i",
			}}},
	})

	return map[string]interface{}{
		"_meta": pageInfo(count, perPage, currentPage),
		"items": many(cursor, err),
	}
}
func GetUserForumCount(id, currentPage int) interface{} {
	cursor, err := mdb.Collection("forums").Aggregate(context.TODO(), bson.A{
		bson.M{"$match": bson.M{
			"user_id": id,
		}},
		bson.M{"$group": bson.M{
			"_id":   "$idol_id",
			"count": bson.M{"$sum": 1},
		}},
		bson.M{"$sort": bson.M{
			"count": -1,
		}},
		bson.M{"$project": bson.M{
			"_id":     0,
			"idol_id": "$_id",
			"count":   1,
		}},
	})

	return many(cursor, err)
}

func groupForumCount(match, group bson.M, currentPage int) interface{} {
	skip := int64(currentPage-1) * perPage
	cursor, err := mdb.Collection("forums").Aggregate(context.TODO(),
		bson.A{
			bson.M{"$match": match},
			bson.M{"$group": group},
			bson.M{"$skip": skip},
			bson.M{"$limit": perPage},
			bson.M{"$sort": bson.M{"count": -1}},
		})

	return many(cursor, err)
}

func skipNum(currentPage int, perPage int64) int64 {
	return int64(currentPage-1) * perPage
}

func count(query bson.D) int {
	res := map[string]int{}
	err := mdb.RunCommand(context.TODO(), query).Decode(&res)
	if err != nil {
		log.Println(err)
		return -1
	}
	return res["n"]
}

func pageInfo(count int, perPage int64, currentPage int) interface{} {
	return map[string]interface{}{
		"count":       count,
		"totalPage":   int(math.Ceil(float64(count) / float64(perPage))),
		"currentPage": currentPage,
		"perPage":     perPage,
	}
}
