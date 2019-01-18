package handletype

import (
	"bytes"

	. "github.com/TISUnion/MCBot-go/datatype"
)

type PlayerMessage struct {
}

func (*PlayerMessage) Handle(pk *Packet, pl *Player) []byte {
	buf := bytes.NewBuffer(*pk.Data)
	uuid := StringInstance.Decode(buf)
	name := StringInstance.Decode(buf)
	pl.UUID = uuid
	pl.Username = name
	// fmt.Println("登陆成功：UUID：", string(uuid), ", 用户名：", string(name))
	return []byte{}
}
