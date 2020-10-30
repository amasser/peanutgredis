/**
 * @Author: realpeanut
 * @Date: 2020/10/27 4:06 下午
 */
package peanutRedis

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)


func TestGetString(t *testing.T)  {

	 var testSlice = []string{
	 	"get a",
	 	"get b",
	 	"get c",
	 	"getset getset 1",
	 	"getset getset 1",
	 }
	var client RedisCli
	conn := client.Connect("localhost",6379)
	for _,q := range testSlice {
		result,err := conn.Q(q)
		if err != nil {
			t.Fatal(err.Error())
		}
		if result != nil && reflect.TypeOf(result).String() == "[]uint8" {
			fmt.Println(string(result.([]uint8)))
		} else {
			fmt.Println(result)
		}
	}
	defer func() {
		conn.close()
	}()
}

func TestSetString(t *testing.T)  {
	var testSlice = []string{
		"set a 1",
		"set b 2",
		"set c 3",
		"del a",
	}

	var client RedisCli
	conn := client.Connect("localhost",6379)
	for _,q := range testSlice {
		result,err := conn.Q(q)
		checkType := reflect.TypeOf(result).String()
		var check interface{}
		switch checkType {
		case "int64":
			fmt.Println(result)
		case "[]uint8":
			check = string(result.([]uint8))
			fmt.Println(check)
		default:

		}

		if err != nil {
			t.Fatal(err.Error())
		}
	}
	conn.close()
}

/**
	测试发送心跳
 */
func TestPING(t *testing.T)  {
	var testSlice = []string{
		"PING",
	}
	var client RedisCli
	conn := client.Connect("localhost",6379)
	time.Sleep(time.Duration(5)*time.Second)
	for _,q := range testSlice {
		result,err := conn.Q(q)
		if err != nil {
			t.Fatal(err.Error())
		}
		if string(result.([]uint8)) != "PONG" {
			t.Fatal(err.Error())
		}
		t.Log(string(result.([]uint8)))
	}
	conn.close()
}

/**
	测试断线重连
 */
func TestReConn(t *testing.T)  {

	var testSlice = []string{
		"get a",
		"get b",
		"get c",
		"getset getset 1",
		"getset getset 1",
	}
	var client RedisCli
	conn := client.Connect("localhost",6379)
	for _,q := range testSlice {
		_,err := conn.Q(q)
		if err != nil {
			t.Fatal(err.Error())
		}
		//time.Sleep(time.Duration(3)*time.Second)
		//fmt.Println(string(result.([]uint8)))
	}
	conn.close()
}
