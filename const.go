/**
 * @Author: realpeanut
 * @Date: 2020/10/30 5:28 下午
 */
package peanutRedis

const TCP4 = "tcp4"          //ipv4
const status_reply = '+'     //状态回复
const error_reply = '-'      //错误回复
const integer_reply = ':'    //整数回复
const bulk_reply = '$'       //批量回复
const multi_bulk_reply = '*' //多条批量回复
const redis_separator ="\r\n"//分隔符
const redis_cut = '\n'