package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"sync"
)

var httpreg = regexp.MustCompile(`<(a|link).*href=["'](.+?)["']`)

var pathList []string

func Control(tg string, deep bool) {
	wg := &sync.WaitGroup{}
	fmt.Println("start")
	webUrl, err := url.Parse(tg)
	if err != nil {
		log.Fatal(err)
	}
	if deep {
		goPathFinder(webUrl, wg)
	} else {
		pathFinder(webUrl, wg)
	}
	wg.Wait()
	if deep {
		// OutPut.TxtOut(webUrl.Host, pathList)
	} else {
		for _, v := range pathList {
			fmt.Println(v)
		}
	}
}

func pathFinder(tg *url.URL, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	// 请求网页
	resp, err := http.Get(tg.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 调用正则，取href
	allUrl := httpreg.FindAllSubmatch(body, -1)
	for _, tmpList := range allUrl {
		tmpUrl, err := url.Parse(string(tmpList[2]))
		if err != nil {
			log.Fatal()
		}
		//判断一下href对象是否为目标域名下，防止跑到其他网站,这里用的host，有点小问题，因为如果用的旁站资源那么就会无法导入。
		if tmpUrl.Host == tg.Host && !IndexOf(tmpUrl.String(), pathList) {
			pathList = append(pathList, tmpUrl.String())
			// 递归调用，跑完全网站
			pathFinder(tmpUrl, wg)
		}
	}

}

func goPathFinder(tg *url.URL, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	// 请求网页
	resp, err := http.Get(tg.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	allUrl := httpreg.FindAllSubmatch(body, -1)
	for _, tmpList := range allUrl {
		tmpUrl, err := url.Parse(string(tmpList[2]))
		if err != nil {
			log.Fatal()
		}
		//判断一下href对象是否为目标域名下，防止跑到其他网站,这里用的host，有点小问题，因为如果用的旁站资源那么就会无法导入。
		if tmpUrl.Host == tg.Host && !IndexOf(tmpUrl.String(), pathList) {
			pathList = append(pathList, tmpUrl.String())
			go goPathFinder(tmpUrl, wg)
		}
	}
}

func IndexOf(atom string, array []string) bool {
	for _, value := range array {
		if atom == value {
			return true
		}
	}
	return false
}