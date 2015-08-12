package is

import (
	"bytes"
	"regexp"
	"testing"
	"time"
	"unsafe"

	"github.com/frozzare/go-assert"
)

func TestArray(t *testing.T) {
	assert.Equal(t, Array([...]int{1, 2, 3}), true)
	assert.Equal(t, Array(false), false)
	assert.Equal(t, Array("string"), false)
}

func TestBool(t *testing.T) {
	assert.Equal(t, Bool(true), true)
	assert.Equal(t, Bool(false), true)
	assert.Equal(t, Bool([...]int{1, 2, 3}), false)
	assert.Equal(t, Bool("string"), false)
}

func TestBuffer(t *testing.T) {
	var b bytes.Buffer
	assert.Equal(t, Buffer(b), true)
	assert.Equal(t, Buffer(false), false)
	assert.Equal(t, Buffer([...]int{1, 2, 3}), false)
	assert.Equal(t, Buffer("string"), false)
}

func TestByte(t *testing.T) {
	assert.Equal(t, Byte(byte(1)), true)
	assert.Equal(t, Byte(byte(127)), true)
	assert.Equal(t, Byte(false), false)
	assert.Equal(t, Byte([...]int{1, 2, 3}), false)
	assert.Equal(t, Byte("string"), false)
}

func TestChan(t *testing.T) {
	ic := make(chan int)
	assert.Equal(t, Chan(ic), true)
	assert.Equal(t, Chan([...]int{1, 2, 3}), false)
	assert.Equal(t, Chan("string"), false)
}

func TestComplex64(t *testing.T) {
	assert.Equal(t, Complex64(complex64(0.1)), true)
	assert.Equal(t, Complex64(complex64(5.1)), true)
	assert.Equal(t, Complex64(false), false)
	assert.Equal(t, Complex64([...]int{1, 2, 3}), false)
	assert.Equal(t, Complex64("string"), false)
}

func TestComplex128(t *testing.T) {
	assert.Equal(t, Complex128(complex128(0.1)), true)
	assert.Equal(t, Complex128(complex128(5.1)), true)
	assert.Equal(t, Complex128(false), false)
	assert.Equal(t, Complex128([...]int{1, 2, 3}), false)
	assert.Equal(t, Complex128("string"), false)
}

func TestFloat32(t *testing.T) {
	assert.Equal(t, Float32(float32(0.1)), true)
	assert.Equal(t, Float32(float32(5.1)), true)
	assert.Equal(t, Float32(false), false)
	assert.Equal(t, Float32([...]int{1, 2, 3}), false)
	assert.Equal(t, Float32("string"), false)
}

func TestFloat64(t *testing.T) {
	assert.Equal(t, Float64(float64(0.1)), true)
	assert.Equal(t, Float64(float64(5.1)), true)
	assert.Equal(t, Float64(false), false)
	assert.Equal(t, Float64([...]int{1, 2, 3}), false)
	assert.Equal(t, Float64("string"), false)
}

func FuncTest() {}

func TestFunc(t *testing.T) {
	assert.Equal(t, Func(FuncTest), true)
	assert.Equal(t, Func(false), false)
	assert.Equal(t, Func([...]int{1, 2, 3}), false)
	assert.Equal(t, Func("string"), false)
}

func TestInt(t *testing.T) {
	assert.Equal(t, Int(1), true)
	assert.Equal(t, Int(100), true)
	assert.Equal(t, Int(false), false)
	assert.Equal(t, Int([...]int{1, 2, 3}), false)
	assert.Equal(t, Int("string"), false)
}

func TestInt8(t *testing.T) {
	assert.Equal(t, Int8(int8(1)), true)
	assert.Equal(t, Int8(int8(127)), true)
	assert.Equal(t, Int8(false), false)
	assert.Equal(t, Int8([...]int{1, 2, 3}), false)
	assert.Equal(t, Int8("string"), false)
}

func TestInt16(t *testing.T) {
	assert.Equal(t, Int16(int16(1)), true)
	assert.Equal(t, Int16(int16(32767)), true)
	assert.Equal(t, Int16(false), false)
	assert.Equal(t, Int16([...]int{1, 2, 3}), false)
	assert.Equal(t, Int16("string"), false)
}

func TestInt32(t *testing.T) {
	assert.Equal(t, Int32(int32(1)), true)
	assert.Equal(t, Int32(int32(1032767)), true)
	assert.Equal(t, Int32(1), false)
	assert.Equal(t, Int32(false), false)
	assert.Equal(t, Int32([...]int{1, 2, 3}), false)
	assert.Equal(t, Int32("string"), false)
}

func TestInt64(t *testing.T) {
	assert.Equal(t, Int64(int64(1)), true)
	assert.Equal(t, Int64(int64(1000032767)), true)
	assert.Equal(t, Int64(false), false)
	assert.Equal(t, Int64([...]int{1, 2, 3}), false)
	assert.Equal(t, Int64("string"), false)
}

func TestInterface(t *testing.T) {
	assert.Equal(t, Interface(false), false)
	assert.Equal(t, Interface([...]int{1, 2, 3}), false)
	assert.Equal(t, Interface("string"), false)
}

func TestMap(t *testing.T) {
	names := map[string]int{}
	assert.Equal(t, Map(names), true)
	assert.Equal(t, Map(false), false)
	assert.Equal(t, Map([...]int{1, 2, 3}), false)
	assert.Equal(t, Map("string"), false)
}

