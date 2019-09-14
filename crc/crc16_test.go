package crc_test

import (
	"fmt"
	"testing"

	. "github.com/UnsafePointer/cyclic-redundancy-check-algorithms/crc"
)

func TestCRC16Check(t *testing.T) {
	b := []byte("Sony Computer Entertainment of America")
	crc := CRC16(b)
	expected := uint16(0x440C)
	if crc != expected {
		t.Error(fmt.Sprintf("Expected check value %#04X, got: %#04X instead", expected, crc))
	}
}

func TestCRC16Full(t *testing.T) {
	b := []byte("Sony Computer Entertainment of America")
	b = append(b, 0x44, 0x0C)
	crc := CRC16(b)
	if crc != 0x0 {
		t.Error(fmt.Sprintf("Expected 0x0, got: %#04X instead", crc))
	}
}
