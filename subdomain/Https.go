package subdomain

import (
	"bytes"
	"github.com/bitly/go-simplejson"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)



func CeFind(target string,wg *sync.WaitGroup) {
	ceWg := sync.WaitGroup{}

	go crtsh(target,&ceWg)
	go certspotter(target,&ceWg)

	ceWg.Wait()
	wg.Done()
}

// use crt.sh to find
func crtsh(target string,wg *sync.WaitGroup) {
	wg.Add(1)

	req,err := http.Get("https://crt.sh/?output=json&q="+ target)
	if err != nil{
		color.Red("crtsh error!",err)
		wg.Done()
		return
	}
	defer req.Body.Close()

	// 获取主体并且进行分割，拆分的结果存入数组body中，每个元素都是一条json
	context,err := ioutil.ReadAll(req.Body)
	if err != nil {
		color.Red("crtsh error!",err)
		wg.Done()
		return
	}

	re := regexp.MustCompile("},(.*?)(\\n*?)(.*?){")
	tmp := re.ReplaceAll(context,[]byte("}#{"))
	tmp = bytes.Trim(tmp,"[")
	tmp = bytes.Trim(tmp,"]")
	body := bytes.Split(tmp,[]byte("#"))

	// 从body中取json进行分析，同时将分析结果内的url加入子域切片内
	for _,cont := range body {
		js,err := simplejson.NewJson(cont)
		if err != nil {
			color.Red("crtsh error!",err)
			wg.Done()
			return
		}
		domain,err := js.Get("name_value").String()
		if err != nil {
			color.Red("crtsh can't use!")
			wg.Done()
			return
		}

		subdomain = append(subdomain,domain)
	}
	wg.Done()
}


func certspotter(tg string ,wg *sync.WaitGroup) {
	wg.Add(1)

	req,err := http.Get("https://api.certspotter.com/v1/issuances?expand=dns_names&include_subdomains="+tg)
	if err != nil {
		color.Red("certspotter error!",err)
		wg.Done()
		return
	}
	defer req.Body.Close()

	context,err := ioutil.ReadAll(req.Body)
	if err != nil {
		wg.Done()
		return
	}
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
			wg.Done()
			return
		}

		//提取dns_names数组,并入subdomain
		subarr,err := js.Get("dns_names").StringArray()
		if err!= nil {
			color.Red("Couldn't use certspotter")
			wg.Done()
			return
		}
		subdomain=append(subdomain,subarr...)
	}
	wg.Done()
}