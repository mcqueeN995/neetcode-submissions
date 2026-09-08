type MyHashSet struct {
    bucket []bool 
}

func Constructor() MyHashSet {
    return MyHashSet{
        bucket: make([]bool, 1000001),
    }
}

func (this *MyHashSet) Add(key int) {
    this.bucket[key] = true
}

func (this *MyHashSet) Remove(key int) {
    this.bucket[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
    return this.bucket[key]
}