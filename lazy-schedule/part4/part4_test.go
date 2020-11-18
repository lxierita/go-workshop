package part4_test

import (
	. "part4"
	"reflect"
	"testing"
	"time"
)

func TestLazyScheduler_should_return_results_of_scheduled_functions_in_order(t *testing.T) {

	ls := Scheduler{}

	want := []int{6, 2}

	f1 := func(args ...int) int {
		r := 0
		for _, n := range args {
			r += n
		}
		time.Sleep(time.Second * 2)
		return r
	}
	f2 := func(args ...int) int {
		r := 1
		for _, n := range args {
			r *= n
		}
		time.Sleep(time.Second * 2)

		return r
	}
	ls.Add(f1, []int{1, 2, 3})
	ls.Add(f2, []int{1, 2})

	if got := ls.Run(); reflect.DeepEqual(want, got) == false {
		t.Errorf("Scheduler.Run() = %d; want [6, 2]", got)
	}

}

func benchmarkRun(b *testing.B, scheduleNum int) {
	ls := Scheduler{}

	f := func(args ...int) int {
		r := 0
		for _, n := range args {
			r += n
		}
		return r
	}

	// lazily add the designated amount of functions
	for i := 0; i < scheduleNum; i++ {
		ls.Add(f, []int{1, 2, 3})
	}

	//reset timer after expensive set up
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ls.Run()
	}

}

func BenchmarkLazyScheduler_run1000000(b *testing.B) {
	benchmarkRun(b, 1000000)
}
