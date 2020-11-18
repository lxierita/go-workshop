package part8_test

import (
	. "part8"
	"reflect"
	"testing"
)

func TestLazyScheduler_can_handle_panics_and_recover(t *testing.T) {

	ls := NewScheduler()

	f := func(args ...int) int {
		panic("panic")
	}

	f1 := func(args ...int) int {
		return 0
	}

	ls.Add(f, []int{1, 2, 3})
	ls.Add(f1, []int{1, 2, 3})

	want := []Job{{0, false, ""}, {0, true, "success"}}

	if got, _ := ls.Run(); reflect.DeepEqual(got, want) == false {
		t.Errorf("Scheduler.Run() = %v; want {{0, false, \"\"}, {6, true, \"success\"}}", got)
	}
}
