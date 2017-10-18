package sock

import (
	"encoding/binary"
	"errors"
	"math"
)

func int2bytes(i int) []byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(i))
	return b[:]
}

func bytes2int(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func int82bytes(i int8) []byte {
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], uint16(i))
	return b[:]
}

func bytes2int8(b []byte) int8 {
	return int8(binary.BigEndian.Uint16(b))
}

func int162bytes(i int16) []byte {
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], uint16(i))
	return b[:]
}

func bytes2int16(b []byte) int16 {
	return int16(binary.BigEndian.Uint16(b))
}

func int322bytes(i int32) []byte {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], uint32(i))
	return b[:]
}

func bytes2int32(b []byte) int32 {
	return int32(binary.BigEndian.Uint32(b))
}

func int642bytes(i int64) []byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(i))
	return b[:]
}

func bytes2int64(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(b))
}

func uint2bytes(u uint) []byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(u))
	return b[:]
}

func bytes2uint(b []byte) uint {
	return uint(binary.BigEndian.Uint64(b))
}

func uint82bytes(u uint8) []byte {
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], uint16(u))
	return b[:]
}

func bytes2uint8(b []byte) uint8 {
	return uint8(binary.BigEndian.Uint16(b))
}

func uint162bytes(u uint16) []byte {
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], u)
	return b[:]
}

func bytes2uint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func uint322bytes(u uint32) []byte {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], u)
	return b[:]
}

func bytes2uint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func uint642bytes(u uint64) []byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], u)
	return b[:]
}

func bytes2uint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func bool2bytes(b bool) []byte {
	if b {
		return []byte{1}
	}
	return []byte{0}
}

func bytes2bool(b []byte) bool {
	if len(b) > 0 && b[0] == 1 {
		return true
	}
	return false
}

func error2bytes(e error) []byte {
	if e == nil {
		return nil
	}
	return []byte(e.Error())
}

func bytes2error(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	return errors.New(string(b))
}

func rune2bytes(r rune) []byte {
	return []byte(string([]rune{r}))
}

func bytes2rune(b []byte) rune {
	return []rune(string(b))[0]
}

func float642bytes(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

func bytes2float64(b []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(b))
}

func float322bytes(f float32) []byte {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], math.Float32bits(f))
	return buf[:]
}

func bytes2float32(b []byte) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(b))
}
