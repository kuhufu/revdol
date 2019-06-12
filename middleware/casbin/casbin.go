package casbin

import (
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	etcdwatcher "github.com/kuhufu/etcd3-watcher"
	"log"
	. "revdol/config"
	"revdol/model"
	"time"
)

var fileEnforce *casbin.SyncedEnforcer
var dbEnforce *casbin.SyncedEnforcer

func init() {
	fileEnforce = casbin.NewSyncedEnforcer(Config.Casbin.Model, Config.Casbin.Policy)
	fileEnforce.StartAutoLoadPolicy(time.Second * 30)
	w := etcdwatcher.NewWatcher(Config.Etcd.URL)
	fileEnforce.SetWatcher(w)
	w.SetUpdateCallback(func(s string) {
		fileEnforce.LoadPolicy()
		log.Println("New Revision detected:", s)
	})

	dbEnforce = casbin.NewSyncedEnforcer(Config.Casbin.Model, gormadapter.NewAdapter("mysql", "kuhufu:12345@tcp(127.0.0.1:3306)/"))
	dbEnforce.StartAutoLoadPolicy(time.Second * 30)
}

func GetEnforce() *casbin.SyncedEnforcer {
	return fileEnforce
}

func Check(c *gin.Context) bool {
	var account *model.Account
	if u, exists := c.Get("account"); !exists {
		return false
	} else {
		account = u.(*model.Account)
	}

	e := GetEnforce()

	identity := account.Username
	role := account.Role
	path := c.Request.URL.Path
	method := c.Request.Method

	// e.Enforce(role, path, method) || e.Enforce(username, path, method)
	// 双重检查，首先检查该用户的角色是否有权限，再检查该用户是否有权限
	// 存在问题：
	// 必须保证用户的用户名不能是角色名
	// username == "account"|"admin"|"anonymous" 是不允许的，否则会造成越权
	// 解决办法：
	// 1. 自己创建 username 与角色名相同的用户
	// 2. 将策略分开存储，角色策略与用户策略不能保存再一起。
	// 3. 修改模型，这种方法角色策略与用户策略放在一起没问题
	// r = sub, role, obj, act
	// (r.role == p.role || g(r.sub, p.role)) && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)

	return e.Enforce(identity, role, path, method)
	//return e.Enforce(role, path, method) || e.Enforce(username, path, method)
}
