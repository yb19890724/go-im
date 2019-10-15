package rest

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/yb19890724/go-im/example/example3/pkg/user/service/auth"
	response "github.com/yb19890724/go-im/example/example3/tools/response/json"
	"net/http"
	"strconv"
)

// router
func Handler(ln auth.Service) http.Handler {
	
	route := httprouter.New()
	
	route.GET("/user/:id", User(ln))
	
	return route
}

func cors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	return w
}

//
func User(l auth.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		
		ID, err := strconv.Atoi(p.ByName("id"))
		
		if err != nil {
			response.BadRequest(w,fmt.Sprintf("%s is not a valid user ID, it must be a number.", p.ByName("id")))
			return
		}
		
		var u auth.User
		
		u.ID=uint(ID)
		res, err := l.User(u)
		
		if err != nil  {
			fmt.Println(err)
			response.WithNotImplemented(w, "抱歉，未找到相关数据")
			return
		}
		
		
		response.ResponseJson(w, "请求成功", res)
		
	}
}
