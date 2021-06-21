package lru

type lRU struct {
	l        *list
	capacity int
	bucket   map[string]*node
}

func NewLRU(capacity int) *lRU {
	return &lRU{capacity: capacity, bucket: make(map[string]*node)}
}

func (lru *lRU) put(key, value string) {
	if lru.l == nil {
		l := newList(key, value)
		lru.bucket[key] = l.head
		return
	}
	data := lru.bucket[key]
	if data != nil && len(lru.bucket) > 1 {
		lru.l.move(data)
	} else {
		lru.l.insert(data)
		if len(lru.bucket) > lru.capacity {
			lru.l.removeFirst()
			delete(lru.bucket, key)
		}
	}
}

func (lru *lRU) get(key string) *string {
	if lru == nil {
		return nil
	}
	val := lru.bucket[key]
	if val == nil {
		return nil
	}
	if len(lru.bucket) > 1 {
		lru.l.move(val)
	}
	return val.data.value
}
