package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"math/rand"
	"sync"
	"time"
)

var (
	mgoSess *mgo.Session
	rwLock  sync.RWMutex
)

const (
	// _mgoHost = "192.168.174.149:27017"
	_mgoHost = "127.0.0.1:27017"
	_mgoDB   = "MzDB"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	getMgoSession()
}

func getMgoSession() *mgo.Session {
	if mgoSess == nil {
		rwLock.Lock()
		defer rwLock.Unlock()
		if mgoSess == nil {
			mgoSess = connectMongo()
		}
	}
	return mgoSess
}

func connectMongo() *mgo.Session {
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{_mgoHost},
		Database: _mgoDB,
		Username: "",
		Password: "",
	})
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	list, _ := PrizeList((2-1)*5, 5)
	for _, prize := range list {
		logrus.Debugf("%+v", prize)
	}
}

func randomUserId(idRange int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(idRange) + 1
}
