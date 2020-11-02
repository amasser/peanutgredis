# peanutgredis
使用golang socket 实现的redis客户端，遵循redis标准协议，本项目主要想通过自己实现redis客户端来学习redigo、io、bufio、strconv、sync源码，学习为主,源码阅读将放在https://github.com/realpeanut/golangSourceCodeRead 项目

目前业内使用较多的是redigo package。阅读源码后，发现redigo作者自己实现了部分转码工作，而没有用官方package strconv,本项目以学习为主，后面转码功能将采用strconv package,顺便再研读一下strconv源码，后续也会增加bufio的源码阅读以及分享

- [redis协议文档](http://redisdoc.com/topic/protocol.html#id8)
# 协议实现列表
- [x] 网络层
- [x] 请求
- [x] 回复
- [x] 状态回复
- [x] 错误回复
- [x] 整数回复
- [x] 批量回复
- [ ] 多条批量回复
- [ ] 多条批量回复中的空元素
- [ ] 多命令和流水线
- [x] 内联命令
# 功能实现列表
- [x] query client
- [x] 连接池
- [ ] auth认证
