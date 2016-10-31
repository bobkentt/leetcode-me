package main

import (
    "fmt"

    stack "leetcode-me/stack_queue"
)


// test func here.
func test_stack_queue() {
    s := stack.Init()
    s.Push(3)
    min , _ := s.GetMin()
    fmt.Printf("%d\n",min)
    s.Push(2)
    min , _ = s.GetMin()
    fmt.Printf("%d\n",min)
    s.Push(4)
    min , _ = s.GetMin()
    fmt.Printf("%d\n",min)
    s.Push(1)
    min, _ = s.GetMin()
    fmt.Printf("%d\n",min)
}

func main() {
    test_stack_queue()
}
