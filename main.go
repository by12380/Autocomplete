package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/by12380/Autocomplete/configs"
	"github.com/by12380/Autocomplete/routers"
	"github.com/by12380/Autocomplete/services/trie"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/weekface/mgorus"
)

type config struct {
	AssetsPath         string `env:"ASSETS_PATH"`
	MongoDBServiceHost string `env:"AUTOCOMPLETE_MONGODB_SERVICE_HOST"`
	MongoDBServicePort string `env:"AUTOCOMPLETE_MONGODB_SERVICE_PORT"`
}

func main() {
	cfg := config{}
	env.Parse(&cfg)

	mongoDBUrl := "mongodb://" + os.Getenv("AUTOCOMPLETE_MONGODB_SERVICE_HOST") + ":" + os.Getenv("AUTOCOMPLETE_MONGODB_SERVICE_PORT")

	log := logrus.New()
	hook, err := mgorus.NewHooker(mongoDBUrl, "autocomplete", "logs")
	if err == nil {
		log.Hooks.Add(hook)
	} else {
		fmt.Print(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	endpoint := os.Getenv("AUTOCOMPLETE_MINIO_SERVICE_HOST") + ":" + os.Getenv("AUTOCOMPLETE_MINIO_SERVICE_PORT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.New()
	r.Use(configs.Logger(log), gin.Recovery())

	routers.InitAutocomplete(r.Group("/autocomplete"))

	bucketName := "drive"
	objectName := "trie-input.json"
	filePath := cfg.AssetsPath + "/trie-input.json"

	// Upload the zip file with FPutObject
	err = minioClient.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
		return
	}

	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalln(err)
		return
	}

	result := gjson.GetBytes(bytes, "items")

	// Build trie
	t := trie.GetInstance()
	for _, item := range result.Array() {
		textItem := trie.TextItem{
			Value:  item.Get("value").String(),
			Weight: int(item.Get("weight").Int()),
		}
		t.Add(&textItem)
	}

	r.Run()
}
