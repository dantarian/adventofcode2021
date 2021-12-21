package bits

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/dgryski/go-bitstream"
)

func ParsePacket(input string) (Packet, error) {
	buffer, err := stringToUint8Slice(input)
	if err != nil {
		return nil, err
	}

	byteReader := bytes.NewReader(buffer)
	bitReader := bitstream.NewReader(byteReader)

	packet, _, err := newPacket(bitReader)
	return packet, err
}

func newPacket(r *bitstream.BitReader) (Packet, int, error) {
	version, err := r.ReadBits(3)
	if err != nil {
		return nil, 0, err
	}

	pktType, err := r.ReadBits(3)
	if err != nil {
		return nil, 0, err
	}

	switch pktType {
	case 4:
		return newLiteralPacket(version, r)
	default:
		return newOperatorPacket(version, pktType, r)
	}
}

func newLiteralPacket(version uint64, r *bitstream.BitReader) (*LiteralPacket, int, error) {
	continueReading := true
	value := uint64(0)
	bitsRead := 6
	for continueReading {
		chunk, err := r.ReadBits(5)
		if err != nil {
			return nil, 0, err
		}
		bitsRead += 5

		continueReading = chunk >= 16
		if continueReading {
			chunk -= 16 // remove the signal bit
		}

		value = (value << 4) + chunk
	}

	return &LiteralPacket{version: version, value: value}, bitsRead, nil
}

func newOperatorPacket(version uint64, pktType uint64, r *bitstream.BitReader) (*OperatorPacket, int, error) {
	subpacketMode, err := r.ReadBit()
	if err != nil {
		return nil, 0, err
	}

	lengthLength := 15
	if subpacketMode {
		lengthLength = 11
	}

	length, err := r.ReadBits(lengthLength)
	if err != nil {
		return nil, 0, err
	}

	var subpackets []Packet
	var bitCount int
	switch subpacketMode {
	case false:
		subpackets, bitCount, err = readSubpacketsByBitLength(int(length), r)
	case true:
		subpackets, bitCount, err = readSubpacketsByCount(int(length), r)
	}

	if err != nil {
		return nil, 0, err
	}

	return &OperatorPacket{version, pktType, subpackets}, 7 + lengthLength + bitCount, nil
}

func readSubpacketsByBitLength(length int, r *bitstream.BitReader) ([]Packet, int, error) {
	packets := []Packet{}
	i := 0
	bitsConsumed := 0
	for ; bitsConsumed < length; i++ {
		packet, bits, err := newPacket(r)
		if err != nil {
			return nil, 0, err
		}
		bitsConsumed += bits
		packets = append(packets, packet)
	}

	if bitsConsumed != length {
		return nil, 0, fmt.Errorf("expected %v bits, read %v bits", length, bitsConsumed)
	}

	return packets, length, nil
}

func readSubpacketsByCount(length int, r *bitstream.BitReader) ([]Packet, int, error) {
	packets := []Packet{}
	bitsConsumed := 0
	for i := 0; i < length; i++ {
		packet, bits, err := newPacket(r)
		if err != nil {
			return nil, 0, err
		}
		bitsConsumed += bits
		packets = append(packets, packet)
	}

	return packets, bitsConsumed, nil
}

func stringToUint8Slice(input string) ([]uint8, error) {
	buffer := []uint8{}
	var currentByte uint8
	for i, r := range input {
		val, err := strconv.ParseUint(string(r), 16, 8)
		if err != nil {
			return nil, err
		}

		if i%2 == 0 {
			currentByte = uint8(val) * uint8(16)
			continue
		}

		currentByte += uint8(val)
		buffer = append(buffer, currentByte)
	}

	// Deal with trailing characters
	if len(input)%2 == 1 {
		buffer = append(buffer, currentByte)
	}

	return buffer, nil
}
