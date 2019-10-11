package discovery

import (
	"context"
	"dipole-gateway/node/common"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"log"
	"strings"
	"time"
)

type ClientDis struct {
	client        *clientv3.Client
	ctx *common.Context
}

func NewClientDis (addr []string)( *ClientDis, error){
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	if client, err := clientv3.New(conf); err == nil {
		return &ClientDis{
			client:client,
		}, nil
	} else {
		client = nil
		return nil ,err
	}
}


func (c *ClientDis) watcher(prefix string) {
	rch := c.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				c.UpdateServiceList(string(ev.Kv.Key),string(ev.Kv.Value))
			case mvccpb.DELETE:
				c.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

//func (c *ClientDis) extractAddrs(resp *clientv3.GetResponse) []string {
//	addrs := make([]string,0)
//	if resp == nil || resp.Kvs == nil {
//		return addrs
//	}
//	for i := range resp.Kvs {
//		if v := resp.Kvs[i].Value; v != nil {
//			c.SetServiceList(string(resp.Kvs[i].Key),string(resp.Kvs[i].Value))
//			addrs = append(addrs, string(v))
//		}
//	}
//	return addrs
//}

func (c *ClientDis) UpdateServiceList(key,val string) {
	service := strings.Split(key,"/")[2]
	serviceNode := strings.Split(key,"/")[3]
	c.ctx.UpdateServices(service, serviceNode, val)
	log.Println("set data key :",key,"val:",val)
}
//
func (c *ClientDis) DelServiceList(key string) {
	service := strings.Split(key,"/")[2]
	serviceNode := strings.Split(key,"/")[3]
	c.ctx.DelServices(service, serviceNode)
	log.Println("del data key:", key)
}
//
//func (c *ClientDis) SerList2Array()[]string {
//	c.lock.Lock()
//	defer c.lock.Unlock()
//	addrs := make([]string,0)
//
//	for _, v := range c.serverList {
//		addrs = append(addrs,v)
//	}
//	return addrs
//}

func (c *ClientDis) InitServices(serviceName string) (*common.Context, error) {
	resp, err := c.client.Get(context.Background(), serviceName, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var services = make(map[string]map[string]string)
	for _, v := range resp.Kvs {
		service := strings.Split(string(v.Key),"/")[2]
		serviceNode := strings.Split(string(v.Key),"/")[3]
		if services[service] == nil {
			services[service] = make(map[string]string)
		}
		services[service][serviceNode] = string(v.Value)
	}
	fmt.Println(services)
	c.ctx = common.InitContext(services)
	go c.watcher(serviceName)
	return c.ctx,nil
}