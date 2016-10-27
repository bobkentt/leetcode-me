package stack_queue

const (
    MAX_STACK_LEN = 1000
)

type mystack struct {
    var stack [MAX_STACK_LEN] int
    var vmin [MAX_STACK_LEN] int
    var len int
}

func (this *mystack) Init() {
    len = 0
    return this
}

func (this *mystack) Pop() (int, bool) {
    if len == 0 {
        return 0, false
    }
    len--
	return this.stack[len-1], true
}

func (this *mystack) Push(value int) bool {
    len++
	stack[len] = value
    if len == 1 {
        vmin[len] = value
        return true
	}
	if this.vmin > value {
		vmin[len] = value
	}
    return true
}

func (this *mystack) GetMin() (int, bool) {
    if len == 0 {
		return 0, false
	}
	return vmin[len], true
}

