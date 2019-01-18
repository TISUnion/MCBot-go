# MCBot-go

模拟MC登录协议

## 介绍

Minecraft版本 1.13.2

参考项目：https://github.com/ammaraskar/pyCraft

Minecraft协议：https://wiki.vg/Protocol

此程序是一个简易的模拟MC登录库，并不是登陆管理程序，相关的多用户模拟登陆需要额外编写goroutine来实现，但是提供一个栗子程序。

之后无重大BUG将不在升级该库

如果想深度开发，并且熟悉python，还是推荐pyCraft


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

## 使用方法

使用connect.NewConnect方法可获得一个连接对象，如果isOnline为true则表示为正版登陆，argv需要传入2个参数，账号以及密码，否则表示为正盗版登陆，只需要传入一个参数，玩家的昵称。
``` golang
connect.NewConnect(host string, port int, isOnline bool, argv ...string) *connect.Connect
```
创建完*connect.Connect后，如果创建失败则会返回nil，然后要启动连接，使用start函数即可启动，值得一提的是，start函数是一个死循环函数，如果需要多玩家管理，则需要自行设计使用go程。
```golang
if c != nil {
    // go c.Start()  
    c.Start()
    defer c.Close()
}
```
