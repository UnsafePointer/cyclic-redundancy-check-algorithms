package crc_test

import (
	"fmt"
	"testing"

	. "github.com/UnsafePointer/cyclic-redundancy-check-algorithms/crc"
)

func TestCRC8Check(t *testing.T) {
	crc := CRC8{
		Polynomial: 0xF1,
	}
	b := []byte("SCEA")
	check := crc.Calculate(b)
	expected := byte(0x0A)
	if check != expected {
		t.Error(fmt.Sprintf("Expected check value %#02X, got: %#02X instead", expected, crc))
	}
}

func TestCRC8Full(t *testing.T) {
	crc := CRC8{
		Polynomial: 0xF1,
	}
	b := []byte("SCEA")
	b = append(b, 0x0A)
	check := crc.Calculate(b)
	if check != 0x0 {
		t.Error(fmt.Sprintf("Expected 0x0, got: %#02X instead", crc))
	}
}
