package PortScan

import (
	"fmt"
	"github.com/fatih/color"
	"net"
	"sync"
	"time"
)

func tcpScan(target string) {
	color.Yellow("base tcp")
	wg := sync.WaitGroup{}
	for i:=1;i<=65535;i++ {
		go tcpAddPortList(i,&wg)
	}
	wg.Wait()
}

func tcpDetectPort(port int) bool{
	addr := fmt.Sprintf("%v:%d",ip,port)
	_,err := net.DialTimeout("tcp",addr,5*time.Second)
	if err != nil {
		return false
	}
	return true
}

func tcpAddPortList(port int,wg *sync.WaitGroup){
	wg.Add(1)
	if tcpDetectPort(port){
		portList = append(portList,port)
	}
	wg.Done()
}