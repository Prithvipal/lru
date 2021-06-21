package lru

import "fmt"

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

func newNode(key string, value *string) *node {
	e := entry{key: key, value: value}
	return &node{data: e}
}

func newList(key, value string) *list {
	nodeN := newNode(key, &value)
	return &list{head: nodeN, tail: nodeN}
}

func (dll *list) insert(key string, value *string) {
	nodeN := newNode(key, value)
	dll.tail.next = nodeN
	nodeN.pre = dll.tail
	dll.tail = dll.tail.next
}

func (dll *list) String() string {
	str := ""
	curr := dll.head
	for curr != nil {
		str = fmt.Sprintf("%s%s=%s,", str, curr.data.key, *curr.data.value)
		curr = curr.next
	}
	return str
}

func (dll *list) move(nodeN *node) {
	if dll.head != dll.tail {
		if nodeN.pre != nil {
			nodeN.pre.next = nodeN.next
		}
		if nodeN.next != nil {
			nodeN.next.pre = nodeN.pre
		}

		dll.tail.next = nodeN
		nodeN.pre = dll.tail
		dll.tail = nodeN
	}
}

func (dll *list) moveWithValue(nodeN *node, value *string) {
	nodeN.data.value = value
	dll.move(nodeN)
}

func (dll *list) removeFirst() {
	dll.head = dll.head.next
	dll.head.pre = nil
}
