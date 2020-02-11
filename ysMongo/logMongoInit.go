package ysMongo

import (
	"fmt"

	"labix.org/v2/mgo"
)

var MongoChain = make(chan *mgo.Session)

func init() {
	for i := 0; i < ChannelNum; i++ {
		GetConnection()
	}
}

// GetConnection 获取一个mongo连接并放入mongo的channel中
func GetConnection() {
	session, err := mgo.Dial(ConnectURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(session)
	MongoChain <- session
	fmt.Println("放入一个mongo连接成功")
}
