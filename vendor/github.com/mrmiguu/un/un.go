package un

import (
	"reflect"
)

func Wrap(err error) {
	if err != nil {
		panic(err)
	}
}

type unknown struct {
	v reflect.Value
}

func (u unknown) T(ptr interface{}) {
	reflect.Indirect(reflect.ValueOf(ptr)).Set(reflect.Indirect(u.v))
}

func Known(v interface{}, err error) unknown             { Wrap(err); return unknown{reflect.ValueOf(v)} }
func Bool(b bool, err error) bool                        { Wrap(err); return b }
func String(s string, err error) string                  { Wrap(err); return s }
func Int(i int, err error) int                           { Wrap(err); return i }
func Int8(i int8, err error) int8                        { Wrap(err); return i }
func Int16(i int16, err error) int16                     { Wrap(err); return i }
func Int32(i int32, err error) int32                     { Wrap(err); return i }
func Int64(i int64, err error) int64                     { Wrap(err); return i }
func Uint(u uint, err error) uint                        { Wrap(err); return u }
func Uint8(u uint8, err error) uint8                     { Wrap(err); return u }
func Uint16(u uint16, err error) uint16                  { Wrap(err); return u }
func Uint32(u uint32, err error) uint32                  { Wrap(err); return u }
func Uint64(u uint64, err error) uint64                  { Wrap(err); return u }
func Uintptr(u uintptr, err error) uintptr               { Wrap(err); return u }
func Byte(b byte, err error) byte                        { Wrap(err); return b }
func Rune(r rune, err error) rune                        { Wrap(err); return r }
func Float32(f float32, err error) float32               { Wrap(err); return f }
func Float64(f float64, err error) float64               { Wrap(err); return f }
func Complex64(c complex64, err error) complex64         { Wrap(err); return c }
func Complex128(c complex128, err error) complex128      { Wrap(err); return c }
func Bools(b []bool, err error) []bool                   { Wrap(err); return b }
func Strings(s []string, err error) []string             { Wrap(err); return s }
func Ints(i []int, err error) []int                      { Wrap(err); return i }
func Int8s(i []int8, err error) []int8                   { Wrap(err); return i }
func Int16s(i []int16, err error) []int16                { Wrap(err); return i }
func Int32s(i []int32, err error) []int32                { Wrap(err); return i }
func Int64s(i []int64, err error) []int64                { Wrap(err); return i }
func Uints(u []uint, err error) []uint                   { Wrap(err); return u }
func Uint8s(u []uint8, err error) []uint8                { Wrap(err); return u }
func Uint16s(u []uint16, err error) []uint16             { Wrap(err); return u }
func Uint32s(u []uint32, err error) []uint32             { Wrap(err); return u }
func Uint64s(u []uint64, err error) []uint64             { Wrap(err); return u }
func Uintptrs(u []uintptr, err error) []uintptr          { Wrap(err); return u }
func Bytes(b []byte, err error) []byte                   { Wrap(err); return b }
func Runes(r []rune, err error) []rune                   { Wrap(err); return r }
func Float32s(f []float32, err error) []float32          { Wrap(err); return f }
func Float64s(f []float64, err error) []float64          { Wrap(err); return f }
func Complex64s(c []complex64, err error) []complex64    { Wrap(err); return c }
func Complex128s(c []complex128, err error) []complex128 { Wrap(err); return c }
func BoolPtr(b *bool, err error) *bool                   { Wrap(err); return b }
func StringPtr(s *string, err error) *string             { Wrap(err); return s }
func IntPtr(i *int, err error) *int                      { Wrap(err); return i }
func Int8Ptr(i *int8, err error) *int8                   { Wrap(err); return i }
func Int16Ptr(i *int16, err error) *int16                { Wrap(err); return i }
func Int32Ptr(i *int32, err error) *int32                { Wrap(err); return i }
func Int64Ptr(i *int64, err error) *int64                { Wrap(err); return i }
func UintPtr(u *uint, err error) *uint                   { Wrap(err); return u }
func Uint8Ptr(u *uint8, err error) *uint8                { Wrap(err); return u }
func Uint16Ptr(u *uint16, err error) *uint16             { Wrap(err); return u }
func Uint32Ptr(u *uint32, err error) *uint32             { Wrap(err); return u }
func Uint64Ptr(u *uint64, err error) *uint64             { Wrap(err); return u }
func UintptrPtr(u *uintptr, err error) *uintptr          { Wrap(err); return u }
func BytePtr(b *byte, err error) *byte                   { Wrap(err); return b }
func RunePtr(r *rune, err error) *rune                   { Wrap(err); return r }
func Float32Ptr(f *float32, err error) *float32          { Wrap(err); return f }
func Float64Ptr(f *float64, err error) *float64          { Wrap(err); return f }
func Complex64Ptr(c *complex64, err error) *complex64    { Wrap(err); return c }
func Complex128Ptr(c *complex128, err error) *complex128 { Wrap(err); return c }
