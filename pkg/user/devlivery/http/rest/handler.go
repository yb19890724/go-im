package rest

import (
	"../../../../../pkg/authorize/service/auth"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// router
func Handler(ln auth.Service) http.Handler {
	
	route := httprouter.New()
	
	route.GET("/user/friends", UserFriends(ln))
	
	route.GET("/user/friend/group", FriendGroup(ln))
	
	route.POST("/user/friend", MakeFriend(ln))
	
	route.POST("/user", User(ln))
	
	return route
}

func cors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Add("Access-Control-Allow-Headers","Content-Type")
	return w
}

// User Friends
func UserFriends(l auth.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		
		w = cors(w)
		
		
	}
	
}

// User Friends Group
func FriendGroup(l auth.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		
	
	}
}

//
func MakeFriend(l auth.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	
	}
}

//
func User(l auth.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	
	}
}