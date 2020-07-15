package modules

import (
	"bytes"
	"github.com/bitly/go-simplejson"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
)

func Find(tg string) []string{
	// 数据合并
	subdomain :=crtsh(tg)
//	subdomain = append(subdomain,certspotter(tg)...)
	return subdomain
}

// use crt.sh to find
func crtsh(tg string) []string{
	req,err := http.Get("https://crt.sh/?output=json&q="+tg)
	if err != nil{
		color.Red("crtsh error!",err)
	}
	defer req.Body.Close()

	// 获取主体并且进行分割，拆分的结果存入数组body中，每个元素都是一条json
	context,err := ioutil.ReadAll(req.Body)

	tmp := bytes.ReplaceAll(context,[]byte("["),[]byte(""))
	tmp = bytes.ReplaceAll(tmp,[]byte("]"),[]byte(""))
	tmp = bytes.ReplaceAll(tmp,[]byte("\\n"),[]byte(""))
	tmp = bytes.ReplaceAll(tmp,[]byte("},{"),[]byte("}\\n{"))

	body := bytes.Split(tmp,[]byte("\\n"))

	var subdomain []string //用来放子域的切片
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

// Not yet completed
/*
func certspotter(tg string) []string{

	req,err := http.Get("https://api.certspotter.com/v1/issuances?expand=dns_names&include_subdomains=true&domain="+tg)
	if err != nil {
		color.Red("certspotter error!",err)
	}
	defer req.Body.Close()


}*/