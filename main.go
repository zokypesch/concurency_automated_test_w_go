package main

import (
	"fmt"
	"net/http"

	"./jwt"

	hello "./pack/hello"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("server")  // name of config file (without extension)
	viper.AddConfigPath("config/") // path to look for the config file in
	err := viper.ReadInConfig()    // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	port := viper.Get("port") //os.Getenv("PORT")

	setup := viper.Get("setup")

	if setup.(string) == "true" {
		//do something
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware := jwt.Auth()

	// routing login
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/hello/execute", hello.Exec)

	http.ListenAndServe(":"+port.(string), r)
}
