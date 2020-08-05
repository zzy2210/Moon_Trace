package subdomain

import (
	"Moon_Trace/Global"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"strings"
	"sync"
)


var subdomain []string

func FindSubdomain(target string){ // use function to find subdomain and organize data
	wg := sync.WaitGroup{}
	wg.Add(2) //如果在子函数里面写 wg.add(1) 这种，会直接跑过去而不是停留。
	color.Yellow("Start subdomain find")

	go DnsData(target,&wg)
	go CeFind(target,&wg)

	wg.Wait()
	sub := unique(subdomain)

	for n,_ := range sub{
		fmt.Println(sub[n])
	}
}

	func unique(ataxic []string) []string{
		//I want to use simHash
		//But now,I don't know how to implement the code
		//So only index of
		var unique []string
		for _,value := range ataxic {
			value = strings.TrimSpace(value)
			if !indexOf(value,unique){
				unique = append(unique,value)
			}
		}

		return unique
	}

	func indexOf(atom string,array []string) bool{
		// Did atom in array?
		for _,value := range array {
			if atom == value {
				return true
			}
		}
		return false
	}


	func init() {
		Global.Moon.Flags = append(Global.Moon.Flags,
			cli.BoolFlag{ // 参数判定，启用子域查询
				Name:  "sub",
				Usage: "to find subdimain",
			},
			)

	}