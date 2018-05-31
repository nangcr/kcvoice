// kcvoice包为舰队Collection游戏实现了一个简单的爬虫。
package kcvoice

import (
	"net/http"
	"fmt"
	"regexp"
	"io/ioutil"
)

// 硬编码数据来源，请勿改动
const (
	MoegirlNameUrl     = `https://zh.moegirl.org/%E8%88%B0%E9%98%9FCollection/%E5%9B%BE%E9%89%B4/%E8%88%B0%E5%A8%98`
	MoegirlNameRegFmt  = `>No.[0-9][0-9][0-9] ([^"]+)</a>`
	MoegirlVoiceUrl    = `https://zh.moegirl.org/舰队Collection:%s`
	MoegirlVoiceRegFmt = `data-filesrc="(https://img.moegirl.org/common/[^"]+)"`
)

// 数据来源是一个接收格式化数据(舰娘名字和语音编号)以及正则表达式的结构体
type Source struct {
	nameUrl     string
	voiceUrl    string
	nameRegfmt  string
	voiceRegfmt string
}

// 自定义数据源
func NewSource(nameUrl string, voiceUrl string, nameRegfmt string, voiceRegfmt string) *Source {
	return &Source{
		nameUrl:     nameUrl,
		voiceUrl:    voiceUrl,
		nameRegfmt:  nameRegfmt,
		voiceRegfmt: voiceRegfmt,
	}
}

// 萌娘百科数据源获取
func NewMoegirlSource() *Source {
	return &Source{
		nameUrl:     MoegirlNameUrl,
		voiceUrl:    MoegirlVoiceUrl,
		nameRegfmt:  MoegirlNameRegFmt,
		voiceRegfmt: MoegirlVoiceRegFmt,
	}
}

// 默认数据源获取
func NewDefaultSource() *Source {
	return NewMoegirlSource()
}

// 获取所有舰娘名字
func (s Source) GetNames() (list []string, err error) {
	client := &http.Client{}
	url := s.nameUrl
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	response, err := client.Do(request)
	if err != nil {
		return
	}

	reg, err := regexp.Compile(s.nameRegfmt)
	if err != nil {
		return
	}

	str, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	res := reg.FindAllStringSubmatch(string(str), -1)
	for _, v := range res {
		list = append(list, v[1])
	}
	return
}

// 根据舰娘名字来获取链接列表
func (s Source) GetUrls(name string) (list []string, err error) {
	client := &http.Client{}
	url := fmt.Sprintf(s.voiceUrl, name)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	response, err := client.Do(request)
	if err != nil {
		return
	}

	reg, err := regexp.Compile(s.voiceRegfmt)
	if err != nil {
		return
	}

	str, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	res := reg.FindAllStringSubmatch(string(str), -1)
	for _, v := range res {
		list = append(list, v[1])
	}
	return
}
