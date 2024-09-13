package networks

type NetworkMessage struct {
	MsgBuf  []byte
	ReadPos uint16
	SizeMsg int
}

const (
	maxBytes uint16 = 16384
)

func (nm *NetworkMessage) GetU16() uint16 {
	v := uint16(nm.MsgBuf[nm.ReadPos]) | uint16(nm.MsgBuf[nm.ReadPos+1])<<8
	nm.ReadPos += 2
	return v
}

func (nm *NetworkMessage) GetU32() uint32 {
	v := uint32(nm.MsgBuf[nm.ReadPos]) |
		uint32(nm.MsgBuf[nm.ReadPos+1])<<8 |
		uint32(nm.MsgBuf[nm.ReadPos+2])<<16 |
		uint32(nm.MsgBuf[nm.ReadPos+3])<<24
	nm.ReadPos += 4
	return v
}

func (nm *NetworkMessage) GetString() string {
	stringlen := nm.GetU16()
	if stringlen >= maxBytes-nm.ReadPos {
		return ""
	}

	v := nm.MsgBuf[nm.ReadPos : nm.ReadPos+stringlen]
	nm.ReadPos += stringlen
	return string(v)
}

func (nw *NetworkMessage) SkipBytes(count uint16) { nw.ReadPos += count }

func (nw *NetworkMessage) Reset() {
	nw.MsgBuf = make([]byte, 0)

}
