package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// create repository interface
type Repository interface {
	User(ps User) (User, error)
	Add(ps User) (uint, error)
	Update(id uint, ps User) (User, error)
}

// create service interface
type Service interface {
	Login(ps User) (string, error)
	Register(ps User) (uint, error)
}

// struct service storage repository
type service struct {
	GR Repository // 寄存器
}

type jwtCustomClaims struct {
	jwt.StandardClaims
	
	// 追加自己需要的信息
	Uid uint   `json:"uid"`
	Up  string `json:"updated"`
}

// new service
func NewService(r Repository) Service {
	return &service{r}
}

// create token
func createToken(uid uint, up string) (token string, err error) {
	var sk []byte
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
		},
		uid,
		up,
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tk.SignedString(sk)
	return
}

// service call repository (user and update)
func (s *service) Login(ps User) (string, error) {
	
	res, err := s.GR.User(ps)
	
	if res.Token != "" {
		
		// 获取token
		tk, err := createToken(uint(res.ID), res.Updated)
		
		if err != nil {
			log.Fatalf("create token err:%s", err)
		}
		
		// 更新数据库 token
		res, err := s.GR.Update(res.ID, User{
			Token: tk,
		})
		
		return res.Token, nil
	}
	return "", err
}

// service call repository Add
func (s *service) Register(ps User) (uint, error) {
	
	User, err := s.GR.User(ps)
	
	// 查询记录 User.ID==0 && err.Error()== "record not found" 没有记录
	if 0 != User.ID || err.Error() != "record not found" { // 避免重复注册
	
		return 0, err
		
	}
	
	return s.GR.Add(ps)
}
