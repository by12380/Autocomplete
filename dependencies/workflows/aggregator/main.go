package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	endpoint := os.Getenv("DEFAULT_MINIO_SERVICE_HOST") + ":" + os.Getenv("DEFAULT_MINIO_SERVICE_PORT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called drive.
	bucketName := "drive"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	mongoDBUrl := "mongodb://" + os.Getenv("DEFAULT_MONGODB_SERVICE_HOST") + ":" + os.Getenv("DEFAULT_MONGODB_SERVICE_PORT")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDBUrl))
	defer client.Disconnect(ctx)

	logs := client.Database("autocomplete").Collection("logs")

	// specify a pipeline that will return the number of times each name appears in the collection
	// specify the MaxTime option to limit the amount of time the operation can run on the server
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$queryString"},
			{"count", bson.D{
				{"$sum", 1},
			}},
		}},
	}
	opts := options.Aggregate().SetMaxTime(5 * time.Second)
	cursor, err := logs.Aggregate(ctx, mongo.Pipeline{groupStage}, opts)
	if err != nil {
		log.Fatal(err)
	}

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	var items []bson.M
	for _, result := range results {
		items = append(items, bson.M{
			"value":  result["_id"],
			"weight": result["count"],
		})
	}
	file, _ := json.MarshalIndent(bson.M{
		"items": items,
	}, "", " ")

	_ = ioutil.WriteFile("trie-input.json", file, 0644)

	// Upload the zip file
	objectName := "trie-input.json"
	filePath := "trie-input.json"
	contentType := "application/json"

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
