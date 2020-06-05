package clazys_test

import (
	. "lazys"
	"testing"
)

// func init() {
// 	Scheduled = Scheduled[:0]
// }

func TestCLazyScheduler_should_execute_functions_in_order(t *testing.T) {
	// setup
	ls := Scheduler{}
	Scheduled = Scheduled[:0]

	want := []int{6, 2}

	f1 := func(args ...int) int {
		r := 0
		for _, n := range args {
			r += n
		}
		return r
	}
	f2 := func(args ...int) int {
		r := 1
		for _, n := range args {
			r *= n
		}
		return r
	}
	ls.Add(f1, []int{1, 2, 3})
	ls.Add(f2, []int{1, 2})

	got := ls.Run()

	if diff := compare(got, want); diff != 0 {
		t.Errorf("Scheduler.Run() = %d; want [6, 2]", got)
	}

}

func compare(got, want []int) int {

	m := make(map[int]int)

	for _, x := range got {
		m[x]++
	}

	var diff []int
	for _, y := range want {
		if m[y] > 0 {
			m[y]--
			continue
		}
		diff = append(diff, y)
	}

	return len(diff)
}

func BenchmarkCLazyScheduler(b *testing.B) {
	ls := Scheduler{}
	Scheduled = Scheduled[:0]

	f1 := func(args ...int) int {
		r := 0
		for _, n := range args {
			r += n
		}
		return r
	}
	f2 := func(args ...int) int {
		r := 1
		for _, n := range args {
			r *= n
		}
		return r
	}
	ls.Add(f1, []int{1, 2, 3})
	ls.Add(f2, []int{1, 2})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ls.Run()
	}

}