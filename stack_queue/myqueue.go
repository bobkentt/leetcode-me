package stack_queue

const (
    MAX_STACK_LEN = 1000
)

type queue struct {
    s1[MAX_STACK_LEN] int
    s2[MAX_STACK_LEN] int
    len int
}

func Init() *mystack {
    return &mystack {
        len : 0,
    }
}
