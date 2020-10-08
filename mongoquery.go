package mongoquery

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Query returns datapoints with some secnod in the past
func Query(myCollection string, seconds int) int {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost/mydb"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	collection := client.Database("mydb").Collection(myCollection)
	// filter := bson.D{bson.E{Key: "_id", Value: bson.D{bson.E{Key: "$gt", Value: 2010081956011234}}}}
	filter := bson.D{{"_id", bson.D{{"$gt", 2010081956011234}}}}
	cursor, err := collection.Find(ctx, filter)
	var entry []bson.M
	if err = cursor.All(ctx, &entry); err != nil {
		log.Fatal(err)
	}
	fmt.Println(entry)

	return cursor.RemainingBatchLength()
}
