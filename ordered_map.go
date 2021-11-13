package orderedmap

type OrderedMap struct {
	dataMap map[interface{}]*Entry
	keys    []interface{}
}

type Entry struct {
	Val   interface{}
	index int
}

func NewEntry(v interface{}, idx int) *Entry {
	return &Entry{
		Val:   v,
		index: idx,
	}
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		dataMap: make(map[interface{}]*Entry),
		keys:    make([]interface{}, 0),
	}
}

func (m *OrderedMap) getEntry(k interface{}) *Entry {
	if ent, ok := m.dataMap[k]; ok {
		return ent
	}
	return nil
}

func (m *OrderedMap) Set(k, v interface{}) {
	if k == nil {
		panic("Key cannot be nil")
	}

	entry := m.getEntry(k)
	if entry != nil {
		entry.Val = v
		return
	}

	m.dataMap[k] = NewEntry(v, len(m.keys))
	m.keys = append(m.keys, k)
}

func (m *OrderedMap) CheckGet(k interface{}) (interface{}, bool) {
	if k != nil {
		entry := m.getEntry(k)
		if entry != nil {
			return entry.Val, true
		}
	}
	return nil, false
}

func (m *OrderedMap) Get(k interface{}) interface{} {
	val, _ := m.CheckGet(k)
	return val
}

func (m *OrderedMap) Delete(k interface{}) {
	if k != nil {
		entry := m.getEntry(k)
		if entry != nil {
			length := len(m.keys)
			index := entry.index

			copy(m.keys[index:], m.keys[index+1:])
			m.keys = m.keys[:length-1]
			delete(m.dataMap, k)
		}
	}
}

func (m *OrderedMap) Exists(k interface{}) bool {
	return k != nil && m.getEntry(k) != nil
}

func (m *OrderedMap) Length() int {
	return len(m.keys)
}

func (m *OrderedMap) ForEach(visitFn func(k, v interface{})) {
	if visitFn != nil {
		for _, key := range m.keys {
			visitFn(key, m.dataMap[key].Val)
		}
	}
}

func (m *OrderedMap) ReverseForEach(visitFn func(k, v interface{})) {
	if visitFn != nil {
		for j := len(m.keys) - 1; j >= 0; j-- {
			k := m.keys[j]
			visitFn(k, m.dataMap[k].Val)
		}
	}
}
