package main

import (
	"github.com/sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"
)

const (
	_mgoUserPrize = "mgo_user_prize"
)

func PrizeList(pageNo, pageSize int) ([]MgoUserPrize, bool, error) {
	session := mgoSess.Copy()
	defer session.Close()

	skip := (pageNo - 1) * pageSize
	var hasMore bool

	collection := session.DB(_mgoDB).C(_mgoUserPrize)
	list := make([]MgoUserPrize, 0)

	total, err := collection.Find(bson.M{"prize_id": bson.M{"$lt": 6}}).Select(bson.M{"prize_id": 1, "_id": 0}).Count()
	if err != nil {
		logrus.Errorf("PrizeList Count error=%v", err.Error())
		return nil, hasMore, err
	}

	// 最新的中奖信息
	// 再来一次 不显示在奖品列表
	err = collection.Find(bson.M{"prize_id": bson.M{"$lt": 6}}).
		Select(bson.M{"user_id": 1, "prize_name": 1, "prize_level": 1, "create_at": 1, "_id": 0}).
		Sort("-create_at").
		Skip(skip).
		Limit(pageSize).
		All(&list)
	if err != nil {
		logrus.Errorf("PrizeList error=%v", err.Error())
		return nil, hasMore, err
	}

	// 判断是否有下一页
	count := skip + len(list)
	if total > count {
		hasMore = true
	}

	return list, hasMore, nil
}
