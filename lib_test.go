/**
 * @Author: realpeanut
 * @Date: 2020/10/27 4:06 下午
 */
package peanutRedis

import (
	"testing"
)

func TestLib(t *testing.T)  {

}

func TestGetString(t *testing.T)  {
	var client redisCli
	result,err := client.conn("localhost",6379).query("get a")
	if string(result.([]uint8)) != "1" {
		t.Fatal("err")
	}
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestSetString(t *testing.T)  {
	var client redisCli
	result,err := client.conn("localhost",6379).query("set b 2")
	check := string(result.([]uint8))
	if check != "OK" {
		t.Fatal("err")
	}
	if err != nil {
		t.Fatal(err.Error())
	}
}
