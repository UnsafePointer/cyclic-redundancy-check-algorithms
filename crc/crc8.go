package crc

type CRC8 struct {
	polynomial byte
	table      [256]byte
}

func NewCRC8(polynomial byte) *CRC8 {
	crc := &CRC8{
		polynomial: polynomial,
	}
	crc.table = crc.generateTable()
	return crc
}

func (crc CRC8) generateTable() [256]byte {
	table := [256]byte{}
	for i := 0; i < 256; i++ {
		currentByte := byte(i)
		for i := 0; i < 8; i++ {
			if (currentByte & 0x80) != 0 {
				currentByte = (currentByte << 1) ^ crc.polynomial
			} else {
				currentByte <<= 1
			}
		}
		table[i] = currentByte
	}

	return table
}

func (crc CRC8) Calculate(bytes []byte) byte {
	check := byte(0)
	for _, currentByte := range bytes {
		data := currentByte ^ check
		check = crc.table[data]
	}

	return check
}
