package bits

func (p *OperatorPacket) Version() uint64 {
	return p.version
}

func (p *LiteralPacket) Version() uint64 {
	return p.version
}

func (p *OperatorPacket) VersionSum() uint64 {
	versionSum := p.version
	for _, packet := range p.subpackets {
		versionSum += packet.VersionSum()
	}
	return versionSum
}

func (p *LiteralPacket) VersionSum() uint64 {
	return p.version
}
