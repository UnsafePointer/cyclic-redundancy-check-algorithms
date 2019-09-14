package crc_test

import (
	"fmt"
	"testing"

	. "github.com/UnsafePointer/cyclic-redundancy-check-algorithms/crc"
)

func TestCRC32Check(t *testing.T) {
	b := []byte("Sony Computer Entertainment of America")
	crc := CRC32(b)
	if crc != 0x49A09045 {
		t.Error(fmt.Sprintf("Expected check value 0x49A09045, got: %#02X instead", crc))
	}
}

func TestCRC32Full(t *testing.T) {
	b := []byte("Sony Computer Entertainment of America")
	b = append(b, 0x49, 0xA0, 0x90, 0x45)
	crc := CRC32(b)
	if crc != 0x0 {
		t.Error(fmt.Sprintf("Expected 0x0, got: %#02X instead", crc))
	}
}
