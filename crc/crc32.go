package crc

type CRC32 struct {
	Polynomial uint32
}

func (crc CRC32) generateTable() [256]uint32 {
	table := [256]uint32{}
	for i := 0; i < 256; i++ {
		currentByte := uint32(i << 24)
		for i := 0; i < 8; i++ {
			if (currentByte & 0x80000000) != 0 {
				currentByte = (currentByte << 1) ^ crc.Polynomial
			} else {
				currentByte <<= 1
			}
		}
		table[i] = currentByte
	}

	return table
}

func (crc CRC32) Calculate(bytes []byte) uint32 {
	table := crc.generateTable()
	check := uint32(0)
	for _, currentByte := range bytes {
		data := ((uint32(currentByte) << 24) ^ check) >> 24
		check = (check << 8) ^ table[data]
	}

	return check
}
