package stack_queue

const (
    MAX_STACK_LEN = 1000
)

type mystack struct {
    stack[MAX_STACK_LEN] int
    vmin[MAX_STACK_LEN] int
    len int
}

func (this *mystack) Pop() (int, bool) {
    if this.len == 0 {
        return 0, false
    }
    this.len--
	return this.stack[this.len-1], true
}

func (this *mystack) Push(value int) bool {
    if this.len == 0 {
        this.vmin[this.len] = value
	    this.stack[this.len] = value
	} else {
        if this.vmin[this.len] > value {
            this.vmin[this.len + 1] = value
        }
        this.stack[this.len + 1] = value
    }
    this.len++
    return true
}

func (this *mystack) GetMin() (int, bool) {
    if this.len == 0 {
		return 0, false
	}
	return this.vmin[this.len], true
}

func Init() *mystack {
    return &mystack {
        len : 0,
    }
}
