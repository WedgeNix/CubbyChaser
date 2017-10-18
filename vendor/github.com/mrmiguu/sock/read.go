package sock

import (
	"bytes"
	"errors"
)

func read(pkt []byte) error {
	parts := bytes.Split(pkt, v)
	if len(parts) != 4 || len(parts[0]) < 1 {
		return errors.New("invalid packet")
	}
	t, key, idx, body := parts[0][0], string(parts[1]), bytes2int(parts[2]), parts[3]

	switch t {
	case Tbool:
		B := findBool(key, idx)
		if B != nil {
			B.r <- bytes2bool(body)
		}
	case Tstring:
		S := findString(key, idx)
		if S != nil {
			S.r <- string(body)
		}
	case Tint:
		I := findInt(key, idx)
		if I != nil {
			I.r <- bytes2int(body)
		}
	case Tint8:
		I := findInt8(key, idx)
		if I != nil {
			I.r <- bytes2int8(body)
		}
	case Tint16:
		I := findInt16(key, idx)
		if I != nil {
			I.r <- bytes2int16(body)
		}
	case Tint32:
		I := findInt32(key, idx)
		if I != nil {
			I.r <- bytes2int32(body)
		}
	case Tint64:
		I := findInt64(key, idx)
		if I != nil {
			I.r <- bytes2int64(body)
		}
	case Tuint:
		U := findUint(key, idx)
		if U != nil {
			U.r <- bytes2uint(body)
		}
	case Tuint8:
		U := findUint8(key, idx)
		if U != nil {
			U.r <- bytes2uint8(body)
		}
	case Tuint16:
		U := findUint16(key, idx)
		if U != nil {
			U.r <- bytes2uint16(body)
		}
	case Tuint32:
		U := findUint32(key, idx)
		if U != nil {
			U.r <- bytes2uint32(body)
		}
	case Tuint64:
		U := findUint64(key, idx)
		if U != nil {
			U.r <- bytes2uint64(body)
		}
	case Tbyte:
		B := findByte(key, idx)
		if B != nil {
			B.r <- body[0]
		}
	case Tbytes:
		B := findBytes(key, idx)
		if B != nil {
			B.r <- body
		}
	case Trune:
		R := findRune(key, idx)
		if R != nil {
			R.r <- bytes2rune(body)
		}
	case Tfloat32:
		F := findFloat32(key, idx)
		if F != nil {
			F.r <- bytes2float32(body)
		}
	case Tfloat64:
		F := findFloat64(key, idx)
		if F != nil {
			F.r <- bytes2float64(body)
		}
	case Terror:
		E := findError(key, idx)
		if E != nil {
			E.r <- bytes2error(body)
		}
	}

	return nil
}
