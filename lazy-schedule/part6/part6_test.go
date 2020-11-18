package part6_test

import (
	. "part6"
	"reflect"
	"testing"
	"time"
)

func TestLazyScheduler_times_out_and_return_unaffected_data_structure(t *testing.T) {

	ls := NewScheduler()

	f := func(args ...int) int {
		r := 0
		for _, n := range args {
			r += n
		}
		return r
	}

	ftimeout := func(args ...int) int {
		r := 0
		for _, n := range args {
			r += n
		}
		time.Sleep(time.Second * 10)
		return r
	}

	ls.Add(f, []int{1, 2, 3})
	ls.Add(ftimeout, []int{1, 2, 3})
	ls.Add(f, []int{1, 2, 3})

	want := []Job{{6, true, "success"}, {0, false, "timed out"}, {6, true, "success"}}

	if _, err := ls.Run(); err != nil {
		t.Error("Scheduler.Run() should have 1 timeout error")
	}

	if got, _ := ls.Run(); reflect.DeepEqual(got, want) == false {
		t.Errorf("Scheduler.Run() = %v should return unaffected work, want %v", got, want)
	}
}
