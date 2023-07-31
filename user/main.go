package main

import (
	"douyinv2/user/dal/db"
	user "douyinv2/user/kitex_gen/user/apiservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	db.Init()
	r, err := etcd.NewEtcdRegistry([]string{"192.168.217.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	svr := user.NewServer(new(ApiServiceImpl), server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Userservice"}),
		server.WithRegistry(r))
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
