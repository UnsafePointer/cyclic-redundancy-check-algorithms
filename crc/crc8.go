package crc

func generateCRC8Table(polynomial byte) [256]byte {
	table := [256]byte{}
	for i := 0; i < 256; i++ {
		currentByte := byte(i)
		for i := 0; i < 8; i++ {
			if (currentByte & 0x80) != 0 {
				currentByte = (currentByte << 1) ^ polynomial
			} else {
				currentByte <<= 1
			}
		}
		table[i] = currentByte
	}

	return table
}

func CRC8(bytes []byte) byte {
	const polynomial = byte(0xF1) // Polynomial: (x^8 +) x^7 + x^6 + x^5 + x^4 + 1

	table := generateCRC8Table(polynomial)
	crc := byte(0)
	for _, currentByte := range bytes {
		data := currentByte ^ crc
		crc = table[data]
	}

	return crc
}
