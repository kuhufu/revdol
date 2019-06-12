package mongoSource

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"revdol/model"
	"testing"
	"time"
)

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	mdb = client.Database("revdol")
}

func TestFind(t *testing.T) {

	cursor, err := mdb.Collection("forums").Find(
		ctx,
		bson.M{
			"idol_id": 4,
			"created_at": bson.M{"$gt": time.Now().Add(-2 * 24 * time.Hour),}},
		&options.FindOptions{
			Sort: bson.M{"created_at": -1},
		})
	if err != nil {
		log.Println(err)
	}
	arr := []map[string]interface{}{}
	cursor.All(ctx, &arr)
	fmt.Println(arr)
}

func TestFindOne(t *testing.T) {
	result := map[string]interface{}{}
	mdb.Collection("forums").FindOne(ctx, bson.M{"user_id": 1001}).Decode(&result)
	raws, e := mdb.Collection("forums").FindOne(ctx, bson.M{"user_id": 1001}).DecodeBytes()
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(result)
	fmt.Println(string(raws))

}

func TestInsertOne(t *testing.T) {
	var mapData interface{} = map[string]interface{}{"name":"kuhufu", "age":11}
	result, err := mdb.Collection("idols").InsertOne(ctx, mapData)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("insert map:", result)

	bytesData, _ := json.Marshal(mapData)
	result, err = mdb.Collection("idols").InsertOne(ctx, mapData)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("insert bytes:", string(bytesData))


	idol := model.Idol{ID:12, Nickname:"tamomo"}
	result, err = mdb.Collection("idols").InsertOne(ctx, idol)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("insert struct:", result)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)
}

func TestDistinct(t *testing.T) {
	arr := []map[string]interface{}{}
	distinct, err := mdb.Collection("forums").Distinct(ctx, "content", bson.M{"sort": -1})
	if err != nil {
		log.Fatal(err)
	}
	mdb.Collection("forums")
	log.Println(distinct)
	fmt.Println(arr)
}