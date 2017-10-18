package sock

func findBool(key string, idx int) *rbool {
	rbooll.RLock()
	defer rbooll.RUnlock()
	v, found := rboolm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findString(key string, idx int) *rstring {
	rstringl.RLock()
	defer rstringl.RUnlock()
	v, found := rstringm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findInt(key string, idx int) *rint {
	rintl.RLock()
	defer rintl.RUnlock()
	v, found := rintm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findInt8(key string, idx int) *rint8 {
	rint8l.RLock()
	defer rint8l.RUnlock()
	v, found := rint8m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findInt16(key string, idx int) *rint16 {
	rint16l.RLock()
	defer rint16l.RUnlock()
	v, found := rint16m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findInt32(key string, idx int) *rint32 {
	rint32l.RLock()
	defer rint32l.RUnlock()
	v, found := rint32m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findInt64(key string, idx int) *rint64 {
	rint64l.RLock()
	defer rint64l.RUnlock()
	v, found := rint64m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findUint(key string, idx int) *ruint {
	ruintl.RLock()
	defer ruintl.RUnlock()
	v, found := ruintm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findUint8(key string, idx int) *ruint8 {
	ruint8l.RLock()
	defer ruint8l.RUnlock()
	v, found := ruint8m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findUint16(key string, idx int) *ruint16 {
	ruint16l.RLock()
	defer ruint16l.RUnlock()
	v, found := ruint16m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findUint32(key string, idx int) *ruint32 {
	ruint32l.RLock()
	defer ruint32l.RUnlock()
	v, found := ruint32m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findUint64(key string, idx int) *ruint64 {
	ruint64l.RLock()
	defer ruint64l.RUnlock()
	v, found := ruint64m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findByte(key string, idx int) *rbyte {
	rbytel.RLock()
	defer rbytel.RUnlock()
	v, found := rbytem[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findBytes(key string, idx int) *rbytes {
	rbytesl.RLock()
	defer rbytesl.RUnlock()
	v, found := rbytesm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findRune(key string, idx int) *rrune {
	rrunel.RLock()
	defer rrunel.RUnlock()
	v, found := rrunem[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findFloat32(key string, idx int) *rfloat32 {
	rfloat32l.RLock()
	defer rfloat32l.RUnlock()
	v, found := rfloat32m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findFloat64(key string, idx int) *rfloat64 {
	rfloat64l.RLock()
	defer rfloat64l.RUnlock()
	v, found := rfloat64m[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}

func findError(key string, idx int) *rerror {
	rerrorl.RLock()
	defer rerrorl.RUnlock()
	v, found := rerrorm[key]
	if !found {
		return nil
	}
	if idx >= len(v) || idx < 0 {
		return nil
	}
	return &v[idx]
}
