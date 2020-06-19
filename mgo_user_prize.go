package main

import (
	"gopkg.in/mgo.v2/bson"
)

type MgoUserPrize struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	UserID         int           `bson:"user_id" json:"user_id"`
	PrizeID        int           `bson:"prize_id" json:"prize_id"`
	ReceiveID      string        `bson:"receive_id" json:"receive_id"`
	Status         ReceiveStatus `bson:"status" json:"status"`
	CreateAt       string        `bson:"create_at" json:"create_at"`
	ReceiveAt      string        `bson:"receive_at" json:"receive_at"`
	ExpireAt       string        `bson:"expire_at" json:"expire_at"`
	PrizeName      string        `bson:"prize_name" json:"prize_name"`
	PrizeLevel     int           `bson:"prize_level" json:"prize_level"`
	ImageUrl       string        `bson:"image_url" json:"image_url"`
	AgainUseStatus AgainStatus   `bson:"again_use_status" json:"again_use_status"`
}

// 奖品领取状态
type ReceiveStatus uint

const (
	// 未领取
	ReceiveStatusWaiting ReceiveStatus = iota
	// 已领取
	ReceiveStatusSucceeded
)

// 再来一次 获取的抽奖次数 使用状态
type AgainStatus uint

const (
	AgainStatusNotUse AgainStatus = iota
	AgainStatusUsed
)

type UserPrize struct {
	UserName   string
	PrizeName  string
	CreateAt   string
	PrizeLevel int
}
