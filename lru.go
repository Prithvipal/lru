package lru

type lRU struct {
	l        *list
	capacity int
	length   int
	bucket   map[string]*node
}

func NewLRU(capacity int) *lRU {
	return &lRU{capacity: capacity, bucket: make(map[string]*node)}
}

func (lru lRU) put(key, value string) {

}

func (lru lRU) get(key, value string) *string {
	return nil
}
