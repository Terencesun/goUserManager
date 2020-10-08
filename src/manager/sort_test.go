package mgr

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T)  {
	var arr = []int{1, 5, 2, 7, 3, 4, 3}
	Sort(&arr)
	fmt.Println(arr)
}

func TestSlice(t *testing.T)  {
	a := make([]int, 0)
	for i:=0; i<10; i++ {
		a = append(a, i)
	}
	for v:=range a{
		fmt.Println(v)
	}
}
