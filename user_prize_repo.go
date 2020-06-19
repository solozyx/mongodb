package main

import (
	"github.com/sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"
)

const (
	_mgoUserPrize = "mgo_user_prize"
)

func PrizeList(skip, limit int) ([]MgoUserPrize, error) {
	session := mgoSess.Copy()
	defer session.Close()

	collection := session.DB(_mgoDB).C(_mgoUserPrize)
	list := make([]MgoUserPrize, 0)

	// 最新的中奖信息
	// 再来一次 不显示在奖品列表
	err := collection.Find(bson.M{"prize_id": bson.M{"$lt": 7}}).
		Select(bson.M{"user_id": 1, "prize_name": 1, "prize_level": 1, "create_at": 1, "_id": 0}).
		Sort("-create_at").
		Skip(skip).
		Limit(limit).
		All(&list)
	if err != nil {
		logrus.Errorf("PrizeList error=%v", err.Error())
		return nil, err
	}

	return list, nil
}
