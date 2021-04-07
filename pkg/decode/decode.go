package decode

import (
	"encoding/binary"
	"errors"
	"io"
)

var HashOne = [256]byte{
	0x1F, 0xCA, 0x29, 0xAC, 0x03, 0x1E, 0x2D, 0xA0, 0x27, 0xB2, 0x71, 0xD4, 0x8B, 0x86, 0xF5, 0x48,
	0x2F, 0x9A, 0xB9, 0xFC, 0x13, 0xEE, 0xBD, 0xF0, 0x37, 0x82, 0x01, 0x24, 0x9B, 0x56, 0x85, 0x98,
	0x3F, 0x6A, 0x49, 0x4C, 0x23, 0xBE, 0x4D, 0x40, 0x47, 0x52, 0x91, 0x74, 0xAB, 0x26, 0x15, 0xE8,
	0x4F, 0x3A, 0xD9, 0x9C, 0x33, 0x8E, 0xDD, 0x90, 0x57, 0x22, 0x21, 0xC4, 0xBB, 0xF6, 0xA5, 0x38,
	0x5F, 0x0A, 0x69, 0xEC, 0x43, 0x5E, 0x6D, 0xE0, 0x67, 0xF2, 0xB1, 0x14, 0xCB, 0xC6, 0x35, 0x88,
	0x6F, 0xDA, 0xF9, 0x3C, 0x53, 0x2E, 0xFD, 0x30, 0x77, 0xC2, 0x41, 0x64, 0xDB, 0x96, 0xC5, 0xD8,
	0x7F, 0xAA, 0x89, 0x8C, 0x63, 0xFE, 0x8D, 0x80, 0x87, 0x92, 0xD1, 0xB4, 0xEB, 0x66, 0x55, 0x28,
	0x8F, 0x7A, 0x19, 0xDC, 0x73, 0xCE, 0x1D, 0xD0, 0x97, 0x62, 0x61, 0x04, 0xFB, 0x36, 0xE5, 0x78,
	0x9F, 0x4A, 0xA9, 0x2C, 0x83, 0x9E, 0xAD, 0x20, 0xA7, 0x32, 0xF1, 0x54, 0x0B, 0x06, 0x75, 0xC8,
	0xAF, 0x1A, 0x39, 0x7C, 0x93, 0x6E, 0x3D, 0x70, 0xB7, 0x02, 0x81, 0xA4, 0x1B, 0xD6, 0x05, 0x18,
	0xBF, 0xEA, 0xC9, 0xCC, 0xA3, 0x3E, 0xCD, 0xC0, 0xC7, 0xD2, 0x11, 0xF4, 0x2B, 0xA6, 0x95, 0x68,
	0xCF, 0xBA, 0x59, 0x1C, 0xB3, 0x0E, 0x5D, 0x10, 0xD7, 0xA2, 0xA1, 0x44, 0x3B, 0x76, 0x25, 0xB8,
	0xDF, 0x8A, 0xE9, 0x6C, 0xC3, 0xDE, 0xED, 0x60, 0xE7, 0x72, 0x31, 0x94, 0x4B, 0x46, 0xB5, 0x08,
	0xEF, 0x5A, 0x79, 0xBC, 0xD3, 0xAE, 0x7D, 0xB0, 0xF7, 0x42, 0xC1, 0xE4, 0x5B, 0x16, 0x45, 0x58,
	0xFF, 0x2A, 0x09, 0x0C, 0xE3, 0x7E, 0x0D, 0x00, 0x07, 0x12, 0x51, 0x34, 0x6B, 0xE6, 0xD5, 0xA8,
	0x0F, 0xFA, 0x99, 0x5C, 0xF3, 0x4E, 0x9D, 0x50, 0x17, 0xE2, 0xE1, 0x84, 0x7B, 0xB6, 0x65, 0xF8,
}

