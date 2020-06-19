package main

import (
	"fmt"
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

func DistinctUserLottery() (int, error) {
	session := mgoSess.Copy()
	defer session.Close()

	collection := session.DB(_mgoDB).C(_mgoUserLottery)
	var result []int
	err := collection.Find(bson.M{"user_id": bson.M{"$gt": 0}}).Distinct("user_id", &result)
	if err != nil {
		return 0, err
	}
	fmt.Printf("result=%v\n", result)
	return len(result), nil
}

// 计算用户参与抽奖次数
func CountUserJoinTimesByUserId(userId int) (int, error) {
	session := mgoSess.Copy()
	defer session.Close()

	collection := session.DB(_mgoDB).C(_mgoUserLottery)
	// count,err := collection.Find(bson.M{"user_id": bson.M{"$eq": userId}}).Select(bson.M{"user_id":1}).Count()
	count, err := collection.Find(bson.M{"user_id": userId}).Select(bson.M{"user_id": 1}).Count()
	if err != nil {
		return 0, err
	}
	fmt.Printf("CountUserJoinTimesByUserId count=%d", count)
	return count, nil
}
