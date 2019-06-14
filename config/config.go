package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/tidwall/pretty"
	"log"
)

func init() {
	err := configor.Load(&Config, "config/config.yaml")

	if err != nil {
		log.Println("load config failure")
		return
	}

	if Config.Config.Print {
		bf := bytes.NewBuffer([]byte{})
		enc := json.NewEncoder(bf)
		enc.SetEscapeHTML(false)
		enc.Encode(Config)
		bs := pretty.Pretty(bf.Bytes())
		fmt.Println(string(bs))
	}
}

var Config = struct {
	APPName string `default:"revdol"`

	Dev bool `default:"true"`

	Config struct {
		Print bool `default:"false"`
	} `json:"config"`

	Gin struct {
		Release bool `default:"false" json:"release"`
		Color   bool `default:"false" json:"color"`
	} `json:"gin"`

	Mongo struct {
		URL string `json:"url"`
	} `json:"mongo"`

	Redis struct {
		URL string `required:"true" json:"url"`
	} `json:"redis"`

	Gorm struct {
		Log      bool   `default:"false" json:"log_mode"`
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
		Cert string `required:"true" json:"cert_file"`
		Key  string `required:"true" json:"key_file"`
	} `json:"cert"`

	Http_port  string `default:":80" json:"http_port"`
	Https_port string `default:":443" json:"https_port"`
}{}
