package main

import (
	"github.com/dgrijalva/jwt-go"
	response "github.com/yb19890724/go-im/tools/response/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type jwtCustomClaims struct {
	jwt.StandardClaims
	
	// 追加自己需要的信息
	Uid uint `json:"uid"`
}

// 创建
func CreateToken(SecretKey []byte, issuer string, uid uint) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    issuer,// 签发者
		},
		uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}

// 解析
func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}


// 登录
func loginHandler(w http.ResponseWriter, r *http.Request) {
	
	r.ParseForm()
	
	if r.PostForm.Get("username") == "test" &&  r.PostForm.Get("password") == "test" {
		
		secretKey := make([]byte,0)
		
		token, _ := CreateToken(secretKey, "test", 2222)
		
		response.ResponseJson(w, "登录成功",map[string]string{
			"token":token,
		})
		return
	}
	
	response.ResponseJson(w ,"登录失败", map[string]string{})
}


func main() {
	http.HandleFunc("/login", loginHandler) //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
