package main

import (
	"context"
	"dipole-gateway/node/common"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	resp, err := cli.Get(ctx, "/app01/", clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Kvs)

	//创建反向代理
	proxy := common.NewReverseProxy()

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		errc <- fmt.Errorf("%s", <-c)
	}()

	//开始监听
	go func() {
		errc <- http.ListenAndServe(":9090", proxy)
	}()

	<-errc
}