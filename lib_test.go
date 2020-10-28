/**
 * @Author: realpeanut
 * @Date: 2020/10/27 4:06 下午
 */
package peanutRedis

import (
	"fmt"
	"reflect"
	"testing"
)


func TestGetString(t *testing.T)  {

	 var testSlice = []string{
	 	"get a",
	 	"get b",
	 	"get c",
	 	"getset getset 1",
	 	"getset getset 1",
	 }
	var client redisCli
	conn := client.conn("localhost",6379)
	for _,q := range testSlice {
		result,err := conn.query(q)
		if err != nil {
			t.Fatal(err.Error())
		}
		fmt.Println(string(result.([]uint8)))
	}
	conn.close()
}

func TestSetString(t *testing.T)  {
	var testSlice = []string{
		"set a 1",
		"set b 2",
		"set c 3",
		"del a",
	}

	var client redisCli
	conn := client.conn("localhost",6379)
	for _,q := range testSlice {
		result,err := conn.query(q)
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

