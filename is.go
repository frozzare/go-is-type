package is

import "reflect"

// Check if obj is of a specified type
func isType(obj interface{}, check reflect.Kind) bool {
	return reflect.TypeOf(obj).Kind() == check
}

// Check if obj is of struct type
func isStructType(obj interface{}, check string) bool {
	value := reflect.ValueOf(obj).Type().String()

	if len(value) == 0 {
		return false
	}

	if string(value[0]) == "*" {
		value = value[1:]
	}

	return value == check
}

// Test if obj is a array
func Array(obj interface{}) bool {
	return isType(obj, reflect.Array)
}

// Test if obj is a bool
func Bool(obj interface{}) bool {
	return isType(obj, reflect.Bool)
}

// Test if obj is of struct bytes.Buffer
func Buffer(obj interface{}) bool {
	return isStructType(obj, "bytes.Buffer")
}

// Test if obj is of byte (alias for uint8)
func Byte(obj interface{}) bool {
	return Uint8(obj)
}

// Test if obj is chan
func Chan(obj interface{}) bool {
	return isType(obj, reflect.Chan)
}

// Test if obj is Complex64
func Complex64(obj interface{}) bool {
	return isType(obj, reflect.Complex64)
}

// Test if obj is Complex128
func Complex128(obj interface{}) bool {
	return isType(obj, reflect.Complex128)
}

// Test if obj is Float32
func Float32(obj interface{}) bool {
	return isType(obj, reflect.Float32)
}

// Test if obj is Float64
func Float64(obj interface{}) bool {
	return isType(obj, reflect.Float64)
}

// Test if obj is Func
func Func(obj interface{}) bool {
	return isType(obj, reflect.Func)
}

// Test if obj is Int
func Int(obj interface{}) bool {
	return isType(obj, reflect.Int)
}

// Test if obj is Int8
func Int8(obj interface{}) bool {
	return isType(obj, reflect.Int8)
}

// Test if obj is Int16
func Int16(obj interface{}) bool {
	return isType(obj, reflect.Int16)
}

// Test if obj is Int32
func Int32(obj interface{}) bool {
	return isType(obj, reflect.Int32)
}

// Test if obj is Int64
func Int64(obj interface{}) bool {
	return isType(obj, reflect.Int64)
}

// Test if obj is Interface
func Interface(obj interface{}) bool {
	return isType(obj, reflect.Interface)
}

// Test if obj is Map
func Map(obj interface{}) bool {
	return isType(obj, reflect.Map)
}

// Test if obj is Nil
func Nil(obj interface{}) bool {
	return obj == nil
}

// Check if struct is of a struct type
func OfStructType(obj interface{}, check string) bool {
	return isStructType(obj, check)
}

// Test if obj is Ptr
func Ptr(obj interface{}) bool {
	return isType(obj, reflect.Ptr)
}

// Test if obj is of struct regexp.Regexp
func Regexp(obj interface{}) bool {
	return isStructType(obj, "regexp.Regexp")
}

// Test if obj is of rune (alias for int32)
func Rune(obj interface{}) bool {
	return Int32(obj)
}

// Test if obj is Slice
func Slice(obj interface{}) bool {
	return isType(obj, reflect.Slice)
}

// Test if obj is String
func String(obj interface{}) bool {
	return isType(obj, reflect.String)
}

// Test if obj is Struct
func Struct(obj interface{}) bool {
	return isType(obj, reflect.Struct)
}

// Test if obj is os of struct time.Time
func Time(obj interface{}) bool {
	return isStructType(obj, "time.Time")
}

// Test if obj is Uint
func Uint(obj interface{}) bool {
	return isType(obj, reflect.Uint)
}

// Test if obj is Uint8
func Uint8(obj interface{}) bool {
	return isType(obj, reflect.Uint8)
}

// Test if obj is Uint16
func Uint16(obj interface{}) bool {
	return isType(obj, reflect.Uint16)
}

// Test if obj is Uint32
func Uint32(obj interface{}) bool {
	return isType(obj, reflect.Uint32)
}

// Test if obj is Uint64
func Uint64(obj interface{}) bool {
	return isType(obj, reflect.Uint64)
}

// Test if obj is Uintptr
func Uintptr(obj interface{}) bool {
	return isType(obj, reflect.Uintptr)
}

// Test if obj is UnsafePointer
func UnsafePointer(obj interface{}) bool {
	return isType(obj, reflect.UnsafePointer)
}
