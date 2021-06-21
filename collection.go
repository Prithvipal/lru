package lru

type node struct {
	data entry
	next *node
	pre  *node
}

type entry struct {
	key   string
	value *string
}

type list struct {
	head *node
	tail *node
}

func newNode(key, value string) *node {
	e := entry{key: key, value: &value}
	return &node{data: e}
}

func newList(key, value string) *list {
	nodeN := newNode(key, value)
	return &list{head: nodeN, tail: nodeN}
}

func (dll *list) insert(data *node) {
	dll.tail.next = data
	data.pre = dll.tail
	dll.tail = dll.tail.next
}

func (dll *list) move(nodeN *node) {
	nodeN.pre.next = nodeN.next
	nodeN.next.pre = nodeN.pre
	dll.tail.next = nodeN
	nodeN.pre = dll.tail
	dll.tail = nodeN
}

func (dll *list) removeFirst() {
	dll.head = dll.head.next
	dll.head.pre = nil
}
