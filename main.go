// go-scanport project main.go
package main

import (
	"fmt"
	"net"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func TestIpPort(address string) {
	defer wg.Done()
	conn, err := net.DialTimeout("tcp", address, time.Nanosecond*10)
	if err != nil {
		//fmt.Print(err)
		return
	}

	fmt.Printf("%s is open\n", address)
	conn.Close()

}

func main() {
	args := os.Args
	StartIP := ""
	EndIP := ""
	StartPort := ""
	EndPort := ""
	if len(args) <= 1 {
		fmt.Printf("No args,exit(0)\n")
		fmt.Printf("usage:a.exe -StartIP [ip] -EndIP [ip] -StartPort [port] -EndPort [port]\n")
		return
	} else if strings.Compare(args[1], "/?") == 0 {
		fmt.Printf("usage:a.exe -StartIP [ip] -EndIP [ip] -StartPort [port] -EndPort [port]\n")
		return
	} else if strings.Compare(args[1], "--help") == 0 {
		fmt.Printf("usage:a.exe -StartIP [ip] -EndIP [ip] -StartPort [port] -EndPort [port]\n")
		return
	}
	for i := 0; i < len(args); i++ {
		if strings.Compare(args[i], "-StartIP") == 0 {
			StartIP = args[i+1]
		} else if strings.Compare(args[i], "-EndIP") == 0 {
			EndIP = args[i+1]
		} else if strings.Compare(args[i], "-StartPort") == 0 {
			StartPort = args[i+1]
		} else if strings.Compare(args[i], "-EndPort") == 0 {
			EndPort = args[i+1]
		}
	}
	s := strings.Split(StartIP, ".")
	e := strings.Split(EndIP, ".")
	var sIP, eIP [4]int
	for i := 0; i < 4; i++ {
		sIP[i], _ = strconv.Atoi(s[i])
		eIP[i], _ = strconv.Atoi(e[i])
	}
	sP, _ := strconv.Atoi(StartPort)
	eP, _ := strconv.Atoi(EndPort)
	fmt.Printf("begin to scan IP:%s---%s,port:%d---%d\n", StartIP, EndIP, sP, eP)
	debug.SetMaxThreads(65535)
	//begin to scan
	for addr1 := sIP[0]; addr1 <= eIP[0]; addr1++ {
		for addr2 := sIP[1]; addr2 <= eIP[1]; addr2++ {
			for addr3 := sIP[2]; addr3 <= eIP[2]; addr3++ {
				for addr4 := sIP[3]; addr4 < eIP[3]; addr4++ {
					for port := 0; port <= 65535; port++ {
						wg.Add(1)
						go TestIpPort(net.JoinHostPort(fmt.Sprintf("%d.%d.%d.%d", addr1, addr2, addr3, addr4), strconv.Itoa(port)))
					}
					fmt.Printf(fmt.Sprintf("%d.%d.%d.%d is complete\n", addr1, addr2, addr3, addr4))
				}
			}

		}
	}
	wg.Wait()
	fmt.Printf("done!\n")
}
