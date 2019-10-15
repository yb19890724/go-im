package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/yb19890724/go-im/example/example3/pkg/connect"
	"github.com/yb19890724/go-im/example/example3/pkg/user/service/auth"
	"time"
)

var Updated = time.Now().Format("2006-01-02 15:04:05")

// struct storage db object
type Storage struct {
	DbConns *connect.MysqlConnectors
}

// created db object
func NewStorage(dbConns *connect.MysqlConnectors) *Storage {
	
	s := &Storage{
		DbConns: dbConns,
	}
	
	return s
}

// 限制范围
func (s *Storage) scope(db *gorm.DB, u auth.User) *gorm.DB {
	
	if u.ID != 0 {
		
		db = db.Where("id=?", u.ID)
	}
	return db
}

// get user info
func (s *Storage) User(u auth.User) (auth.User, error) {

	lu := auth.User{}
	
	db, err := s.DbConns.GetMysqlConnection("slave")
	
	if err != nil {
		return lu, err
	}
	
	if err := s.scope(db, u).Select([]string{"id"}).First(&lu).Error; err != nil {
		return lu, err
	}
	
	return lu, nil
}
