package crc_test

import (
	"fmt"
	"testing"
	"unsafe"

	. "github.com/UnsafePointer/cyclic-redundancy-check-algorithms/crc"
)

func TestCRC32Check(t *testing.T) {
	crc := NewCRC32(0xF10FF0F1)
	b := []byte("Sony Computer Entertainment of America")
	expected := uint32(0x49A09045)
	check := uint32(0)

	t.Run("CRC32Check", func(t *testing.T) {
		check = crc.Calculate(b)
		if check != expected {
			t.Error(fmt.Sprintf("Expected check value %#08X, got: %#08X instead", expected, crc))
		}
	})

	// Endianness handling, might not work in your machine
	bCheck := (*[4]byte)(unsafe.Pointer(&check))[:]
	for i := len(bCheck) - 1; i >= 0; i-- {
		b = append(b, bCheck[i])
	}

	t.Run("CRCFull", func(t *testing.T) {
		check := crc.Calculate(b)
		if check != 0x0 {
			t.Error(fmt.Sprintf("Expected 0x0, got: %#08X instead", crc))
		}
	})
}
