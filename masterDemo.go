package main

import (
	"log"
	"time"
	"fmt"
	dis "discovery"
)

func main() {

	m, err := dis.NewMaster([]string{
		"http://192.168.1.17:2379",
		"http://192.168.1.17:2479",
		"http://192.168.1.17:2579",
	}, "services/")

	if err != nil {
		log.Fatal(err)
	}

	for {
		for k, v := range  m.Nodes {
			fmt.Printf("node:%s, ip=%s\n", k, v.Info.IP)
		}
		fmt.Printf("nodes num = %d\n",len(m.Nodes))
		time.Sleep(time.Second * 5)
	}
}


