package main

import (
	exam "gitee.com/erdanli/ipproxypool/internal/examination"
	"gitee.com/erdanli/ipproxypool/internal/storage"
)

func main() {
	avaliableProxyIP, _ := exam.TestJiangXianLi()
	storage.MangoDB(avaliableProxyIP)

}
