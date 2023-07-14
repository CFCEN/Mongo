package main

import (
	"Mongo/Db"
	"context"
	"fmt"
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
	var userSetting UserSetting
	//mongoTemplate.FindOneById(ctx, "60595428721f2314cbbf6c65", &userSetting)

	query := Db.Where("app", "nezha")

	list := make([]UserSetting, 0)
	mongoTemplate.FindMany(ctx, query.GetCriteria(), list, &userSetting)

	fmt.Println(userSetting)
}
