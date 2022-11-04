package mongorepo

import (
	"cicd-controller/help"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var databaseName = "cicd_monitor"

func CreateConnection() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + "/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	help.MyPanic(err)
	mongoClient = client
	err = mongoClient.Ping(ctx, readpref.Primary())
	help.MyPanic(err)
}

func InsertOrUpdate(collectionName string, data interface{}, filter map[string]string) {
	collection := mongoClient.Database(databaseName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	opts := options.Update().SetUpsert(true)
	defer cancel()
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"name": filter["name"]},
		bson.D{
			{"$set", toDoc(data)},
		},
		opts,
	)
	help.MyPanic(err)
}

func Insert(collectionName string, data []interface{}) {
	collection := mongoClient.Database(databaseName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertMany(ctx, data)
	help.MyPanic(err)
}

func Count(collectionName string, filter WithFilter) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := mongoClient.Database(databaseName).Collection(collectionName)

	query := buildQuery(filter)
	count, err := collection.CountDocuments(ctx, query)
	help.MyPanic(err)
	return count
}

func FindAll(collectionName string, output interface{}, filter WithFilter) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := mongoClient.Database(databaseName).Collection(collectionName)

	query := buildQuery(filter)
	opts := options.Find().
		SetSort(bson.D{{"rating", -1}}).
		SetLimit(int64(filter.Size)).
		SetSkip(int64((filter.Page - 1) * filter.Size))
	cursor, err := collection.Find(ctx, query, opts)
	help.MyPanic(err)
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &output)
	help.MyPanic(err)
}

func buildQuery(filter WithFilter) bson.M {
	query := bson.M{}
	if len(filter.Name) > 0 {
		query = bson.M{"name": bson.M{"$regex": ".*" + filter.Name + ".*"}}
	}
	return query
}

func toDoc(v interface{}) (doc *bson.D) {
	data, err := bson.Marshal(v)
	help.MyPanic(err)
	err = bson.Unmarshal(data, &doc)
	help.MyPanic(err)
	return
}
