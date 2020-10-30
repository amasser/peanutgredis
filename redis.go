/**
 * @Author: realpeanut
 * @Date: 2020/10/30 3:25 下午
 */
package peanutRedis

type c struct {}


type redis interface{
	Get(string)interface{}
	Set(string,string)interface{}
	GetSet(string,string)interface{}
	GetBit(string)interface{}
	MGet(string)interface{}
	SetBit(string)interface{}
	SetEx(string)interface{}
	SetNx(string)interface{}
	SetRange(string,int8,string)interface{}
	StrLen(string)interface{}
	Incr(string)interface{}
	IncrBy(string)interface{}
	Decr(string)interface{}
	DecrBy(string)interface{}
	HSet(string,string,string)interface{}
	HGet(string,string)interface{}
	//......................其他方法待添加 TODO
}

func (c c) Get(key string) interface{} {
	var client RedisCli
	conn := client.conn("localhost",6379)
	res,_:=conn.query("get" + key)
	return res
}