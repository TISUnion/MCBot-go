package connect

import (
	"fmt"
	"net"
	"os"
	"time"

	. "github.com/TISUnion/MCBot-go/datatype"
)

/**
 * 创建连接
 * argv只传一个参数，则为昵称
 * 传二个则为账号密码
 */
func NewConnect(host string, port int, isOnline bool, argv ...string) *Connect {
	var pl *Player
	if (len(argv) == 1) && !isOnline {
		pl = &Player{Username: argv[0]}
	} else if (len(argv) == 2) && isOnline {
		pl = &Player{Account: argv[0], Password: argv[1]}
		if !pl.GetToken() { //获取账号信息
			fmt.Println("获取账号信息失败！")
			time.Sleep(5 * time.Second)
			os.Exit(1)
		}
	} else {
		fmt.Println("创建TCP连接失败, 因为：参数错误")
		return nil
	}
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Second*5)
	if err != nil {
		fmt.Println("TCP连接失败, 因为：", err)
		return nil
	}
	c := &Connect{
		Conn:     conn,
		Host:     host,
		Port:     uint16(port),
		IsOnline: isOnline,
		player:   pl,
	}
	return c
}
