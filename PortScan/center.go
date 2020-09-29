package PortScan

import (
	"Moon_Trace/Global"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"net"
)

var ip interface{}
var portList []int
func PortScan(target,scantype string) {
	var err error
	ip,err = net.ResolveIPAddr("ip4",target)
	if err != nil {
		color.Red("error!")
		panic(err)
	}
	color.Yellow("Start Port Scan ")

	if scantype == "tcp" {
		tcpScan(target)
	}else if scantype == "udp"{
		udpScan(target)
	}else {
		fmt.Println("Sorry,this version don't have this scan_type")
	}

	portList = unique(portList)

	for _,port := range portList{
		fmt.Println("Open Port:",port)
	}


}

func unique(ataxic []int) []int{
	//I want to use simHashs
	//But now,I don't know how to implement the code
	//So only index of
	var unique []int
	for _,value := range ataxic {
		if !indexOf(value,unique){
			unique = append(unique,value)
		}
	}

	return unique
}

func indexOf(atom int,array []int) bool{
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
		cli.BoolFlag{
			Name: "port",
			Usage: "PortScan",
		},
		cli.BoolFlag{
			Name:        "udp",
			Usage:       "UdpScan",
		},
		cli.BoolFlag{
			Name:        "tcp",
			Usage:       "TcpScan",
		},
	)
}

