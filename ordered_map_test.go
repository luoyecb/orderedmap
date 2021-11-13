package orderedmap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testmapdata = []struct {
	key string
	val string
}{
	{"hello", "world"},
	{"lang", "golang"},
	{"addr", "beijing"},
	{"compony", "cc"},
}

func TestOrderedMap(t *testing.T) {
	assert := assert.New(t)

	om := NewOrderedMap()

	// test Set
	for _, data := range testmapdata {
		om.Set(data.key, data.val)
	}
	assert.Equal(om.Length(), len(testmapdata))

	// test Exists
	for _, data := range testmapdata {
		assert.True(om.Exists(data.key))
	}

	// test Get
	for _, data := range testmapdata {
		assert.Equal(data.val, om.Get(data.key))
	}

	// test Delete
	delKey := "addr"
	om.Delete(delKey)
	assert.Equal(om.Length(), len(testmapdata)-1)
	assert.False(om.Exists(delKey))

	fmt.Println("===")
	om.ForEach(func(k, v interface{}) {
		fmt.Printf("key=%v, val=%v\n", k, v)
	})
	fmt.Println("===")
}
