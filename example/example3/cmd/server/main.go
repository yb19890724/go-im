package main

import (
	"fmt"
	"github.com/yb19890724/go-im/example/example3/pkg/config"
	"github.com/yb19890724/go-im/example/example3/pkg/connect"
	"github.com/yb19890724/go-im/example/example3/pkg/consul"
	"github.com/yb19890724/go-im/example/example3/pkg/devlivery/http/rest"
	"github.com/yb19890724/go-im/example/example3/pkg/user/service/auth"
	"github.com/yb19890724/go-im/example/example3/pkg/user/storage"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`   // 列名为 `id`
	Name string `json:"name"` // 列名为 `username`
}

var Users []User

func main() {
	
	// 加载consul服务
	consulClient, _ := consul.NewConsulServer("http://localhost:8500","/cluster/database")
	
	conf := config.NewMysqlConfig(consulClient) // 创建配置协议
	
	err := conf.LoadMysql("/cluster/database")
	
	if err!= nil { // 预热配置 缓存
		fmt.Println("auth server run err :%s",err)
		return
	}
	
	dbConns := connect.NewMysqlConnectors(conf) // 加载数据库连接
	
	
	// 监听配置 是否发生变化做热处理
	/*changConf:=make(chan int,1)
	ctx, cancel := context.WithCancel(context.Background())
	// 监听k/v是否发生变
	conf.Watch(ctx, cancel, "http://localhost:8500", "/cluster/database", changConf)
	dbConns.ResetDbsConn(ctx, cancel,changConf)
	*/
	// -----------------
	
	
	s := storage.NewStorage(dbConns) // 存储
	u := auth.NewService(s)     // 服务
	
	fmt.Printf("The auth server is on tap now: http://localhost %s", ":8080")
	
	router := rest.Handler(u)
	log.Fatal(http.ListenAndServe(":8080", router))

}
