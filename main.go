package main

import (
    "fmt"

    "stack_queue"
)


// write test func here.
func test_stack_queue() {
    s := stack_queue.mystack.Init()
    s.Push(3)
    min := s.GetMin()
    fmt.Printf("%d\n",min)
    s.Push(2)
    min := s.GetMin()
    fmt.Printf("%d\n",min)
    s.Push(4)
    min := s.GetMin()
    fmt.Printf("%d\n",min)
    s.Push(1)
    min := s.GetMin()
    fmt.Printf("%d\n",min)
}

func main() {
    test_stack_queue()
}
