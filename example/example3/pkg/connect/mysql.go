package connect

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yb19890724/go-im/example/example3/pkg/config"
	"log"
	"sync"
)

var mysqlConnectors *MysqlConnectors

// 连接器结构体
type MysqlConnectors struct {
	Conf *config.Configs
	rwMute sync.RWMutex
	Conns map[string]*gorm.DB
}

// new service
func NewMysqlConnectors(conf *config.Configs) (*MysqlConnectors) {
	
	mysqlConnectors := &MysqlConnectors{
		Conf: conf,
		Conns:make(map[string]*gorm.DB),
	}
	
	return mysqlConnectors
}


// 创建mysql连接
func (mc *MysqlConnectors) CreateMysqlConnection(conf config.DbConf) (*gorm.DB, error) {
	
	db, err := gorm.Open("mysql", conf.Dsn)
	
	if err != nil {
		
		return nil, errors.New(fmt.Sprintf("connect mysql fail %s", err))
	}
	// 设置连接池
	db.DB().SetMaxIdleConns(conf.MaxIdleConns)
	db.DB().SetMaxOpenConns(conf.MaxOpenConns)
	db.DB().SetConnMaxLifetime(conf.ConnMaxLifetime)
	db.SingularTable(true)
	db.BlockGlobalUpdate(false)
	return db, nil
}

// 获取指定 db连接
func (mc *MysqlConnectors) GetDbConn(dbName string) (*gorm.DB, bool) {
	mc.rwMute.RLock()
	dbConn, ok := mc.Conns[dbName]
	mc.rwMute.RUnlock()
	return dbConn, ok
}

// 存储db连接
func (mc *MysqlConnectors) SetDbsConn(dbName string, db *gorm.DB) {
	mc.rwMute.Lock()
	mc.Conns[dbName] = db
	mc.rwMute.Unlock()
}

// 删除db连接
func (mc *MysqlConnectors) DelDbsConn(dbName string) {
	mc.rwMute.Lock()
	delete(mc.Conns, dbName)
	mc.rwMute.Unlock()
}

// 卸载db连接
func (mc *MysqlConnectors) ResetDbsConn(ctx context.Context,cancel context.CancelFunc,changConf <-chan int)  {
	
	for {
		select {
		case <-changConf:
			
			fmt.Println("重置db")
			
			for index, v := range mc.Conns{
				
				// 延时操作
				err :=v.Close()// 关闭数据库连接
				
				mc.DelDbsConn(index)// 删除连接存储位
				
				if err !=nil {
					log.Printf("db close err %s",err)
					goto ERROR
				}
			}
		case <-ctx.Done():
			goto CLOSED
		}
	}
ERROR:
	cancel()
CLOSED:
	fmt.Println("重置机制退出")
}

// 获取mysql配置
func (mc *MysqlConnectors) GetMysqlConnection(dbName string) (*gorm.DB, error) {
	dbConn, ok := mc.GetDbConn(dbName)
	
	if !ok{
		// 获取配置
		conf, ok := mc.Conf.GetMysqlConf(dbName)
		
		if !ok {
			return nil, errors.New("get mysql config is nil")
		}
		
		// 创建连接
		dbConn, err := mc.CreateMysqlConnection(conf)
	
		if err != nil {
			return nil, err
		}
		
		mc.SetDbsConn(dbName, dbConn)
		
		return dbConn,nil
	}
	
	return dbConn,nil
}



