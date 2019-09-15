package crc_test

import (
	"fmt"
	"testing"

	. "github.com/UnsafePointer/cyclic-redundancy-check-algorithms/crc"
)

func TestCRC8(t *testing.T) {
	crc := NewCRC8(0xF1)
	b := []byte("SCEA")
	expected := byte(0x0A)
	check := byte(0)

	t.Run("CRC8Check", func(t *testing.T) {
		check = crc.Calculate(b)
		if check != expected {
			t.Error(fmt.Sprintf("Expected check value %#02X, got: %#02X instead", expected, crc))
		}
	})

	b = append(b, check)

	t.Run("CRC8Full", func(t *testing.T) {
		check := crc.Calculate(b)
		if check != 0x0 {
			t.Error(fmt.Sprintf("Expected 0x0, got: %#02X instead", crc))
		}
	})
}
