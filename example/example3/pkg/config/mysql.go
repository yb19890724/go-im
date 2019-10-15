package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"github.com/yb19890724/go-im/example/example3/pkg/consul"
	"gopkg.in/yaml.v2"
	"log"
	"sync"
	"time"
)


type Configs struct {
	Consul *consul.ConsulConn
	Conf sync.Map
}


// mysql配置 结构体
type MysqlClusterConfig struct {
	Mysql MysqlConfig `json:"mysql"`
}

// 集群标识
type MysqlConfig struct {
	Cluster map[string]DbConf `json:"cluster"`
}


// db配置项
type DbConf struct {
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
	Dsn             string        `json:"dsn"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	MaxOpenConns    int           `json:"max_open_conns"`
}


//
type MysqlConfer interface {
	GetMysqlConfAll(dbConfName string) ([]byte, error)
	GetMysqlConf(confName string) (interface{}, bool)
	SetMysqlConf(confName string, conf []byte)
}


// 创建数据库配置
func NewMysqlConfig(consul *consul.ConsulConn) *Configs {
	return &Configs{
		Consul: consul,
	}
}

// 加载配置
func (c *Configs) LoadMysql(confName string) error {
	
	mysqlConf, err := c.GetMysqlConfAll(confName)
	
	if err != nil {
		return errors.New(fmt.Sprintf("conf get error :%s", err))
	}
	
	// 获取数据库配置
	mConf, err := c.MysqlFormatConf(mysqlConf)
	
	if err != nil {
		return errors.New(fmt.Sprintf("conf format error :%s", err))
	}
	
	for index, v := range mConf.Mysql.Cluster {
		
		c.SetMysqlConf(index, v) // 设置所有配置
		
	}
	return nil
}

// 存储指定配置
func (c *Configs) SetMysqlConf(confName string, conf DbConf) {
	c.Conf.Store(confName, conf)
}

// 获取指定配置
func (c *Configs) GetMysqlConf(confName string) (DbConf, bool) {
	
	v, ok := c.Conf.Load(confName)
	// 找不到缓存配置 去consul获取
	if !ok {
		
		_, conf, err := c.Consul.Get(c.Consul.ConfName)
		
		if err == nil { // 转化配置
			mConf, err := c.MysqlFormatConf(conf)
			
			if err == nil { // 转化成功设置配置
				c.SetMysqlConf(confName, mConf.Mysql.Cluster[confName])
				return mConf.Mysql.Cluster[confName], true
			}
		}
	}
	
	res, ok := v.(DbConf)
	return res, ok
}

// 获取所有配置项
func (c *Configs) GetMysqlConfAll(dbConf string) ([]byte, error) {
	_, conf, err := c.Consul.Get(dbConf)
	return conf, err
}


// 配置格式化
func (c *Configs ) MysqlFormatConf(conf []byte) (MysqlClusterConfig, error) {
	var mysqlConf MysqlClusterConfig
	
	err := yaml.Unmarshal(conf, &mysqlConf)
	
	if err != nil {
		return MysqlClusterConfig{}, err
	}
	return mysqlConf, err
}



// 重置配置
func (c *Configs) ResetConf(dbConfs map[string]DbConf) {
	for index, v := range dbConfs {
		c.SetMysqlConf(index, v)
	}
}

// 监控配置 变化
func (c *Configs) Watch(ctx context.Context, cancel context.CancelFunc, address string, dbConfName string, changeConf chan<- int) {
	
	watchConfig := make(map[string]interface{})
	
	watchConfig["type"] = "key"
	watchConfig["key"] = dbConfName
	watchPlan, err := watch.Parse(watchConfig)
	
	if err != nil {
		fmt.Println(err)
	}
	
	watchPlan.Handler = func(idx uint64, data interface{}) {
		
		fmt.Println("配置发生变换")
		
		d, ok := data.(*api.KVPair)
		
		if ok {
			mConf, err := c.MysqlFormatConf(d.Value)
			// 重置配置
			if err == nil {
				fmt.Println("重置配置")
				c.ResetConf(mConf.Mysql.Cluster) // 重置配置
				changeConf <- 1
			}
		}
		
	}
	
	if err := watchPlan.Run(address); err != nil {
		log.Fatalf("start watch error, error message: %s", err.Error())
	}
	
	for {
		select {
		case <-ctx.Done(): // 检测接受端是否关闭
			goto CLOSED
		}
	}

CLOSED:
	fmt.Println("监听配置变化 退出")
	// 关闭监听
	watchPlan.Stop()
	cancel() // 通知接收端关闭
}
