package validation

import "hash/crc32"

func PayloadChecksum(values []float64) uint32 {
	buffer := make([]byte, 0, len(values)*8)
	for _, value := range values {
		bits := uint64(value * 1000)
		for shift := uint(0); shift < 64; shift += 8 {
			buffer = append(buffer, byte(bits>>shift))
		}
	}
	return crc32.ChecksumIEEE(buffer)
}
