package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/yb19890724/go-im/pkg/user/service/auth"
	response "github.com/yb19890724/go-im/tools/response/json"
	"net/http"
)

// router
func Handler(ln auth.Service) http.Handler {
	
	route := httprouter.New()
	
	route.POST("/login", Login(ln))
	
	route.POST("/register", Register(ln))
	
	return route
}

func cors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Add("Access-Control-Allow-Headers","Content-Type")
	return w
}

// user login
func Login(l auth.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		
		w =cors(w)
		
		decoder := json.NewDecoder(r.Body)
		
		var rq auth.User
		
		if err := decoder.Decode(&rq); err != nil {
			
			http.Error(w, err.Error(), http.StatusBadRequest)
			
			return
		}
		
		res, err := l.Login(rq)
		
		if err != nil {
			fmt.Println("登录失败")
		}
		
		data := map[string]string{
			"token": res,
		}
		response.ResponseJson(w, "登录成功",data)
	}
	
}

// 注册
func Register(l auth.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		
		decoder := json.NewDecoder(r.Body)
		
		var rq auth.User
		
		if err := decoder.Decode(&rq); err != nil {
			
			response.BadRequest(w,err.Error())
			
			return
		}
		
		res, err := l.Register(rq)
		
		if res == 0 || err != nil {
			
			response.WithNotImplemented(w,err.Error())
			
			return
		}
		
		response.WithCreated(w)
	}
}
