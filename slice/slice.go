package slice

type Types interface {
	~bool |
		~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func Find[T Types](slice []T, val T) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}

func HasIntersect[T Types](slice1, slice2 []T) bool {
	if (slice1 == nil) != (slice2 == nil) {
		return false
	}
	same_elemnt := Intersect(slice1, slice2)
	return len(same_elemnt) != 0
}

func Intersect[T Types](slice1, slice2 []T) []T {
	m := make(map[T]int)
	nn := make([]T, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

func IsSame[T Types](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type IntegerType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8
}

func MakeRepeat[I IntegerType, T any](length I, fill T) []T {
	s := make([]T, length)
	for i := 0; i < int(length); i++ {
		s[i] = fill
	}
	return s
}
