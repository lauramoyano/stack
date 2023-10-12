package main

import (
	"fmt"

	"github.com/lauramoyano/stack/pop"
	"github.com/lauramoyano/stack/push"
)

func main() {
	miSlice := []int{1, 2, 3, 4, 5}

	push.Push(&miSlice, 3)
	result, _ := pop.Pop(&miSlice)
	fmt.Println(miSlice)
	fmt.Println(result)

}
