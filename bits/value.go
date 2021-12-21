package bits

import "math"

func (p *LiteralPacket) Value() uint64 {
	return p.value
}

func (p *OperatorPacket) Value() uint64 {
	var result uint64

	switch p.pktType {
	case 0: // Sum
		for _, sub := range p.subpackets {
			result += sub.Value()
		}
	case 1: // Product
		result = 1
		for _, sub := range p.subpackets {
			result *= sub.Value()
		}
	case 2: // Minimum
		result = math.MaxUint64
		for _, sub := range p.subpackets {
			val := sub.Value()
			if val > result {
				continue
			}
			result = val
		}
	case 3: // Maximum
		for _, sub := range p.subpackets {
			val := sub.Value()
			if val < result {
				continue
			}
			result = val
		}
	case 5: // Greater than
		if p.subpackets[0].Value() > p.subpackets[1].Value() {
			result = 1
		}
	case 6: // Less than
		if p.subpackets[0].Value() < p.subpackets[1].Value() {
			result = 1
		}
	case 7: // Equal to
		if p.subpackets[0].Value() == p.subpackets[1].Value() {
			result = 1
		}
	}

	return result
}
