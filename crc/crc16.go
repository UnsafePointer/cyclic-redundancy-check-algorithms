package crc

type CRC16 struct {
	polynomial uint16
	table      [256]uint16
}

func NewCRC16(polynomial uint16) *CRC16 {
	crc := &CRC16{
		polynomial: polynomial,
	}
	crc.table = crc.generateTable()
	return crc
}

func (crc CRC16) generateTable() [256]uint16 {
	table := [256]uint16{}
	for i := 0; i < 256; i++ {
		currentByte := uint16(i << 8)
		for i := 0; i < 8; i++ {
			if (currentByte & 0x8000) != 0 {
				currentByte = (currentByte << 1) ^ crc.polynomial
			} else {
				currentByte <<= 1
			}
		}
		table[i] = currentByte
	}

	return table
}

func (crc CRC16) Calculate(bytes []byte) uint16 {
	check := uint16(0)
	for _, currentByte := range bytes {
		data := ((uint16(currentByte) << 8) ^ check) >> 8
		check = (check << 8) ^ crc.table[data]
	}

	return check
}
