package crc

func generateCRC32Table(polynomial uint32) [256]uint32 {
	table := [256]uint32{}
	for i := 0; i < 256; i++ {
		currentByte := uint32(i << 24)
		for i := 0; i < 8; i++ {
			if (currentByte & 0x80000000) != 0 {
				currentByte = (currentByte << 1) ^ polynomial
			} else {
				currentByte <<= 1
			}
		}
		table[i] = currentByte
	}

	return table
}

func CRC32(bytes []byte) uint32 {
	const polynomial = uint32(0xF10FF0F1)

	table := generateCRC32Table(polynomial)
	crc := uint32(0)
	for _, currentByte := range bytes {
		data := ((uint32(currentByte) << 24) ^ crc) >> 24
		crc = (crc << 8) ^ table[data]
	}

	return crc
}
