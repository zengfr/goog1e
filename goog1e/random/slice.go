package random

// SliceAny 生成随机切片
func SliceAny(t []any, length int) []any {
	if length <= 0 {
		return []any{}
	}
	var tLen = len(t)
	var rt = make([]any, 0, length)
	for i := 0; i < length; i++ {
		rt = append(rt, t[Rand.Intn(tLen)])
	}
	return rt
}

type IntUintStringFloat interface {
	~int8 | ~int16 | ~int32 | ~int64 |
		~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~string |
		~float32 | ~float64
}

// Slice 生成随机切片
func Slice[T IntUintStringFloat](t []T, length int) []T {
	if length <= 0 {
		return []T{}
	}
	var tLen = len(t)
	var rt = make([]T, 0, length)
	for i := 0; i < length; i++ {
		rt = append(rt, t[Rand.Intn(tLen)])
	}
	return rt
}

// ShuffleAny 洗牌，打乱顺序
func ShuffleAny(t []any) []any {
	length := len(t)
	if length <= 0 || length == 1 {
		return t
	}
	for i := 0; i < length; i++ {
		index := Rand.Intn(length)
		t[i], t[index] = t[index], t[i]
	}
	return t
}

// Shuffle 洗牌，打乱顺序
func Shuffle[T IntUintStringFloat](t []T) []T {
	length := len(t)
	if length <= 0 || length == 1 {
		return t
	}
	for i := 0; i < length; i++ {
		index := Rand.Intn(length)
		t[i], t[index] = t[index], t[i]
	}
	return t
}
