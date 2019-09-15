package crc

type CRC8 struct {
	Polynomial byte
}

func (crc CRC8) generateTable() [256]byte {
	table := [256]byte{}
	for i := 0; i < 256; i++ {
		currentByte := byte(i)
		for i := 0; i < 8; i++ {
			if (currentByte & 0x80) != 0 {
				currentByte = (currentByte << 1) ^ crc.Polynomial
			} else {
				currentByte <<= 1
			}
		}
		table[i] = currentByte
	}

	return table
}

func (crc CRC8) Calculate(bytes []byte) byte {
	table := crc.generateTable()
	check := byte(0)
	for _, currentByte := range bytes {
		data := currentByte ^ check
		check = table[data]
	}

	return check
}