var HashTwo = [256]byte{
	0x3F, 0x50, 0x95, 0x8A, 0x23, 0xFC, 0x49, 0x86, 0x67, 0x08, 0xDD, 0x62, 0x0B, 0x74, 0x51, 0x1E,
	0x0F, 0x40, 0xA5, 0xBA, 0x73, 0x6C, 0xD9, 0x36, 0x37, 0xF8, 0xED, 0x92, 0x5B, 0xE4, 0xE1, 0xCE,
	0xDF, 0x30, 0xB5, 0xEA, 0xC3, 0xDC, 0x69, 0xE6, 0x07, 0xE8, 0xFD, 0xC2, 0xAB, 0x54, 0x71, 0x7E,
	0xAF, 0x20, 0xC5, 0x1A, 0x13, 0x4C, 0xF9, 0x96, 0xD7, 0xD8, 0x0D, 0xF2, 0xFB, 0xC4, 0x01, 0x2E,
	0x7F, 0x10, 0xD5, 0x4A, 0x63, 0xBC, 0x89, 0x46, 0xA7, 0xC8, 0x1D, 0x22, 0x4B, 0x34, 0x91, 0xDE,
	0x4F, 0x00, 0xE5, 0x7A, 0xB3, 0x2C, 0x19, 0xF6, 0x77, 0xB8, 0x2D, 0x52, 0x9B, 0xA4, 0x21, 0x8E,
	0x1F, 0xF0, 0xF5, 0xAA, 0x03, 0x9C, 0xA9, 0xA6, 0x47, 0xA8, 0x3D, 0x82, 0xEB, 0x14, 0xB1, 0x3E,
	0xEF, 0xE0, 0x05, 0xDA, 0x53, 0x0C, 0x39, 0x56, 0x17, 0x98, 0x4D, 0xB2, 0x3B, 0x84, 0x41, 0xEE,
	0xBF, 0xD0, 0x15, 0x0A, 0xA3, 0x7C, 0xC9, 0x06, 0xE7, 0x88, 0x5D, 0xE2, 0x8B, 0xF4, 0xD1, 0x9E,
	0x8F, 0xC0, 0x25, 0x3A, 0xF3, 0xEC, 0x59, 0xB6, 0xB7, 0x78, 0x6D, 0x12, 0xDB, 0x64, 0x61, 0x4E,
	0x5F, 0xB0, 0x35, 0x6A, 0x43, 0x5C, 0xE9, 0x66, 0x87, 0x68, 0x7D, 0x42, 0x2B, 0xD4, 0xF1, 0xFE,
	0x2F, 0xA0, 0x45, 0x9A, 0x93, 0xCC, 0x79, 0x16, 0x57, 0x58, 0x8D, 0x72, 0x7B, 0x44, 0x81, 0xAE,
	0xFF, 0x90, 0x55, 0xCA, 0xE3, 0x3C, 0x09, 0xC6, 0x27, 0x48, 0x9D, 0xA2, 0xCB, 0xB4, 0x11, 0x5E,
	0xCF, 0x80, 0x65, 0xFA, 0x33, 0xAC, 0x99, 0x76, 0xF7, 0x38, 0xAD, 0xD2, 0x1B, 0x24, 0xA1, 0x0E,
	0x9F, 0x70, 0x75, 0x2A, 0x83, 0x1C, 0x29, 0x26, 0xC7, 0x28, 0xBD, 0x02, 0x6B, 0x94, 0x31, 0xBE,
	0x6F, 0x60, 0x85, 0x5A, 0xD3, 0x8C, 0xB9, 0xD6, 0x97, 0x18, 0xCD, 0x32, 0xBB, 0x04, 0xC1, 0x6E,
}

type Decode struct {
	Len           int
	Buffer        []byte
	OPCode        int
	DecodedBuffer []byte
	RawBuffer     []byte
}

// NewDecoder - Decode incoming traffic
func NewDecoder(reader io.Reader, buff []byte, hashPointer *int) (Decode, error) {
	var decoded Decode
	decoded.RawBuffer = buff

	ttl, err := reader.Read(buff)
	if err != nil {
		return decoded, err
	}

	if ttl < 4 {
		return decoded, errors.New("invalid packet")
	}


	temp := make([]byte, ttl)
	copy(temp, buff[0:ttl])
	crypt(temp, hashPointer)

	dc := make([]byte, len(temp)-4)
	copy(dc, temp[4:])

	decoded.Len = ttl
	decoded.OPCode = int(binary.LittleEndian.Uint16(temp[2:]))
	decoded.Buffer = temp
	decoded.DecodedBuffer = dc
	return decoded, nil
}

func crypt(packet []byte, hashPointer *int) {
	for x := 0; x < len(packet); x++ {
		packet[x] = (packet[x] ^ HashOne[*hashPointer]) ^ HashTwo[*hashPointer]
		*hashPointer = (*hashPointer + 1) % 256
	}
}

func Crypt(packet []byte, hashPointer *int) []byte {
	//hashPointer := 0
	decrypt := make([]byte, len(packet))
	for x := 0; x < len(packet); x++ {
		decrypt[x] = (packet[x] ^ HashOne[*hashPointer]) ^ HashTwo[*hashPointer]
		*hashPointer = (*hashPointer + 1) % 256
	}
	return decrypt
}

func getInt16(b []byte) int16 {
	_ = b[1]
	return int16(b[0]) | int16(b[1])<<8
}
