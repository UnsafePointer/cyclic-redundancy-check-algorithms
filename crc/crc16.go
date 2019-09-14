package crc

func generateCRC16Table(polynomial uint16) [256]uint16 {
	table := [256]uint16{}
	for i := 0; i < 256; i++ {
		currentByte := uint16(i << 8)
		for i := 0; i < 8; i++ {
			if (currentByte & 0x8000) != 0 {
				currentByte = (currentByte << 1) ^ polynomial
			} else {
				currentByte <<= 1
			}
		}
		table[i] = currentByte
	}

	return table
}

func CRC16(bytes []byte) uint16 {
	const polynomial = uint16(0x00F1)

	table := generateCRC16Table(polynomial)
	crc := uint16(0)
	for _, currentByte := range bytes {
		data := ((uint16(currentByte) << 8) ^ crc) >> 8
		crc = (crc << 8) ^ table[data]
	}

	return crc
}
