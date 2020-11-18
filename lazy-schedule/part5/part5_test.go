package part5_test

import (
	. "part5"
	"reflect"
	"runtime"
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

func TestLazyScheduler_should_have_no_goroutine_leaks_after_run(t *testing.T) {

	ls := NewScheduler()

	want := 2

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

	ls.Run()

	if got := runtime.NumGoroutine(); want != got {
		t.Errorf("Scheduler.Run() should have 2 goroutine after run; want 2; got: %d", got)
	}

}
