package slice

import (
	"testing"
)

type TestA struct {
	Data int
}

func TestMakeRepeat(t *testing.T) {
	MakeRepeat(30, TestA{})
}

func BenchmarkMakeFunctions(b *testing.B) {
	b.Run("CustomMake", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice := MakeRepeat(1000000, 0)
			_ = slice
		}
	})

	b.Run("NativeMake", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice := make([]int, 1000000)
			_ = slice
		}
	})
}
