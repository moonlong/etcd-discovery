package discovery

import (
	"github.com/coreos/etcd/clientv3"
	"context"
	"log"
	"time"
	"encoding/json"
	"fmt"
)

type Master struct {
	Path 		string
	Nodes 		map[string] *Node
	Client 		*clientv3.Client
}

//node is a client 
type Node struct {
	State	bool
	Key		string
	Info    ServiceInfo
}


func NewMaster(endpoints []string, watchPath string) (*Master,error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:	endpoints,
		DialTimeout: time.Second,
	})

	if err != nil {
		log.Fatal(err)
		return nil,err
	}

	master := &Master {
		Path:	watchPath,
		Nodes:	make(map[string]*Node),
		Client: cli,
	}

	go master.WatchNodes()
	return master,err
}

func (m *Master) AddNode(key string,info *ServiceInfo) {
	node := &Node{
		State:	true,
		Key:	key,
		Info:	*info,
	}

	m.Nodes[node.Key] = node
}


func GetServiceInfo(ev *clientv3.Event) *ServiceInfo {
	info := &ServiceInfo{}
	err := json.Unmarshal([]byte(ev.Kv.Value), info)
	if err != nil {
		log.Println(err)
	}
	return info
}

func (m *Master) WatchNodes()  {
	rch := m.Client.Watch(context.Background(), m.Path, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
				case clientv3.EventTypePut:
					fmt.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
					info := GetServiceInfo(ev)	
					m.AddNode(string(ev.Kv.Key),info)
				case clientv3.EventTypeDelete:
					fmt.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
					delete(m.Nodes, string(ev.Kv.Key))
			}
		}
	}
}
