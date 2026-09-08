type Pair struct {
	key   int
	value int
}

type MyHashMap struct {
	bucket [10007][]Pair
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func hash(key int) int {
	return key % 10007
}

func (this *MyHashMap) Put(key int, value int) {
	h := hash(key)
	
	for i, p := range this.bucket[h] {
		if p.key == key {
			this.bucket[h][i].value = value 
			return
		}
	}
	this.bucket[h] = append(this.bucket[h], Pair{key, value})
}

func (this *MyHashMap) Get(key int) int {
	h := hash(key)
	for _, p := range this.bucket[h] {
		if p.key == key {
			return p.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := hash(key)
	slice := this.bucket[h]
	
	for i, p := range slice {
		if p.key == key {
			this.bucket[h] = append(slice[:i], slice[i+1:]...)
			return
		}
	}
}