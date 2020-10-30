/**
 * @Author: realpeanut
 * @Date: 2020/10/30 11:13 上午
 */
package peanutRedis

import (
	"net"
	"time"
)

const POOL_MEMBER  = 2



var POOL = make(chan *net.TCPConn,POOL_MEMBER)

type PoolInterface interface {
	Get() *net.TCPConn
	set() bool
}

/**
	假装这是连接池
 */
type RedisPool struct {
	pool chan *net.TCPConn
}

func ( rp *RedisPool) Get(dsn string) *net.TCPConn {
	if len(rp.pool) == 0 {
		rp.pool = POOL
		go rp.iniSet(dsn)
	}
	for  {
		select {
		case co:=<-rp.pool:
			//fmt.Println(len(rp.pool))
			time.Sleep(time.Duration(2)*time.Second)
			return co
			break
		default:

		}
	}
}

func ( rp *RedisPool) iniSet(dsn string) error {
	for i:=1 ;i<=2;i++ {
		tcpAddr, err := net.ResolveTCPAddr(TCP4, dsn)
		if err != nil{
			return err
		}
		conn, err := net.DialTCP(TCP4, nil, tcpAddr)
		if err != nil{
			return err
		}
		err = conn.SetKeepAlive(true)
		if err != nil{
			return err
		}
		rp.pool<-conn
	}
	return  nil
}

func ( rp *RedisPool) close( b *net.TCPConn)  {
	rp.pool<-b
}