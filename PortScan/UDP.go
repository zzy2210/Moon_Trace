package PortScan

import (
	"fmt"
	"github.com/fatih/color"
	"net"
	"sync"
	"time"
)

func udpScan(target string) {
	color.Yellow("base ucp")
	wg := sync.WaitGroup{}
	for i:=1;i<=65535;i++ {
		go udpAddPortList(i,&wg)
	}
	wg.Wait()
}

// TODO:udp响应需要修改，因为udp的无连接性，这里的判定只要发送就可以所以一直成立
func udpDetectPort(port int)bool {
	addr := fmt.Sprintf("%v:%d",ip,port)
	_,err := net.DialTimeout("udp",addr,5*time.Second)
	if err != nil {
		return false
	}
	return true
}

func udpAddPortList(port int,wg *sync.WaitGroup){
	wg.Add(1)
	if udpDetectPort(port){
		portList = append(portList,port)
	}
	wg.Done()
}