package str

import (
	"strconv"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func ToInt[T Numeric](str string) T {
	i, _ := strconv.Atoi(str)
	return T(i)
}
