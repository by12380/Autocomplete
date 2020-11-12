package main

import (
	"fmt"
	"os"

	"github.com/by12380/Autocomplete/configs"
	"github.com/by12380/Autocomplete/dependencies/api-gateway/routers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
)

func main() {
	mongoDBUrl := "mongodb://" + os.Getenv("DEFAULT_MONGODB_SERVICE_HOST") + ":" + os.Getenv("DEFAULT_MONGODB_SERVICE_PORT")

	log := logrus.New()
	hook, err := mgorus.NewHooker(mongoDBUrl, "autocomplete", "logs")

	if err == nil {
		log.Hooks.Add(hook)
	} else {
		fmt.Print(err)
	}

	r := gin.New()
	r.Use(configs.Logger(log), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "success")
	})
	routers.InitAutocomplete(r.Group("/autocomplete"))
	r.Run()
}
