package main

import (
	"context"
	"github.com/CFCEN/Mongo/Db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserSetting struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	App       string             `bson:"app"`
	Uid       primitive.ObjectID `bson:"uid"`
	CreatedAt time.Time          `bson:"createdAt"`
	Oid       primitive.ObjectID `bson:"oid"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func main() {

	var ctx = context.TODO()
	mongoTemplate := Db.Init("mongodb://root:admin@10.5.17.107:27017/?authMechanism=SCRAM-SHA-1&directConnection=true", "nezha_core", "users.settings")
	/*userSetting := UserSetting{
		App:       "nezha_core",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}*/
	//mongoTemplate.InsertOne(ctx, userSetting)
	mongoTemplate.DeleteOne(ctx, Db.Where("app", "nezha_core").GetCriteria())
}
