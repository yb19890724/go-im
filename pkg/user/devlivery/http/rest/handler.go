package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/yb19890724/go-im/pkg/user/service/auth"
	"net/http"
)

// router
func Handler( ln auth.Service) http.Handler  {
	
	route:=httprouter.New()
	
	route.POST("/login",Login(ln))
	
	return route
}

// user login
func Login(l auth.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		
		decoder := json.NewDecoder(r.Body)
		
		var rq auth.User
		
		if err := decoder.Decode(&rq);err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			
			return
		}
		
		res,err := l.Login(rq)
		
		if err != nil {
			fmt.Println("登录失败")
		}
		if res.ID != 0 {
			fmt.Println("登录成功")
		}
	}
}