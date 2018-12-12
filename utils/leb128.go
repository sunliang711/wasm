package utils

import "io"

func DecodeUInt32(rd io.Reader) (uint32, error) {
	buf := make([]byte, 1)
	maxBits := 32
	maxBytes := (maxBits + 6) / 7
	numBytes := 0
	bytes := make([]byte, maxBytes)
	for numBytes < maxBytes {
		_, err := rd.Read(buf)
		if err != nil {
			return 0, err
		}

		byt := buf[0]
		bytes[numBytes] = byt
		numBytes += 1
		if byt&0x80 == 0 {
			break
		}
	}

	var result uint32
	for i := 0; i < maxBytes; i++ {
		result |= uint32(bytes[i]&0x7f) << uint32(i*7)
	}

	return result, nil
}

