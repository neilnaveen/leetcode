type MyHashMap struct {
    k map[int]int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    x := MyHashMap{}
    x.k = make(map[int]int)
    return x
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    this.k[key] = value+1
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    if this.k[key] > 0{
        return this.k[key]-1
    }
    return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    delete(this.k, key)
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */