package db_helper

import (
	"context"
	"log"
	"os"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NewsItem struct {
	Url        string
	UploadTime int64
}

func News(soure string) []string {

	mongo_url := "mongodb://" + os.Getenv("MONGO_URL") + ":27017/"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	news_collection := client.Database("did_homa").Collection(soure)

	cursor, err := news_collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var news []NewsItem
	if err = cursor.All(context.TODO(), &news); err != nil {
		log.Fatal(err)
	}

	// sorting
	sort.Slice(news, func(i, j int) bool {
		return news[i].UploadTime > news[j].UploadTime
	})

	var final_slice []string
	for _, s := range news {
		final_slice = append(final_slice, s.Url)
	}

	return final_slice[:20]
}
