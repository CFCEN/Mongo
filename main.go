package main

import (
	"github.com/CFCEN/Mongo/Db"
	"context"
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
	mongoTemplate := Db.Init("", "nezha_core", "users.settings")
	query := Db.Where("app", "nezha_core")
	update := Db.Update().Set("app", "nezha")
	mongoTemplate.UpdateMany(ctx, query.GetCriteria(), update.GetUpdate())

}
