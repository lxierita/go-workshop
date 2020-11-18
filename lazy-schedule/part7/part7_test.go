package part7_test

import (
	. "part7"
	"reflect"
	"testing"
)

func TestLazyScheduler_cannot_add_and_run_functions_at_the_same_time(t *testing.T) {

	ls := NewScheduler()

	f := func(args ...int) int {
		r := 0
		for _, n := range args {
			r += n
		}

		return r
	}
	ls.Add(f, []int{1, 2, 3})

	want := []int{6}
	got := make(chan []int)
	wait := make(chan bool)

	go func() {
		wait <- true
		result, _ := ls.Run()
		got <- result
	}()

	<-wait
	go func() {
		for i := 0; i < 5; i++ {
			ls.Add(f, []int{1, 2, 3})
		}
	}()

	result := <-got
	if reflect.DeepEqual(result, want) == false {
		t.Errorf("Scheduler.Run() = %d;  want [6]", result)
	}
}
