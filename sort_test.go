package channel_sorter

import (
	"reflect"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	z := time.Now()
	a := make(chan int, 10)
	c := <-Sort[int](time.Second*2, a, func(i int, j int) bool { return i < j })
	n := time.Now().Sub(z)
	if n < time.Second*2 {
		t.Fatal("Inactive channel should finish after 2 seconds")
	}
	if n > time.Second*3 {
		t.Fatal("Inactive channel should finish within 3 seconds")
	}
	if len(c) != 0 {
		t.Fatal("Output length should be 0 as no numbers were provided")
	}
}

func TestSortWithDelay(t *testing.T) {
	z := time.Now()
	a := make(chan int, 10)
	go func() {
		time.AfterFunc(time.Second/2, func() { a <- 2 })
		time.AfterFunc(time.Second, func() { a <- 1 })
	}()
	c := <-Sort[int](time.Second*2, a, func(i int, j int) bool { return i < j })
	n := time.Now().Sub(z)
	if n < time.Second*3 {
		t.Fatal("Inactive channel should finish after 3 seconds")
	}
	if n > time.Second*4 {
		t.Fatal("Inactive channel should finish within 4 seconds")
	}
	if len(c) != 2 {
		t.Fatal("Output length should be 2 as 2 numbers were provided")
	}
	if !reflect.DeepEqual(c, []int{1, 2}) {
		t.Fatal("Output slice should match [1, 2]")
	}
}
