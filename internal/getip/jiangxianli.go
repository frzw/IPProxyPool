package getip

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

// 获取分页URL
func GetJianXianLiPaginationURL() []string {
	setPageLastNum := 8
	var result []string
	for i := 1; i < setPageLastNum; i++ {
		url := "https://ip.jiangxianli.com/?page=" + strconv.Itoa(i) + "&protocol=http"
		result = append(result, url)
	}
	return result
}

// 获取页面列表的代理IP
func GetJianXianLiProxyIP() []string {
	var result []string
	resultURL := GetJianXianLiPaginationURL()
	for _, url := range resultURL {
		resp, err := http.Get(url)
		if err != nil {
			errors.New("err is not nil")
		}
		if resp.StatusCode != 200 {
			errors.New("StatusCode != 200")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		// 解决中文乱码
		html := mahonia.NewDecoder("utf8").ConvertString(string(body))
		// goquery
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			errors.New("err is not nil!")
		}
		dom.Find("div.layui-form > table > tbody > tr > td:nth-child(1)").Each(func(i int, selection *goquery.Selection) {
			ip := selection.Text()
			port := selection.Next().Text()
			http := "http://" + ip + ":" + port
			result = append(result, http)
		})
	}
	return result
}
