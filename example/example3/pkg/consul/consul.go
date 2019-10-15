package consul

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
)



type ConsulConn struct {
	Conn *api.Client
	ConfName string
}

// 创建consul连接对象
func NewConsulServer(address string,confName string) (*ConsulConn,error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = address
	client, err := api.NewClient(consulConfig)
	
	if err != nil {
		return nil,errors.New(fmt.Sprintf("NewClient error\n%v", err))
	}
	
	cs := &ConsulConn{ // 存储consul 服务
		Conn: client,
		ConfName:confName,
	}
	return cs,nil
}

// 获取k/v
func (cs *ConsulConn) Get(k string) (string, []byte, error) {
	
	p, _, err := cs.Conn.KV().Get(k, nil)
	
	if err != nil {
		return "", []byte(""), err
	}
	return p.Key, p.Value, err
}

// 删除k/v
func (cs *ConsulConn) Delete(k string) error {
	_, err := cs.Conn.KV().Delete(k, nil)
	return err
}
