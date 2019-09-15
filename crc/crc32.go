package crc

type CRC32 struct {
	polynomial uint32
	table      [256]uint32
}

func NewCRC32(polynomial uint32) *CRC32 {
	crc := &CRC32{
		polynomial: polynomial,
	}
	crc.table = crc.generateTable()
	return crc
}

func (crc CRC32) generateTable() [256]uint32 {
	table := [256]uint32{}
	for i := 0; i < 256; i++ {
		currentByte := uint32(i << 24)
		for i := 0; i < 8; i++ {
			if (currentByte & 0x80000000) != 0 {
				currentByte = (currentByte << 1) ^ crc.polynomial
			} else {
				currentByte <<= 1
			}
		}
		table[i] = currentByte
	}

	return table
}

func (crc CRC32) Calculate(bytes []byte) uint32 {
	check := uint32(0)
	for _, currentByte := range bytes {
		data := ((uint32(currentByte) << 24) ^ check) >> 24
		check = (check << 8) ^ crc.table[data]
	}

	return check
}
