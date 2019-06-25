package config

import (
	"github.com/jinzhu/configor"
	"github.com/kuhufu/revdol/util"
	"log"
)

func init() {
	err := configor.Load(&Config, "config/config.yaml")

	if err != nil {
		log.Println("load config failure")
		return
	}
	if Config.Config.Print {
		util.Pretty(Config)
	}
}

var Config = struct {
	APPName string

	//影响的中间件：Secure，jwt
	Dev bool `default:"false"`

	//是否开启 auth
	EnableAuth bool `default:"false" json:"enable_auth" yaml:"enable_auth"`

	Config struct {
		Print bool `default:"false"`
	} `json:"config"`

	Gin struct {
		Release bool `default:"false" json:"release"`
		Color   bool `default:"false" json:"color"`
	} `json:"gin"`

	Mongo struct {
		URL string `required:"true" json:"url"`
	} `json:"mongo"`

	Redis struct {
		URL string `required:"true" json:"url"`
	} `json:"redis"`

	Gorm struct {
		LogMode  bool   `default:"false" json:"log_mode" yaml:"log_mode"`
		Provider string `required:"true" json:"provider"`
		URL      string `required:"true" json:"url"`
	} `json:"gorm"`

	Etcd struct {
		URL string
	} `json:"etcd"`

	Casbin struct {
		Model  string `required:"true"`
		Policy string `required:"true"`
	} `json:"casbin"`

	Cert struct {
		CertFile string `required:"true" json:"cert_file" yaml:"cert_file"`
		KeyFile  string `required:"true" json:"key_file" yaml:"key_file"`
	} `json:"cert"`

	HttpPort  string `default:":80" json:"http_port" yaml:"http_port"`
	HttpsPort string `default:":443" json:"https_port" yaml:"https_port"`
}{}
