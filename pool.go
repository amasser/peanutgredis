/**
 * @Author: realpeanut
 * @Date: 2020/10/30 11:13 上午
 */
package peanutRedis

import "net"

const POOL_MEMBER  = 2

var pool = make(chan *net.TCPConn,POOL_MEMBER)

type PoolInterface interface {
	Get() *net.TCPConn
	set() bool
}

/**
	假装这是连接池
 */
type RedisPool struct {}

func ( rp RedisPool) Get(dsn string) *net.TCPConn {
	connMap := <-pool
	if connMap == nil{
		go rp.iniSet(dsn)
	}
	for  {
		select {
		case co:=<-pool:
			return co
		default:

		}
	}
}

func ( rp RedisPool) iniSet(dsn string) error {
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
	pool<-conn
	return  nil
}

func ( rp RedisPool) close( b *net.TCPConn)  {
	pool<-b
}
