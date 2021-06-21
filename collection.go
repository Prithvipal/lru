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

func (dll list) insert(data int) {}

func (dll list) move(data int) {}

func (dll list) remove(data int) {}
