package storage

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

// ...
type ProxyIP struct {
	IP string
}

// ...avaliableProxyIP []string
func MangoDB() (*mgo.Database, *mgo.Session, []string) {

	mgoURL := "mongodb://localhost:27017/test"

	session, err := mgo.Dial(mgoURL)

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
	// for _, ip := range avaliableProxyIP {
	// 	result := ProxyIP{
	// 		IP: ip,
	// 	}
	// 	//单个对象数据插入
	// 	err = collection.Insert(result)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("------插入数据成功------------")

	//*****查询单条数据*******
	var proxyIP []ProxyIP
	collection.Find(nil).All(&proxyIP)
	log.Println("proxyIP:", proxyIP)
	var result []string
	for _, v := range proxyIP {
		result = append(result, v.IP)
	}
	fmt.Println("------查询数据成功------------")
	return db, session, result
}
