package main

import (
	"fmt"
	dis "discovery"
	"log"
	"time"
)

func main() {

	serviceName := "s-test"
	serviceInfo := dis.ServiceInfo{IP:"192.168.1.26"}

	s, err := dis.NewService(serviceName, serviceInfo,[]string {
		"http://192.168.1.17:2379",
 		"http://192.168.1.17:2479",
 		"http://192.168.1.17:2579",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name:%s, ip:%s\n", s.Name, s.Info.IP)


	go func() {
		time.Sleep(time.Second*20)
		s.Stop()
	}()
	
	s.Start()
}
