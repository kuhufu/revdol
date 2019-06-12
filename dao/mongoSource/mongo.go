package mongoSource

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	. "revdol/config"
	"time"
)

var ctx context.Context
var mdb *mongo.Database

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Config.Mongo.URL))
	if err != nil {
		log.Fatal(err)
	}
	mdb = client.Database("revdol")
}

func AddForum(data []byte) {
	save(data, "forums")
}

func SaveUser(data []byte) {
	save(data, "users")
}

func SaveContribute(data []byte) {
	save(data, "contributes")
}

func SaveIdol(data []byte) {
	save(data, "idols")
}

func save(data []byte, collectionName string) {
	collection := mdb.Collection(collectionName)
	v := bson.M{}
	err := bson.UnmarshalExtJSON(data, false, &v)
	if err != nil {
		log.Println(err)
	}
	UpdateOrInsert(v, collection, "id")
}

func UpdateOrInsert(m bson.M, collection *mongo.Collection, id string) {
	c, _ := collection.CountDocuments(ctx, bson.M{id: m[id]})
	if c == 0 {
		_, err := collection.InsertOne(ctx, m)
		if err != nil {
			log.Println(err)
		}
	} else {
		_, err := collection.UpdateOne(ctx, bson.M{id: m[id]}, bson.M{"$set": m})
		if err != nil {
			log.Println(err)
		}
	}
}
