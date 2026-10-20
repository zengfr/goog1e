package goutil

type IntUintBoolByteStringFloat interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int |
		~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
		~string |
		~bool |
		~float32 | ~float64 | ~complex64 | ~complex128 |
		~uintptr |
		any | interface{}
}
