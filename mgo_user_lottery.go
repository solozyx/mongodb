package main

import (
	"gopkg.in/mgo.v2/bson"
)

type MgoUserLottery struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	UserID   int           `bson:"user_id" json:"user_id"`
	CreateAt string        `bson:"create_at" json:"create_at"`
}
