package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	. "github.com/kuhufu/revdol/config"
	. "github.com/kuhufu/revdol/swagger"
	"log"
)

func init() {
	flag.StringVar(&Config.HttpPort, "p", Config.HttpPort, "http port")
	flag.StringVar(&Config.HttpsPort, "s", Config.HttpsPort, "http port")
	flag.Parse()

	if Config.Gin.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	if Config.Gin.Color {
		gin.ForceConsoleColor()
	} else {
		gin.DisableConsoleColor()
	}
}

// @title EnableSwagger revdol API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @schemes https http
// @Security ApiKeyAuth

// @host localhost
// @BasePath /v2

func main() {

	r := gin.Default()
	EnableSwagger(r)
	router(r)

	// HTTP
	go func() {
		log.Fatal(r.Run(Config.HttpPort))
	}()

	// HTTPS
	cert, key := certPath()
	log.Fatal(r.RunTLS(Config.HttpsPort, cert, key))
}

func certPath() (cert, key string) {
	return Config.Cert.CertFile, Config.Cert.KeyFile
}
