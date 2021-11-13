package orderedmap

type LinkOrderedMap struct {
	dataMap map[interface{}]*Node
	head    *Node
	tail    *Node
}

type Node struct {
	Key  interface{}
	Val  interface{}
	prev *Node
	next *Node
}

func NewNode(k, v interface{}) *Node {
	return &Node{
		Key: k,
		Val: v,
	}
}

func NewLinkOrderedMap() *LinkOrderedMap {
	return &LinkOrderedMap{
		dataMap: make(map[interface{}]*Node),
	}
}

// Insert at the end of the linked list
func (m *LinkOrderedMap) insertNode(n *Node) {
	if m.head == nil {
		m.head = n
		m.tail = n
	} else {
		m.tail.next = n
		n.prev = m.tail
		m.tail = n
	}
}

// Remove node from linked list
func (m *LinkOrderedMap) deleteNode(n *Node) {
	if n.prev != nil {
		n.prev.next = n.next
		n.prev = nil
	} else {
		m.head = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
		n.next = nil
	} else {
		m.tail = n.prev
	}
}

func (m *LinkOrderedMap) getNode(k interface{}) *Node {
	if node, ok := m.dataMap[k]; ok {
		return node
	}
	return nil
}

func (m *LinkOrderedMap) Set(k, v interface{}) {
	if k == nil {
		panic("Key cannot be nil")
	}

	node := m.getNode(k)
	if node != nil {
		node.Val = v
		return
	}

	node = NewNode(k, v)
	m.dataMap[k] = node
	m.insertNode(node)
}

func (m *LinkOrderedMap) CheckGet(k interface{}) (interface{}, bool) {
	if k != nil {
		node := m.getNode(k)
		if node != nil {
			return node.Val, true
		}
	}
	return nil, false
}

func (m *LinkOrderedMap) Get(k interface{}) interface{} {
	val, _ := m.CheckGet(k)
	return val
}

func (m *LinkOrderedMap) Delete(k interface{}) {
	if k != nil {
		node := m.getNode(k)
		if node != nil {
			delete(m.dataMap, node.Key)
			m.deleteNode(node)
		}
	}
}

func (m *LinkOrderedMap) Exists(k interface{}) bool {
	return k != nil && m.getNode(k) != nil
}

func (m *LinkOrderedMap) Length() int {
	return len(m.dataMap)
}

func (m *LinkOrderedMap) ForEach(visitFn func(k, v interface{})) {
	if visitFn != nil {
		n := m.head
		for n != nil {
			visitFn(n.Key, n.Val)
			n = n.next
		}
	}
}

func (m *LinkOrderedMap) ReverseForEach(visitFn func(k, v interface{})) {
	if visitFn != nil {
		n := m.tail
		for n != nil {
			visitFn(n.Key, n.Val)
			n = n.prev
		}
	}
}
