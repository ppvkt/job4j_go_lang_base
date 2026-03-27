package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	size int
	Head *Node
	Tail *Node
}

func NewLruCache(size int) *LruCache {
	if size < 0 {
		panic("cache size cannot be negative")
	}

	return &LruCache{
		size: size,
	}
}

func (l *LruCache) GetSize() int {
	return l.size
}

func (l *LruCache) Put(key string, value string) bool {
	if l.size <= 0 {
		return false
	}

	node := l.find(key)
	if node != nil {
		node.Value = value
		l.moveToHead(node)
		return true
	}

	newNode := &Node{Key: key, Value: value}

	if l.Len() >= l.size {
		l.removeTail()
	}

	l.insertHead(newNode)
	return true
}

func (l *LruCache) Get(key string) *string {
	node := l.find(key)

	if node == nil {
		return nil
	}

	l.moveToHead(node)

	return &node.Value

}

func (l *LruCache) moveToHead(node *Node) {
	if node == l.Head {
		return
	}

	l.remove(node)
	l.insertHead(node)
}

func (l *LruCache) insertHead(newHead *Node) {
	newHead.Next = l.Head
	newHead.Prev = nil

	if l.Head != nil {
		l.Head.Prev = newHead
	}

	l.Head = newHead

	if l.Tail == nil {
		l.Tail = newHead
	}
}

func (l *LruCache) remove(node *Node) {
	next := node.Next
	prev := node.Prev

	if prev != nil {
		prev.Next = next
	}

	if next != nil {
		next.Prev = prev
	}

	if next == nil {
		l.Tail = prev
	}

	if prev == nil {
		l.Head = next
	}
}

func (l *LruCache) removeTail() {
	if l.Tail != nil {
		l.remove(l.Tail)
	}
}

func (l *LruCache) find(key string) *Node {
	currentNode := l.Head

	for currentNode != nil {
		if currentNode.Key == key {
			return currentNode
		}
		currentNode = currentNode.Next
	}

	return nil
}

func (l *LruCache) Len() int {
	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}
