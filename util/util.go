package util

import (
	"fmt"
	"math"
)

func Recovers(name string) {
	if err := recover(); err != nil {
		fmt.Print("ruisRecover(" + name + "):")
		fmt.Println(err) // 这里的err其实就是panic传入的内容，55
	}
}

// BigEndian
func BigByteToInt(data []byte) int64 {
	ln := len(data)
	rt := int64(0)
	for i := 0; i < ln; i++ {
		rt |= int64(data[ln-1-i]) << (i * 8)
	}
	return rt
}
func BigIntToByte(data int64, ln int) []byte {
	rt := make([]byte, ln)
	for i := 0; i < ln; i++ {
		rt[ln-1-i] = byte(data >> (i * 8))
	}
	return rt
}

// LittleEndian
func LitByteToInt(data []byte) int64 {
	ln := len(data)
	rt := int64(0)
	for i := 0; i < ln; i++ {
		rt |= int64(data[i]) << (i * 8)
	}
	return rt
}
func LitIntToByte(data int64, ln int) []byte {
	rt := make([]byte, ln)
	for i := 0; i < ln; i++ {
		rt[i] = byte(data >> (i * 8))
	}
	return rt
}

// BigEndian
func BigByteToFloat32(data []byte) float32 {
	v := BigByteToInt(data)
	return math.Float32frombits(uint32(v))
}
func BigFloatToByte32(data float32) []byte {
	v := math.Float32bits(data)
	return BigIntToByte(int64(v), 4)
}
func BigByteToFloat64(data []byte) float64 {
	v := BigByteToInt(data)
	return math.Float64frombits(uint64(v))
}
func BigFloatToByte64(data float64) []byte {
	v := math.Float64bits(data)
	return BigIntToByte(int64(v), 8)
}

// LittleEndian
func LitByteToFloat32(data []byte) float32 {
	v := LitByteToInt(data)
	return math.Float32frombits(uint32(v))
}
func LitFloatToByte32(data float32) []byte {
	v := math.Float32bits(data)
	return LitIntToByte(int64(v), 4)
}
func LitByteToFloat64(data []byte) float64 {
	v := LitByteToInt(data)
	return math.Float64frombits(uint64(v))
}
func LitFloatToByte64(data float64) []byte {
	v := math.Float64bits(data)
	return LitIntToByte(int64(v), 8)
}
