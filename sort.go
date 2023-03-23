package channel_sorter

import (
	"github.com/MrMelon54/inactive"
	"sort"
	"time"
)

func Sort[T any](d time.Duration, in chan T, less func(T, T) bool) chan []T {
	z := make(chan []T, 1)
	a := make([]T, 0)
	t := inactive.NewTimer(d)
	go func() {
		for {
			select {
			case <-t.C:
				sort.Slice(a, func(i, j int) bool {
					return less(a[i], a[j])
				})
				z <- a
				a = make([]T, 0)
			case b := <-in:
				t.Tick()
				a = append(a, b)
			}
		}
	}()
	return z
}
