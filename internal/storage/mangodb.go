package storage

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// ...
type ProxyIP struct {
	IP string
}

// ...
func MangoDB(avaliableProxyIP []string) (*mgo.Database, *mgo.Session) {

	mgoURL := "mongodb://localhost:27017/test" // 可以代替下面两步

	session, err := mgo.Dial(mgoURL)

	// session.SetMode(mgo.Monotonic, true)

	db := session.DB("test") //数据库名称

	if err != nil {
		fmt.Println("------连接数据库失败------------")
		panic(err)
	}
	fmt.Println("------ConnectionDb-----test-------")

	collection := db.C("AvaliableProxyIP")

	countNum, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("Things objects count: ", countNum)

	//*******插入元素*******
	for _, ip := range avaliableProxyIP {
		result := ProxyIP{
			IP: ip,
		}
		//单个对象数据插入
		err = collection.Insert(result)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("------插入数据成功------------")

	//*****查询单条数据*******
	return db, session
}
