package crc

func CRC8(bytes []byte) byte {
	const polynomial = byte(0xF1) // Polynomial: (x^8 +) x^7 + x^6 + x^5 + x^4 + 1

	crc := byte(0)
	for _, currentByte := range bytes {
		crc ^= currentByte
		for i := 0; i < 8; i++ {
			if (crc & 0x80) != 0 {
				crc = (crc << 1) ^ polynomial
			} else {
				crc <<= 1
			}
		}
	}

	return crc
}
