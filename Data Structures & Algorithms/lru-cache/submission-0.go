type Node struct {
	key int
	val int
	next *Node
	prev *Node
}

type LRUCache struct {
   capacity int
   dict map[int]*Node
   head *Node
   tail *Node 
}

func Constructor(capacity int) LRUCache {
   h := &Node{0,0,nil, nil} 
   t := &Node{0,0,nil, nil} 
   h.next = t
   t.prev = h
   return LRUCache{capacity, map[int]*Node{}, h, t}
}


func (this *LRUCache) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
}

func (this *LRUCache) add(n *Node) {
	temp := this.head.next 
	this.head.next = n
	n.prev = this.head
	n.next = temp
	temp.prev = n
}

func (this *LRUCache) Get(key int) int {
	v,ok := this.dict[key]
	if !ok{
		return -1
	}
	this.remove(v)
	this.add(v)
	return v.val
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity <= 0 {
        return
    }

    // Update an existing key.
    if n, ok := this.dict[key]; ok {
        this.remove(n)
        delete(this.dict, key)
    }

    // Evict the least recently used entry.
    if len(this.dict) == this.capacity {
        lru := this.tail.prev
        this.remove(lru)
        delete(this.dict, lru.key)
    }

    // Insert the new entry as most recently used.
    n := &Node{key: key, val: value}
    this.add(n)
    this.dict[key] = n
}
