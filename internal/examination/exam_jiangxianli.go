package examination

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	storage "gitee.com/erdanli/ipproxypool/internal/dbops"
	getip "gitee.com/erdanli/ipproxypool/internal/getip"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ProxyJiangXianLiTest(proxy_addr string) (Speed int, Status int) {
	//检测代理iP访问地址
	var testUrl string
	//判断传来的代理IP是否是https
	if strings.Contains(proxy_addr, "https") {
		testUrl = "https://icanhazip.com"
	} else {
		testUrl = "http://icanhazip.com"
	}
	// 解析代理地址
	proxy, err := url.Parse(proxy_addr)
	//设置网络传输
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(3),
	}
	// 创建连接客户端
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	begin := time.Now() //判断代理访问时间
	// 使用代理IP访问测试地址
	res, err := httpClient.Get(testUrl)

	if err != nil {
		return
	}
	defer res.Body.Close()
	speed := int(time.Now().Sub(begin).Nanoseconds() / 1000 / 1000) //ms
	//判断是否成功访问，如果成功访问StatusCode应该为200
	if res.StatusCode != http.StatusOK {
		return
	}
	return speed, res.StatusCode
}

// ...
func TestJiangXianLi() ([]string, []string) {
	var availableProxyIP []string
	var unavailableProxyIP []string
	resultIP := getip.GetJianXianLiProxyIP()
	log.Println("Please Wating ...")
	for _, ip := range resultIP {
		var _, status = ProxyJiangXianLiTest(ip)
		if status == 200 {
			availableProxyIP = append(availableProxyIP, ip)
		} else {
			unavailableProxyIP = append(unavailableProxyIP, ip)
		}
	}
	log.Println("可用代理IP:", availableProxyIP)
	return availableProxyIP, unavailableProxyIP
}

// ...
func DeleteUnavailableProxyIP() {
	_, _, resultIP := storage.MangoDB()
	log.Println("Please Wating ...")
	for _, ip := range resultIP {
		var _, status = ProxyJiangXianLiTest(ip)
		if status != 200 {
			connMongoDB(ip)
			log.Println("删除无用proxyip:", ip)
		}
	}
}

func connMongoDB(ip string) {
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
	err = collection.Remove(bson.M{"ip": ip})
	if err != nil {
		fmt.Println("collection.RemoveAll Error:", err)
		return
	} else {
		fmt.Println("删除成功！")
	}
	fmt.Println("Things objects count: ", countNum)
}
