package kcvoice

import (
	"sync"
	"fmt"
	"os"
	"os/exec"
)

// 下载传入舰娘的所有语音
func (s Source) Download(names ...string) (failure []string) {
	var wg sync.WaitGroup
	for _, v := range names {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			err := os.Mkdir("./"+v, 0755)
			if err != nil {
				panic(err)
			}
			fmt.Println(v + " 文件夹已创建")
			file, err := os.Create("./" + v + "/" + v)
			if err != nil {
				panic(err)
			}
			fmt.Println(v + " 正在写入数据")
			list, err := s.GetUrls(v)
			if err != nil {
				panic(err)
			}
			if len(list) == 0 {
				failure = append(failure, v)
			}
			for _, v := range list {
				fmt.Fprintln(file, v)
			}
			fmt.Println(v + " 写入完成")
			file.Close()

			fmt.Println(v + " 开始下载")
			cmd := exec.Command("wget", "-P", "./"+v, "-c", "-i", "./"+v+"/"+v)
			err = cmd.Run()
			if err != nil {
				panic(err)
			}
			fmt.Println(v + " 下载完成")
		}(v)
	}
	wg.Wait()
	return
}

// 实验性质的下载函数,请在阅读源码后使用
func (s Source) DownloadAll() (failure []string) {
	var wg sync.WaitGroup
	names, _ := s.GetNames()
	fmt.Println("名字获取完成")
	err := os.Mkdir("./kcvoice/", 0755)
	if err != nil {
		panic(err)
	}
	for _, v := range names {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			err := os.Mkdir("./kcvoice/"+v, 0755)
			if err != nil {
				panic(err)
			}
			fmt.Println(v + " 文件夹已创建")
			file, err := os.Create("./kcvoice/" + v + "/" + v)
			if err != nil {
				panic(err)
			}
			fmt.Println(v + " 正在写入数据")
			list, err := s.GetUrls(v)
			if err != nil {
				panic(err)
			}
			if len(list) == 0 {
				failure = append(failure, v)
			}
			for _, v := range list {
				fmt.Fprintln(file, v)
			}
			fmt.Println(v + " 写入完成")
			file.Close()

			fmt.Println(v + " 开始下载")
			cmd := exec.Command("wget", "-P", "./"+v, "-c", "-i", "./"+v+"/"+v)
			err = cmd.Run()
			if err != nil {
				panic(err)
			}
			fmt.Println(v + " 下载完成")
		}(v)
	}
	wg.Wait()
	return
}
