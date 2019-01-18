package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/TISUnion/MCBot-go/connect"
	"github.com/go-ini/ini"
)

func connectOffline(host string, port int, username string) *connect.Connect {
	c := connect.NewConnect(host, port, false, username)
	return c
}

func connectOnline(host string, port int, account string, password string) *connect.Connect {
	c := connect.NewConnect(host, port, true, account, password)
	return c
}

func main() {
	//加载配置文件
	Cfg, err := ini.Load("account.ini")
	if err != nil {
		fmt.Printf("读取玩家信息失败: %v", err)
		os.Exit(1)
	}
	var g sync.WaitGroup
	CfgArr := Cfg.Sections()
	//根据配置文件开启连接
	for _, section := range CfgArr {
		if section.HasKey("online-mode") {
			host := section.Key("host").String()
			port, _ := section.Key("port").Int()
			if (host == "") || (port == 0) {
				fmt.Println(section.Name, "：参数错误")
				continue
			}
			onlineMode, _ := section.Key("online-mode").Bool()
			if onlineMode {
				account := section.Key("account").String()
				password := section.Key("password").String()
				if (account == "") || (password == "") {
					fmt.Println(section.Name, "：参数错误")
					continue
				}
				c := connectOnline(host, port, account, password)
				if c != nil {
					g.Add(1)
					go c.Start()
					defer c.Close()
				}
			} else {
				username := section.Key("username").String()
				if username == "" {
					fmt.Println(section.Name, "：参数错误")
					continue
				}
				c := connectOffline(host, port, username)
				if c != nil {
					g.Add(1)
					go c.Start()
					defer c.Close()
				}
			}

		} else {
			fmt.Println("玩家：", section.Name, "未设置online-mode")
		}
	}
	g.Wait()
}
