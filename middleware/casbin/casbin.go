package casbin

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	etcdwatcher "github.com/kuhufu/etcd3-watcher"
	. "github.com/kuhufu/revdol/config"
	"github.com/kuhufu/revdol/constant"
	"github.com/kuhufu/revdol/model"
	"log"
	"time"
)

var enforcer *casbin.SyncedEnforcer

func init() {
	enforcer = casbin.NewSyncedEnforcer(Config.Casbin.Model, Config.Casbin.Policy)
	enforcer.StartAutoLoadPolicy(time.Second * 30)

	if Config.Etcd.URL == "" {
		return
	}
	w := etcdwatcher.NewWatcher(Config.Etcd.URL)
	enforcer.SetWatcher(w)

	w.SetUpdateCallback(func(s string) {
		enforcer.LoadPolicy()
		log.Println("New Revision detected:", s)
	})
}

func GetEnforce() *casbin.SyncedEnforcer {
	return enforcer
}

func Check(c *gin.Context) bool {
	var account *model.Account
	if u, exists := c.Get(constant.AccountKey); !exists {
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
