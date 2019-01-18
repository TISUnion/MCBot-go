package datatype

type Packet struct {
	Id     int
	Length int
	Data   *[]byte
}

func (p *Packet) GetId() {
	id := VarIntInstance.DecodeByte(p.Data)
	if id == 0 {
		p.Id = VarIntInstance.DecodeByte(p.Data)
	} else {
		p.Id = id
	}
}

func (p *Packet) GetLength() {
	p.Length = VarIntInstance.DecodeByte(p.Data)
}
