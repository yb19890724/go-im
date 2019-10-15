package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	response "github.com/yb19890724/go-im/tools/response/json"
	"github.com/yb19890724/go-im/example/example2/model"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
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
			Issuer:    issuer, // 签发者
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

type loginRequest struct {
	Username string
	Password string
}

//  注册
func registerHandler(w http.ResponseWriter, r *http.Request) {
	w = cors(w)

	decoder := json.NewDecoder(r.Body)

	var lr loginRequest

	if err := decoder.Decode(&lr); err != nil {

		response.BadRequest(w,"抱歉！参数错误")

		return
	}

	if lr.Username != "" && lr.Password != "" {

		var user model.User

		user.Username = lr.Username
		user.Password = lr.Password
		user.Nickname = fmt.Sprintf("user%d", rand.Int())

		_, err := user.Insert()

		if err != nil {
			response.ResponseJson(w, "注册成功，返回登录页", map[string]string{})
		}

	}

	response.ResponseJson(w, "注册失败", map[string]string{})

	return
}

// 登录
func loginHandler(w http.ResponseWriter, r *http.Request) {

	w = cors(w)

	decoder := json.NewDecoder(r.Body)

	var lr loginRequest

	if err := decoder.Decode(&lr); err != nil {

		response.BadRequest(w,"抱歉！参数错误")

		return
	}

	var user model.User
	user.Username = lr.Username
	user.Password = lr.Password

	err := user.Login()
	if err != nil {
		response.ResponseJson(w, "账号或密码错误", map[string]string{})
	}

	secretKey := make([]byte, 0)
	token, _ := CreateToken(secretKey, "im", uint(user.ID))

	user.Token = token
	err = user.Update(user.ID)

	if err == nil {
		response.ResponseJson(w, "登录成功", map[string]string{
			"token": token,
		})
	}

	response.ResponseJson(w, "登录失败", map[string]string{})
}

// 跨域
func cors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	return w
}

func main() {
	http.HandleFunc("/login", loginHandler)  //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
