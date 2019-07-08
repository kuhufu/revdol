package mongo

import (
	"context"
	"fmt"
	"github.com/kuhufu/revdol/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI(Config.Mongo.URL))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	mdb = client.Database("revdol")
}

func TestGetForumCount(t *testing.T) {
	res := GetIdolForumCount(4, 1)
	fmt.Println(res)
}

func TestGetFansNumById(t *testing.T) {
	res := GetFansNumById(4, 1)
	fmt.Println(res)
}
func TestGetAllIdolMeta(t *testing.T) {
	res := GetAllIdolMeta(1)
	util.Pretty(res)
}

func TestGetForumById(t *testing.T) {
	res := GetForumById(1)
	fmt.Println(res)
}

func TestGetAllForum(t *testing.T) {
	res := GetAllForum(map[string]interface{}{
		"page": 1,
	})
	util.Pretty(res)
}
func TestGetPopularNumById(t *testing.T) {
	res := GetPopularNumById(4, 1)
	util.Pretty(res)
}

func TestSearchUser(t *testing.T) {
	res := SearchUser("鱿鱼丝", 1)
	util.Pretty(res)
}

func TestSource_SearchForumTitle(t *testing.T) {
	res := SearchForum("title", "好", 1)
	util.Pretty(res)
}

func TestGetUserForumCount(t *testing.T) {
	res := GetUserForumCount(1001, 1)
	util.Pretty(res)
}

func TestRunCommand(t *testing.T) {
	count := map[string]int{}
	mdb.RunCommand(context.TODO(), bson.M{"count": "forums"}).Decode(&count)
	fmt.Println(count)
}

//BenchmarkGetForumById-8   	    5000	    294830 ns/op
func BenchmarkGetForumById(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetForumById(17115)
	}
}

func BenchmarkGetAllForum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mdb.Collection("forums").CountDocuments(context.TODO(), bson.M{})
	}
}
