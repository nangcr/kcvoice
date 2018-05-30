// kcvoice包为舰队Collection游戏语音实现了一个简单的爬虫。
package kcvoice

import (
	"net/http"
	"fmt"
	"regexp"
	"io/ioutil"
	"errors"
)

// 硬编码数据来源，请勿改动
const (
	MoegirlUrl    = "https://zh.moegirl.org/File:%s%s.mp3"
	MoegirlRegFmt = `<div class="fullMedia"><a href="([^"]+)"`
)

// 数据来源是一个接收格式化数据(舰娘名字和语音编号)以及正则表达式的结构体
type Source struct {
	url    string
	regfmt string
}

// 自定义数据源
func NewSource(url string, regfmt string) (this Source) {
	this.url = url
	this.regfmt = regfmt
	return
}

// 萌娘百科数据源获取
func NewMoegirlSource() (this Source) {
	this.url = MoegirlUrl
	this.regfmt = MoegirlRegFmt
	return
}

// 默认数据源获取
func NewDefaultSource() (this Source) {
	this = NewMoegirlSource()
	return
}

// 根据舰娘名字和语音编号来获取链接
func (a Source) GetUrl(name string, id int) (result string, err error) {
	client := &http.Client{}

	sid := fmt.Sprintf("%d", id)
	if id < 10 {
		sid = fmt.Sprintf("0%d", id)
	}
	url := fmt.Sprintf(a.url, name, sid)

	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	response, err := client.Do(reqest)
	if err != nil {
		return
	}

	reg, err := regexp.Compile(a.regfmt)
	if err != nil {
		return
	}

	str, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	res := reg.FindStringSubmatch(string(str))
	if len(res) == 0 {
		errinfo := fmt.Sprintf("Page Not Found(%s %s)", name, sid)
		err = errors.New(errinfo)
		return
	}

	result = res[1]
	return
}

// 根据舰娘名字获取所有链接,参数limite用以设置失败重试次数
func (a Source) GetUrls(name string, limit int) (urls []string) {
	var wrongAnswer int
	for i := 1; wrongAnswer <= limit; i++ {
		str, err := a.GetUrl(name, i)
		if err != nil {
			fmt.Println(err)
			wrongAnswer++
			continue
		}
		urls = append(urls, str)
	}
	return
}
