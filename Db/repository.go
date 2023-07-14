package Db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoTemplate struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func Init(url string, database string, collection string) MongoTemplate {
	mongoTemplate := MongoTemplate{}
	mongoTemplate.client = CreateClient(url)
	mongoTemplate.collection = mongoTemplate.client.Database(database).Collection(collection)
	return mongoTemplate
}

// InsertOne add one new document , return the value id in mongodb
// must add objectId
func (mongoTemplate *MongoTemplate) InsertOne(ctx context.Context, T interface{}) interface{} {
	insertValue, _ := mongoTemplate.collection.InsertOne(ctx, T)
	return insertValue.InsertedID
}

func (mongoTemplate *MongoTemplate) DeleteIneById(ctx context.Context, id string) int64 {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filed := bson.M{"_id": objectId}
	return mongoTemplate.DeleteOne(ctx, filed)
}

func (mongoTemplate *MongoTemplate) DeleteOne(ctx context.Context, criteria bson.M) int64 {
	deleteResult, _ := mongoTemplate.collection.DeleteOne(ctx, criteria)
	return deleteResult.DeletedCount
}

func (mongoTemplate *MongoTemplate) DeleteMany(ctx context.Context, criteria bson.M) int64 {
	deleteResult, _ := mongoTemplate.collection.DeleteMany(ctx, criteria)
	return deleteResult.DeletedCount
}

// FindOneById When using, _id must belong to primitive.ObjectID, if _id belongs to other types, please use your method
func (mongoTemplate *MongoTemplate) FindOneById(ctx context.Context, id string, T interface{}) *mongo.SingleResult {
	objectId, _ := primitive.ObjectIDFromHex(id)
	criteria := bson.M{"_id": objectId}
	return mongoTemplate.FindOne(ctx, criteria, T)
}

// FindOne T is the type of the value you want to get,will convert the value to T
// if exist many value,return first value
func (mongoTemplate *MongoTemplate) FindOne(ctx context.Context, criteria bson.M, T interface{}) *mongo.SingleResult {
	singleResult := mongoTemplate.collection.FindOne(ctx, criteria)
	err := singleResult.Decode(T)
	if err != nil {
		log.Fatal(err)
	}
	return singleResult
}

// FindMany if criteria := bson.M{} and opts == nil ,return all value,
// criteria is query condition,opts is Data display mode{sort,limit.....}
func (mongoTemplate *MongoTemplate) FindMany(ctx context.Context, criteria bson.M, T interface{}, opts ...*options.FindOptions) interface{} {
	cursor, _ := mongoTemplate.collection.Find(ctx, criteria, opts...)
	list := make([]interface{}, 0)
	for cursor.Next(ctx) {
		err := cursor.Decode(T)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, T)
	}
	return list
}

func (mongoTemplate *MongoTemplate) UpdateOneById(ctx context.Context, id string, update bson.M) int64 {
	objectId, _ := primitive.ObjectIDFromHex(id)
	criteria := bson.M{"_id": objectId}
	return mongoTemplate.UpdateOne(ctx, criteria, update)
}

func (mongoTemplate *MongoTemplate) UpdateOne(ctx context.Context, criteria bson.M, update bson.M) int64 {
	updateResult, _ := mongoTemplate.collection.UpdateOne(ctx, criteria, update)
	return updateResult.ModifiedCount
}

func (mongoTemplate *MongoTemplate) UpdateMany(ctx context.Context, criteria bson.M, update bson.M) int64 {
	updateResult, _ := mongoTemplate.collection.UpdateMany(ctx, criteria, update)
	return updateResult.ModifiedCount
}
