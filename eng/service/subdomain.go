package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/bitly/go-simplejson"
	"github.com/labstack/gommon/color"
)

var subdomain []string

func FindSubdomain(target string) { // use function to find subdomain and organize data
	wg := sync.WaitGroup{}
	wg.Add(2) //如果在子函数里面写 wg.add(1) 这种，会直接跑过去而不是停留。
	go CeFind(target, &wg)
	wg.Wait()
	sub := uniqueString(subdomain)
	for n, _ := range sub {
		fmt.Println(sub[n])
	}
}

func indexOfString(atom string, array []string) bool {
	// Did atom in array?
	for _, value := range array {
		if atom == value {
			return true
		}
	}
	return false
}

func CeFind(target string, wg *sync.WaitGroup) {
	ceWg := sync.WaitGroup{}
	go crtsh(target, &ceWg)
	go certspotter(target, &ceWg)
	ceWg.Wait()
	wg.Done()
}

// use crt.sh to find
func crtsh(target string, wg *sync.WaitGroup) {
	wg.Add(1)
	req, err := http.Get("https://crt.sh/?output=json&q=" + target)
	if err != nil {
		wg.Done()
		return
	}
	defer req.Body.Close()
	// 获取主体并且进行分割，拆分的结果存入数组body中，每个元素都是一条json
	context, err := ioutil.ReadAll(req.Body)
	if err != nil {
		wg.Done()
		return
	}
	re := regexp.MustCompile("},(.*?)(\\n*?)(.*?){")
	tmp := re.ReplaceAll(context, []byte("}#{"))
	tmp = bytes.Trim(tmp, "[")
	tmp = bytes.Trim(tmp, "]")
	body := bytes.Split(tmp, []byte("#"))
	// 从body中取json进行分析，同时将分析结果内的url加入子域切片内
	for _, cont := range body {
		js, err := simplejson.NewJson(cont)
		if err != nil {
			wg.Done()
			return
		}
		domain, err := js.Get("name_value").String()
		if err != nil {
			wg.Done()
			return
		}
		subdomain = append(subdomain, domain)
	}
	wg.Done()
}

func certspotter(tg string, wg *sync.WaitGroup) {
	wg.Add(1)
	req, err := http.Get("https://api.certspotter.com/v1/issuances?expand=dns_names&include_subdomains=true&domain=" + tg)
	if err != nil {
		wg.Done()
		return
	}
	defer req.Body.Close()
	context, err := ioutil.ReadAll(req.Body)
	if err != nil {
		wg.Done()
		return
	}
	//分割json数据
	re := regexp.MustCompile("},(.*?)(\\n*?)(.*?){")
	tmp := re.ReplaceAll(context, []byte("}#{"))
	tmp = bytes.Trim(tmp, "[")
	tmp = bytes.Trim(tmp, "]")
	body := bytes.Split(tmp, []byte("#"))
	var subdomain []string //用来放子域的切片
	for _, value := range body {
		// 解析json
		js, err := simplejson.NewJson(value)
		if err != nil {
			wg.Done()
			return
		}
		//提取dns_names数组,并入subdomain
		subarr, err := js.Get("dns_names").StringArray()
		if err != nil {
			color.Red("Couldn't use certspotter")
			wg.Done()
			return
		}
		subdomain = append(subdomain, subarr...)
	}
	wg.Done()
}

func uniqueString(ataxic []string) []string {
	//I want to use simHash
	//But now,I don't know how to implement the code
	//So only index of
	var unique []string
	for _, value := range ataxic {
		value = strings.TrimSpace(value)
		if !indexOfString(value, unique) {
			unique = append(unique, value)
		}
	}

	return unique
}
