package service

import (
	"fmt"
	"github.com/labstack/gommon/color"
	"net"
	"sync"
	"time"
)

var ip interface{}
var portList []int

func PortScan(target string) {
	var err error
	ip, err = net.ResolveIPAddr("ip4", target)
	if err != nil {
		color.Red("error!")
		panic(err)
	}
	color.Yellow("Start Port Scan ")

	tcpScan(target)

	portList = unique(portList)

	for _, port := range portList {
		fmt.Println("Open Port:", port)
	}

}

func unique(ataxic []int) []int {
	var unique []int
	for _, value := range ataxic {
		if !indexOf(value, unique) {
			unique = append(unique, value)
		}
	}
	return unique
}

func indexOf(atom int, array []int) bool {
	// Did atom in array?
	for _, value := range array {
		if atom == value {
			return true
		}
	}
	return false
}

func tcpScan(target string) {
	wg := sync.WaitGroup{}
	for i := 1; i <= 65535; i++ {
		go tcpAddPortList(i, &wg)
	}
	wg.Wait()
}

func tcpDetectPort(port int) bool {
	addr := fmt.Sprintf("%v:%d", ip, port)
	_, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return false
	}
	return true
}

func tcpAddPortList(port int, wg *sync.WaitGroup) {
	wg.Add(1)
	if tcpDetectPort(port) {
		portList = append(portList, port)
	}
	wg.Done()
}
