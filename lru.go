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
		lru.l = newList(key, value)
		lru.bucket[key] = lru.l.head
		return
	}
	data := lru.bucket[key]

	if data != nil {
		lru.l.moveWithValue(data, &value)
	} else {
		nodeN := lru.l.insert(key, &value)
		lru.bucket[key] = nodeN
		if len(lru.bucket) > lru.capacity {
			lru.l.removeFirst()
			delete(lru.bucket, key)
		}
	}
}

func (lru *lRU) get(key string) *string {
	val := lru.bucket[key]
	if val == nil {
		return nil
	}
	if len(lru.bucket) > 1 {
		lru.l.move(val)
	}
	return val.data.value
}
