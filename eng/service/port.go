package service

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var ip interface{}

type PortResponse struct {
	PortList []int
}

func PortScan(target string) *PortResponse {
	var err error
	ip, err = net.ResolveIPAddr("ip4", target)
	if err != nil {
		panic(err)
	}
	portResponse := &PortResponse{}
	portResponse.tcpScan(target)
	fmt.Println("ports:", portResponse.PortList)
	portResponse.PortList = unique(portResponse.PortList)
	fmt.Println("after:", portResponse.PortList)
	return portResponse
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

func (p *PortResponse) tcpScan(target string) {
	wg := sync.WaitGroup{}
	for i := 1; i <= 65535; i++ {
		go p.tcpAddPortList(i, &wg)
	}
	wg.Wait()
}

func tcpDetectPort(port int) bool {
	addr := fmt.Sprintf("%v:%d", ip, port)
	_, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return false
	}
	fmt.Println("good port:", addr)
	return true
}

func (p *PortResponse) tcpAddPortList(port int, wg *sync.WaitGroup) {
	wg.Add(1)
	if tcpDetectPort(port) {
		p.PortList = append(p.PortList, port)
	}
	wg.Done()
}
