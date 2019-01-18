# MCBot-go

模拟MC登录协议

## 介绍
此程序是一个简易的模拟MC登录库，并不是登陆管理程序，相关的多用户模拟登陆需要额外编写goroutine来实现，但是提供一个栗子程序。

## 快速开始

对于不了解go的玩家。提供一下栗子程序的使用方法：

- 下载[release程序](https://github.com/TISUnion/MCBot-go/releases)
- 填写配置文件

```config
    [player1]    ;玩家标识，可以自定义，方便区分不同玩家
    host = localhost  ;连接地址
    port = 25565      ;连接端口
    online-mode=false ;是否为正版
    ;以下为正版参数
    account = XXXXX   ;登陆账号
    password = XXXXX  ;登陆密码
    ;以下为盗版登陆参数
    username = XXXXX  ;玩家昵称
    ;可配置多个玩家
    [player2]
    ...
    ...
```

- 启动程序，windows双击example.exe即可；linux运行example可执行文件