package main

import exam "gitee.com/erdanli/ipproxypool/internal/examination"

func main() {
	// avaliableProxyIP, _ := exam.TestJiangXianLi()
	// _, _, result := storage.MangoDB()
	exam.DeleteUnavailableProxyIP()
}
