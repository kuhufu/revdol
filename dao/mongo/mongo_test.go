package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"revdol/util"
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
	res := GetForumCount(4, 1)
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
	res := GetAllForum(2)
	util.Pretty(res)
}
func TestGetPopularNumById(t *testing.T) {

}