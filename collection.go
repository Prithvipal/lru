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

func (dll list) insert(data *node) {}

func (dll list) move(nodeN *node) {}

func (dll list) removeFirst() {}
