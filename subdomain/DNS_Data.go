package subdomain

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

func DnsData(target string,wg *sync.WaitGroup)  {
	ceBaidu(target)
	//subdomain = append(subdomain,dnsBufferOver(target)...)
	wg.Done()
}

// cebaidu has some problems
func ceBaidu(target string) {
	// 虽然是ce.baidu 但是其实这是一个百度的安全观测站。
	req,err := http.Get("https://ce.baidu.com/index/getRelatedSites?site_address="+ target)
	if err != nil {
		color.Red("Baidu DNS error!",err)
		return
	}
	defer req.Body.Close()
	context,err := ioutil.ReadAll(req.Body)

	//嵌套json，不知为何无法直接取值。
	js,err := simplejson.NewJson(context)
	if err != nil{
			color.Red("Can't use CeBaidu,Json wrong!")
			return
		}
	//要想办法将interface转换为[]byte
	jsarr,err := js.Get("data").Array()
	if err!= nil {
		color.Red("Can't use CeBaidu,may be frequent requests")
		return
	}

	for _,tmp := range jsarr {
		if m,ok := tmp.(map[string]interface{});ok {
			for _,sub := range m{
				switch sub.(type) {
				case string:
					subdomain = append(subdomain,sub.(string))
					fmt.Println(sub)
				}
			}
		}
	}

}

//TODO:
// 使用dns.bufferover.run
// cloudfare拦截，无法使用。
func dnsBufferOver(target string) []string {
	req,err := http.NewRequest("GET","http://dns.bufferover.run/dns?q="+target,nil)
	if err != nil {
		color.Red("Can't use dns.bufferover.run")
		return subdomain
	}
	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")
	clt := http.Client{}
	resp,err := clt.Do(req)
	if err != nil {
		color.Red("Can't use dns.bufferover.run")
		return subdomain
	}
	defer resp.Body.Close()
	context,err := ioutil.ReadAll(resp.Body)


	re := regexp.MustCompile("\"(.+?),")
	body := re.ReplaceAll(context,[]byte("\""))


	js,err := simplejson.NewJson(body)
	if err != nil {
		color.Red("Error!",err)
	}
	subArr,err := js.Get("FDNS_A").StringArray()
	if err != nil {
		color.Red("JS Error!")
		panic(err)
	}
	tmp,err := js.Get("RDNS").StringArray()
	if err != nil {
		color.Red("JS Error!")
		panic(err)
	}

	subArr = append(subArr,tmp...)
	for _,value := range subArr {
		subdomain = append(subdomain,value)
	}
	return subdomain

}