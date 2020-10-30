/**
 * @Author: realpeanut
 * @Date: 2020/10/30 5:04 下午
 */
package peanutRedis

import (
	"strconv"
)

func ParseInt(p []byte) (interface{}, error) {
	if len(p) == 0 {
		return 0, nil
	}

	var negate bool
	if p[0] == '-' {
		negate = true
		p = p[1:]
		if len(p) == 0 {
			return 0, nil
		}
	}

	var n int64
	for _, b := range p {
		n *= 10
		if b < '0' || b > '9' {
			return 0, nil
		}
		n += int64(b - '0')
	}

	if negate {
		n = -n
	}
	return n, nil
}

func ParseLen(p []byte) (int, error) {


	if len(p) == 0 {

		return -1, nil
	}

	if p[0] == '-' && len(p) == 2 && p[1] == '1' {
		return -1, nil
	}

	var n int
	for _, b := range p {
		n *= 10
		if b < '0' || b > '9' {
			return -1, nil
		}
		n += int(b - '0')
	}

	return n, nil
}

func GetDsn(host string,port int16) string {
	return host + ":" + Int16ToString(port)
}

func Int16ToString(c int16) string {
	return strconv.FormatInt(int64(c), 10)
}