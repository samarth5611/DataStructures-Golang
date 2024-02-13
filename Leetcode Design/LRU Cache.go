type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	Head *Node
	Tail *Node
	Len  int
	M    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{}
	cache.Head = &Node{}
	cache.Tail = &Node{}
	cache.M = make(map[int]*Node)
	cache.Len = capacity
	cache.Head.Next = cache.Tail
	cache.Tail.Prev = cache.Head
	return cache
}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.M[key]; ok == true {
		cache.moveToHead(node)
		return node.Value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	node, ok := cache.M[key]
	if ok {
		node.Value = value
		cache.moveToHead(node)
		return
	}
	if len(cache.M) >= cache.Len {
		delete(cache.M, cache.Tail.Prev.Key)
		cache.deleteNode(cache.Tail.Prev)
	}
	newnode := &Node{Key: key, Value: value}
	cache.M[key] = newnode
	cache.addNode(newnode)
}

func (cache *LRUCache) moveToHead(node *Node) {
	cache.deleteNode(node)
	cache.addNode(node)
}

func (cache *LRUCache) deleteNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (cache *LRUCache) addNode(node *Node) {
	node.Prev = cache.Head
	node.Next = cache.Head.Next
	cache.Head.Next.Prev = node
	cache.Head.Next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */