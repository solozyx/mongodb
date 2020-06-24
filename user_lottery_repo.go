package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	_mgoUserLottery = "mgo_user_lottery"
)

func CreateUserLottery(entity MgoUserLottery) error {
	session := mgoSess.Copy()
	defer session.Close()

	entity.ID = bson.NewObjectId()
	entity.CreateAt = time.Now().Format(time.RFC3339)

	collection := session.DB(_mgoDB).C(_mgoUserLottery)
	if err := collection.Insert(&entity); err != nil {
		return err
	}
	return nil
}

func DelByUserId(userId int) error {
	session := mgoSess.Copy()
	defer session.Close()
	collection := session.DB(_mgoDB).C(_mgoUserLottery)
	return collection.Remove(bson.M{"user_id": userId})
}
