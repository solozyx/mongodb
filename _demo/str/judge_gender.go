package main

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

// idolSn 第4位,女1男2 赵丽颖0801000003 蔡徐坤0802000002
func judgeGender(idolSn string) int {
	flag := idolSn[3:4]
	flagInt, _ := strconv.Atoi(flag)
	return flagInt % 2
}

func main() {
	gender := judgeGender("0802000002")
	logrus.Debugf("gender=%d", gender)

	gender = judgeGender("0801000003")
	logrus.Debugf("gender=%d", gender)
}
