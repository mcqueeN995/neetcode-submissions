type MyStack struct {
	queue []int
}

func Constructor()MyStack {
	return MyStack{
		queue : []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)

	if len(this.queue) == 1{
		return 
	}

	for i := 0; i < len(this.queue) - 1; i++ {
		first := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, first)
	}
}

func (this *MyStack) Pop() int {
	if len(this.queue) == 0 {
		return 0
	}
	top := this.queue[0]
	this.queue = this.queue[1:]
	return top
}

func (this *MyStack) Top() int {
	if len(this.queue) == 0 {
		return 0
	}
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