func TestNil(t *testing.T) {
	assert.Equal(t, Nil(nil), true)
	assert.Equal(t, Nil(false), false)
	assert.Equal(t, Nil([...]int{1, 2, 3}), false)
	assert.Equal(t, Nil("string"), false)
}

func TestOfStructType(t *testing.T) {
	assert.Equal(t, OfStructType(time.Now(), "time.Time"), true)
	assert.Equal(t, OfStructType("string", "string"), true)
	assert.Equal(t, OfStructType(false, "bool"), true)
	assert.Equal(t, OfStructType(false, "false"), false)
	assert.Equal(t, OfStructType([...]int{1, 2, 3}, "array"), false)
}

func TestRune(t *testing.T) {
	assert.Equal(t, Rune(rune(1)), true)
	assert.Equal(t, Rune(rune(1032767)), true)
	assert.Equal(t, Rune(1), false)
	assert.Equal(t, Rune(false), false)
	assert.Equal(t, Rune([...]int{1, 2, 3}), false)
	assert.Equal(t, Rune("string"), false)
}

func TestPtr(t *testing.T) {
	i := 0
	assert.Equal(t, Ptr(&i), true)
	assert.Equal(t, Ptr(false), false)
	assert.Equal(t, Ptr([...]int{1, 2, 3}), false)
	assert.Equal(t, Ptr("string"), false)
}

func TestRegexp(t *testing.T) {
	r, _ := regexp.Compile("p([a-z]+)ch")
	assert.Equal(t, Regexp(r), true)
	assert.Equal(t, Regexp(false), false)
	assert.Equal(t, Regexp([...]int{1, 2, 3}), false)
	assert.Equal(t, Regexp("string"), false)
}

func TestSlice(t *testing.T) {
	s := make([]string, 3)
	assert.Equal(t, Slice(s), true)
	assert.Equal(t, Slice(false), false)
	assert.Equal(t, Slice([...]int{1, 2, 3}), false)
	assert.Equal(t, Slice("string"), false)
}

func TestString(t *testing.T) {
	s := make([]string, 3)
	assert.Equal(t, String("hello"), true)
	assert.Equal(t, String(false), false)
	assert.Equal(t, String([...]int{1, 2, 3}), false)
	assert.Equal(t, String(s), false)
}

type person struct {
	name string
	age  int
}

func TestStruct(t *testing.T) {
	assert.Equal(t, Struct(person{"Bob", 20}), true)
	assert.Equal(t, Struct(false), false)
	assert.Equal(t, Struct([...]int{1, 2, 3}), false)
	assert.Equal(t, Struct("string"), false)
}

func TestTime(t *testing.T) {
	assert.Equal(t, Time(time.Now()), true)
	assert.Equal(t, Time(time.Now().UTC()), true)
	assert.Equal(t, Time(false), false)
	assert.Equal(t, Time([...]int{1, 2, 3}), false)
	assert.Equal(t, Time("string"), false)
}

func TestUint(t *testing.T) {
	assert.Equal(t, Uint(uint(1)), true)
	assert.Equal(t, Uint(uint(100)), true)
	assert.Equal(t, Uint(false), false)
	assert.Equal(t, Uint([...]int{1, 2, 3}), false)
	assert.Equal(t, Uint("string"), false)
}

func TestUint8(t *testing.T) {
	assert.Equal(t, Uint8(uint8(1)), true)
	assert.Equal(t, Uint8(uint8(127)), true)
	assert.Equal(t, Uint8(false), false)
	assert.Equal(t, Uint8([...]int{1, 2, 3}), false)
	assert.Equal(t, Uint8("string"), false)
}

func TestUint16(t *testing.T) {
	assert.Equal(t, Uint16(uint16(1)), true)
	assert.Equal(t, Uint16(uint16(1000)), true)
	assert.Equal(t, Uint16(false), false)
	assert.Equal(t, Uint16([...]int{1, 2, 3}), false)
	assert.Equal(t, Uint16("string"), false)
}

func TestUint32(t *testing.T) {
	assert.Equal(t, Uint32(uint32(1)), true)
	assert.Equal(t, Uint32(uint32(105000)), true)
	assert.Equal(t, Uint32(false), false)
	assert.Equal(t, Uint32([...]int{1, 2, 3}), false)
	assert.Equal(t, Uint32("string"), false)
}

func TestUint64(t *testing.T) {
	assert.Equal(t, Uint64(uint64(1)), true)
	assert.Equal(t, Uint64(uint64(100000000)), true)
	assert.Equal(t, Uint64(false), false)
	assert.Equal(t, Uint64([...]int{1, 2, 3}), false)
	assert.Equal(t, Uint64("string"), false)
}

func TestUintptr(t *testing.T) {
	assert.Equal(t, Uintptr(uintptr(1)), true)
	assert.Equal(t, Uintptr(uintptr(100000000)), true)
	assert.Equal(t, Uintptr(false), false)
	assert.Equal(t, Uintptr([...]int{1, 2, 3}), false)
	assert.Equal(t, Uintptr("string"), false)
}

func TestUnsafePointer(t *testing.T) {
	n := 4
	m := make([]int, n)
	assert.Equal(t, UnsafePointer(unsafe.Pointer(&m)), true)
	assert.Equal(t, UnsafePointer(false), false)
	assert.Equal(t, UnsafePointer([...]int{1, 2, 3}), false)
	assert.Equal(t, UnsafePointer("string"), false)
}
