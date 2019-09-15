package crc_test

import (
	"fmt"
	"testing"
	"unsafe"

	. "github.com/UnsafePointer/cyclic-redundancy-check-algorithms/crc"
)

func TestCRC16(t *testing.T) {
	crc := CRC16{
		Polynomial: 0x00F1,
	}
	b := []byte("Sony Computer Entertainment of America")
	expected := uint16(0x440C)
	check := uint16(0)

	t.Run("CRC16Check", func(t *testing.T) {
		check = crc.Calculate(b)
		if check != expected {
			t.Error(fmt.Sprintf("Expected check value %#04X, got: %#04X instead", expected, crc))
		}
	})

	// Endianness handling, might not work in your machine
	bCheck := (*[2]byte)(unsafe.Pointer(&check))[:]
	for i := len(bCheck) - 1; i >= 0; i-- {
		b = append(b, bCheck[i])
	}

	t.Run("CRC16Full", func(t *testing.T) {
		check := crc.Calculate(b)
		if check != 0x0 {
			t.Error(fmt.Sprintf("Expected 0x0, got: %#04X instead", crc))
		}
	})
}
