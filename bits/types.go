package bits

type Packet interface {
	Version() uint64
	VersionSum() uint64
	Value() uint64
}

type OperatorPacket struct {
	version    uint64
	pktType    uint64
	subpackets []Packet
}

type LiteralPacket struct {
	version uint64
	value   uint64
}
