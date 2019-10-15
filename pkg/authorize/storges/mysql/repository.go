package mysql

import (
	"../../../../pkg/authorize/service/auth"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var Updated = time.Now().Format("2006-01-02 15:04:05")

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

// 限制范围
func (s *Storage) Scope(ps auth.User) *Storage {

	if ps.Username != "" {
		s.DB = s.DB.Where("username=?", ps.Username)
	}
	return s
}

// get user info
func (s *Storage) User(u auth.User) (auth.User, error) {

	lu := auth.User{}

	if err := s.Scope(u).DB.Select([]string{"id"}).First(&lu).Error; err != nil {

		return lu, err

	}
	return lu, nil
}

// update user
func (s *Storage) Update(id uint, ps auth.User) (user auth.User, err error) {

	ps.Updated = Updated

	lu := auth.User{}
	// 参数1:是要修改的数据
	// 参数2:是修改的数据
	if err = s.DB.Model(&lu).Where("id=?", id).Updates(&ps).Error; err != nil {
		return
	}

	return ps, nil
}

// create user
func (s *Storage) Add(u auth.User) (id uint, err error) {

	newU := User{
		Username: u.Username,
		Password: u.Password,
		Nickname: u.Nickname,
		Created:  Updated,
		Updated:  Updated,
	}

	res := s.DB.Create(&newU)

	id = newU.ID

	if res.Error != nil {
		err = res.Error
		return
	}
	return
}
