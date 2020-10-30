/**
 * @Author: realpeanut
 * @Date: 2020/10/30 12:02 下午
 */
package peanutRedis

import (
	"testing"
)

func TestPool(t *testing.T)  {
	var testSlice = []string{
		"PING",
		"PING",
	}
	var client redisCli
	for _,q := range testSlice {

		conn := client.conn("localhost",6379)
		result,err := conn.query(q)
		if err != nil {
			t.Fatal(err.Error())
		}
		if string(result.([]uint8)) != "PONG" {
			t.Fatal(err.Error())
		}
		t.Log(string(result.([]uint8)))

	}

	if len(client.pool.pool) != POOL_MEMBER {
		t.Fatal("err")
	}
	//var loop chan int
	//<-loop
}
