package getip

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/go-clog/clog"
)

// ...
func GetIP3306PaginationURL() []string {
	urlPath := "http://www.ip3366.net/free/?stype=1&page=1"
	resp, err := http.Get(urlPath)
	if err != nil {
		clog.Info("[IP3306]] parse pollUrl error")
		clog.Warn(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// 解决中文乱码
	html := mahonia.NewDecoder("gbk").ConvertString(string(body))

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		errors.New("err is not nil!")
	}
	lastNum := 0
	dom.Find("#listnav > ul > a:nth-child(10)").Each(func(i int, selection *goquery.Selection) {
		if href, ok := selection.Attr("href"); ok {
			lastNum, _ = strconv.Atoi(href[len(href)-1:])
		}
	})

	var result []string
	for i := 1; i <= lastNum; i++ {
		pageURL := "http://www.ip3366.net/free/?stype=1&page=" + strconv.Itoa(i)
		result = append(result, pageURL)
	}
	return result
}

// ...
func GetIP3306ProxyIP() []string {
	var result []string
	pageURL := GetIP3306PaginationURL()
	for _, urlPath := range pageURL {
		resp, err := http.Get(urlPath)
		if err != nil {
			clog.Info("[IP3306]] parse pollUrl error")
			clog.Warn(err.Error())
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		// 解决中文乱码
		html := mahonia.NewDecoder("gbk").ConvertString(string(body))

		dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			log.Println("[IP3306]] parse pollUrl error")
		}
		// #list > table > thead > tr > th:nth-child(4)
		// #list > table > tbody > tr:nth-child(1) > td:nth-child(1)
		dom.Find("#list > table > tbody > tr > td:nth-child(1)").Each(func(i int, selection *goquery.Selection) {
			ip := selection.Text()
			port := selection.Next().Text()
			http := "http://" + ip + ":" + port
			result = append(result, http)
		})
	}
	return result
}
