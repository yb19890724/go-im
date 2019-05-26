package main

import (
	"fmt"
	"github.com/yb19890724/go-im/pkg/user/devlivery/http/rest"
	"github.com/yb19890724/go-im/pkg/user/service/auth"
	"github.com/yb19890724/go-im/pkg/user/storges/mysql"
	"log"
	"net/http"
)

func main() {
	
	var login auth.Service
	
	var dbConfig string = "default:secret@tcp(192.168.1.100:3306)/default?charset=utf8mb4&parseTime=True&loc=Local&timeout=10ms"
	
	s, err := mysql.NewStorage(dbConfig)
	
	defer s.DB.Close()
	
	if err != nil {
		log.Println(err)
	}
	
	// 注入存储库aa
	
	login = auth.NewService(s)
	
	router := rest.Handler(login)
	
	fmt.Printf("The product server is on tap now: http://localhost %s", ":8080")
	log.Fatal(http.ListenAndServe(":8080", router))
	
}
