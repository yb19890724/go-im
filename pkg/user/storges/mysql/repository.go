package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yb19890724/go-im/pkg/user/service/auth"
)

// struct storage db object
type Storage struct {
	DB *gorm.DB
}

// created db object
func NewStorage(dbConfig string) (*Storage, error) {
	
	var err error
	
	s := new(Storage)
	
	s.DB, err = gorm.Open("mysql", dbConfig)
	
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	
	if s.DB.Error != nil {
		fmt.Printf("database error %v", s.DB.Error)
	}
	
	return s, err
	
}

// get user info
func (s *Storage) User(params auth.User) (auth.User, error) {
	lu := auth.User{}
	err := s.DB.Select([]string{"id"}).
		Where("username = ?", params.Username).
		Where("password = ?", params.Password).
		First(&lu).Error;
	if err != nil {
		return lu, nil
	}
	return lu, err
}

// create user
func (s *Storage) Add(params auth.User) (id int, err error) {
	
	user := User{
		Username: params.Username,
		Password: params.Password,
	}
	
	result := s.DB.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return id,err
	}
	return id,nil
}