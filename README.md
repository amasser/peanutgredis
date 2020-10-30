# peanutgredis

## Overview

使用golang socket 实现的redis客户端，遵循redis标准协议。
全部功能实现后，将增加完整注释。

- [协议文档](http://redisdoc.com/topic/protocol.html#id8)
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
