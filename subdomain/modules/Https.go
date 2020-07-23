package modules

import (
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"regexp"
)

func CeFind(target string) []string{
	var subdomain []string
	// 数据合并
	subdomain =append(subdomain,crtsh(target)...)
	subdomain = append(subdomain,certspotter(target)...)

	return subdomain
}

// use crt.sh to find
func crtsh(target string) []string{
	req,err := http.Get("https://crt.sh/?output=json&q="+ target)
	if err != nil{
		color.Red("crtsh error!",err)
	}
	defer req.Body.Close()

	// 获取主体并且进行分割，拆分的结果存入数组body中，每个元素都是一条json
	context,err := ioutil.ReadAll(req.Body)

	re := regexp.MustCompile("},(.*?)(\\n*?)(.*?){")
	tmp := re.ReplaceAll(context,[]byte("}#{"))
	tmp = bytes.Trim(tmp,"[")
	tmp = bytes.Trim(tmp,"]")
	body := bytes.Split(tmp,[]byte("#"))

	var subdomain []string
	// 从body中取json进行分析，同时将分析结果内的url加入子域切片内
	for _,cont := range body {
		js,err := simplejson.NewJson(cont)
		if err != nil {
			panic(err)
		}
		domain,err := js.Get("name_value").String()
		if err != nil {
			panic(err)
		}

		subdomain = append(subdomain,domain)
	}
	return subdomain
}


func certspotter(tg string) []string{

	req,err := http.Get("https://api.certspotter.com/v1/issuances?expand=dns_names&include_subdomains="+tg)
	if err != nil {
		color.Red("certspotter error!",err)
	}
	defer req.Body.Close()

	context,err := ioutil.ReadAll(req.Body)
	//分割json数据
	re := regexp.MustCompile("},(.*?)(\\n*?)(.*?){")
	tmp := re.ReplaceAll(context,[]byte("}#{"))
	tmp = bytes.Trim(tmp,"[")
	tmp = bytes.Trim(tmp,"]")
	body := bytes.Split(tmp,[]byte("#"))

	var subdomain []string //用来放子域的切片
	for _,value := range body{
		// 解析json
		js,err := simplejson.NewJson(value)
		if err != nil{
			panic(err)
		}

		//提取dns_names数组,并入subdomain
		subarr,err := js.Get("dns_names").StringArray()
		if err!= nil {
			color.Red("Couldn't use certspotter")
			fmt.Println("")
			return subdomain
		}
		subdomain=append(subdomain,subarr...)
	}

	return subdomain
}