type MyQueue struct {
    in  []int
    out []int
}

func Constructor() MyQueue {
    return MyQueue{
        in:  []int{},
        out: []int{},
    }
}

func (this *MyQueue) Push(x int) {
    this.in = append(this.in, x)
}

func (this *MyQueue) move() {
    if len(this.out) == 0 {
        for i := len(this.in) - 1; i >= 0; i-- {
            this.out = append(this.out, this.in[i])
        }
        this.in = []int{}
    }
}

func (this *MyQueue) Pop() int {
    this.move()
    if len(this.out) == 0 {
        return 0 
    }
    val := this.out[len(this.out)-1]
    this.out = this.out[:len(this.out)-1]
    return val
}

func (this *MyQueue) Peek() int {
    this.move()
    if len(this.out) == 0 {
        return 0
    }
    return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
    return len(this.in) == 0 && len(this.out) == 0
}