/**
 * @Author: realpeanut
 * @Date: 2020/10/27 3:57 下午
 */
package peanutRedis

import (
	"bufio"
	"io"
	"net"
	"strconv"
	"strings"
)

type Query struct {
	command string
}

type Conn struct {
	conn *net.TCPConn
	dsn  string
	pool *RedisPool
}

type redisCli struct {
	Query
	Conn
}

func (rc *redisCli) conn(host string, port int16) *redisCli {

	rc.Conn.dsn = GetDsn(host,port)

	pool := &RedisPool{}

	rc.Conn.conn = pool.Get(rc.Conn.dsn)

	rc.Conn.pool = pool

	return rc
}

func (rc *redisCli) query(command string) (interface{}, error) {
	commandSlice := strings.Split(command, " ")
	rc.Query.command = string(multi_bulk_reply) + strconv.Itoa(len(commandSlice)) + redis_separator

	for _, v := range commandSlice {

		rc.Query.command += string(bulk_reply)
		rc.Query.command += strconv.Itoa(len(v)) + redis_separator
		rc.Query.command += v + redis_separator

	}

	ln, err := rc.Conn.conn.Write([]byte(rc.Query.command))

	if err != nil || ln == 0{
		return nil, err
	}

	res, err := rc.readLineGetSizeAndReply()
	if res == nil {
		//键不存在
		return nil, err
	}

	return res, err
}

func (rc *redisCli) close() {
	rc.pool.close(rc.Conn.conn)
	rc.Conn.conn = nil
}

func (rc *redisCli) readLineGetSizeAndReply() (interface{}, error) {

	r := bufio.NewReader(rc.Conn.conn)
	p, err := r.ReadSlice(redis_cut)
	//减去 $ 和\r两个字节  例如 $5\r\n 此时p为$5\r p[i] == \r
	if string(p[1:3]) == "-1" {
		return nil, nil
	}
	p = p[:len(p)-2]

	if err == bufio.ErrBufferFull {
		buf := append([]byte{}, p...)
		//如果缓存区读完后仍然没有找到\n 则报ErrBufferFull错误，此时继续读取buff直到找到\n
		//如果找到最后任然没有找到\n 则err非空
		for err == bufio.ErrBufferFull {
			p, err = r.ReadSlice(redis_cut)
			buf = append(buf, p...)
		}
		p = buf
	}
	//此时response为$-1或者response格式❌
	if err != nil {
		return nil, nil
	}
	switch p[0] {

	case status_reply:

		switch string(p[1:]) {
		case "OK":

			return []byte("OK"), nil
		case "PONG":

			return []byte("PONG"), nil
		default:
			return p[1:], nil
		}

	case error_reply:

		return p[1:], nil

	case integer_reply:

		return ParseInt(p[1:])

	case bulk_reply:

		if len(p) > 0 {
			n, err := ParseLen(p)
			if err != nil {
				return nil,err
			}
			rs := make([]byte, n)
			ln, err := io.ReadFull(r, rs)
			if err != nil || ln == 0 {
				return nil,err
			}
			return rs, nil
		}

	case multi_bulk_reply:

		n, err := ParseLen(p[1:])
		if n < 0 || err != nil {
			return nil, err
		}
		r := make([]interface{}, n)
		for i := range r {
			r[i], err = rc.readLineGetSizeAndReply()
		}
		return r, nil
	}
	return nil, nil
}
