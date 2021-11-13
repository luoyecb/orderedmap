package orderedmap

import (
	"testing"
	urand "util/rand"
)

func BenchmarkOrderedMap(b *testing.B) {
	om := NewOrderedMap()
	for i := 0; i < b.N; i++ {
		str := urand.RandString(4, 10)
		om.Set(str, str)
		om.Get(str)
	}
}

func BenchmarkLinkOrderedMap(b *testing.B) {
	om := NewLinkOrderedMap()
	for i := 0; i < b.N; i++ {
		str := urand.RandString(4, 10)
		om.Set(str, str)
		om.Get(str)
	}
}
