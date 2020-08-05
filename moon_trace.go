package main

import (
	"Moon_Trace/PortScan"
	"Moon_Trace/subdomain"
	"Moon_Trace/Global"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main(){
	logo := `
 __  __                     _____                   
|  \/  | ___   ___  _ __   |_   _| __ __ _  ___ ___ 
| |\/| |/ _ \ / _ \| '_ \    | || '__/ _  |/ __/ _ \
| |  | | (_) | (_) | | | |   | || | | (_| | (_|  __/
|_|  |_|\___/ \___/|_| |_|___|_||_|  \__,_|\___\___|	
                        |_____|
                                                     ————————author:y1nhui
`
color.Cyan(logo)

var target string

Global.Moon.Flags = append(Global.Moon.Flags,
	cli.StringFlag{ // 目标url的flag，参数绑定到target
		Name:        "u",
		Usage:       "target url",
		Value:       "nil",
		Destination: &target,
	})

Global.Moon.Action = func(c *cli.Context) {

	if c.Bool("sub"){  //调用子域查询
		subdomain.FindSubdomain(target)
	}
	if c.Bool("port"){
		PortScan.PortScan(target)
	}
}


err := Global.Moon.Run(os.Args)
if err!= nil {
	log.Fatal(err)
	}
}

