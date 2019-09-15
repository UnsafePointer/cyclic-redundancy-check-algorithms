package crc_test

import (
	"fmt"
	"testing"

	. "github.com/UnsafePointer/cyclic-redundancy-check-algorithms/crc"
)

func TestCRC32Check(t *testing.T) {
	crc := CRC32{
		Polynomial: 0xF10FF0F1,
	}
	b := []byte("Sony Computer Entertainment of America")
	check := crc.Calculate(b)
	expected := uint32(0x49A09045)
	if check != expected {
		t.Error(fmt.Sprintf("Expected check value %#08X, got: %#08X instead", expected, crc))
	}
}

func TestCRC32Full(t *testing.T) {
	crc := CRC32{
		Polynomial: 0xF10FF0F1,
	}
	b := []byte("Sony Computer Entertainment of America")
	b = append(b, 0x49, 0xA0, 0x90, 0x45)
	check := crc.Calculate(b)
	if check != 0x0 {
		t.Error(fmt.Sprintf("Expected 0x0, got: %#08X instead", crc))
	}
}
