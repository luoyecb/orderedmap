package orderedmap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testmapdata2 = []struct {
	key string
	val string
}{
	{"hello", "world"},
	{"lang", "golang"},
	{"addr", "beijing"},
	{"compony", "cc"},
}

func TestLinkOrderedMap(t *testing.T) {
	assert := assert.New(t)

	om := NewLinkOrderedMap()

	// test Set
	for _, data := range testmapdata2 {
		om.Set(data.key, data.val)
	}
	assert.Equal(om.Length(), len(testmapdata2))

	// test Exists
	for _, data := range testmapdata2 {
		assert.True(om.Exists(data.key))
	}

	// test Get
	for _, data := range testmapdata2 {
		assert.Equal(data.val, om.Get(data.key))
	}

	// test Delete
	delKey := "addr"
	om.Delete(delKey)
	assert.Equal(om.Length(), len(testmapdata2)-1)
	assert.False(om.Exists(delKey))

	fmt.Println("===")
	om.ForEach(func(k, v interface{}) {
		fmt.Printf("key=%v, val=%v\n", k, v)
	})
	fmt.Println("===")
}
