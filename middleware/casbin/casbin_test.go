package casbin

import (
	"fmt"
	"github.com/casbin/casbin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func TestLoadFromDB(t *testing.T) {
	//a := gormadapter.NewAdapter("mysql", "kuhufu:12345@tcp(127.0.0.1:3306)/")
	e := casbin.NewEnforcer("model.conf", "policy.csv")
	e.LoadPolicy()
	if e.Enforce("user1", "admin", "/user/info", "GET") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	//e.SavePolicy()

}
