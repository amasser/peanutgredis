/**
 * @Author: realpeanut
 * @Date: 2020/10/30 12:02 下午
 */
package peanutRedis

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T)  {
	var testSlice = []string{
		"PING",
		"PING",
		"PING",
	}
	var client RedisCli

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
	fmt.Println(len(client.pool.pool))
	if len(client.pool.pool) != POOL_MEMBER {
		t.Fatal("err")
	}
}


