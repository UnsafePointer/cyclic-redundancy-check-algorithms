package crc_test

import (
	"fmt"
	"testing"

	. "github.com/UnsafePointer/cyclic-redundancy-check-algorithms/crc"
)

func TestCRC8Check(t *testing.T) {
	b := []byte("SCEA")
	crc := CRC8(b)
	if crc != 0x0A {
		t.Error(fmt.Sprintf("Expected check value 0x0A, got: %#02X instead", crc))
	}
}

func TestCRC8Full(t *testing.T) {
	b := []byte("SCEA")
	b = append(b, 0x0A)
	crc := CRC8(b)
	if crc != 0x0 {
		t.Error(fmt.Sprintf("Expected 0x0, got: %#02X instead", crc))
	}
}
