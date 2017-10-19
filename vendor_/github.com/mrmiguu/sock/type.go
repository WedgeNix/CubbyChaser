package sock

type keyAndIdx struct {
	key string
	idx int
}

type rbool struct {
	keyAndIdx
	r chan bool
}

type rstring struct {
	keyAndIdx
	r chan string
}

type rint struct {
	keyAndIdx
	r chan int
}

type rint8 struct {
	keyAndIdx
	r chan int8
}

type rint16 struct {
	keyAndIdx
	r chan int16
}

type rint32 struct {
	keyAndIdx
	r chan int32
}

type rint64 struct {
	keyAndIdx
	r chan int64
}

type ruint struct {
	keyAndIdx
	r chan uint
}

type ruint8 struct {
	keyAndIdx
	r chan uint8
}

type ruint16 struct {
	keyAndIdx
	r chan uint16
}

type ruint32 struct {
	keyAndIdx
	r chan uint32
}

type ruint64 struct {
	keyAndIdx
	r chan uint64
}

type rbyte struct {
	keyAndIdx
	r chan byte
}

type rbytes struct {
	keyAndIdx
	r chan []byte
}

type rrune struct {
	keyAndIdx
	r chan rune
}

type rfloat32 struct {
	keyAndIdx
	r chan float32
}

type rfloat64 struct {
	keyAndIdx
	r chan float64
}

type rerror struct {
	keyAndIdx
	r chan error
}
