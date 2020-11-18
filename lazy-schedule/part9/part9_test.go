package part9_test

import (
	. "part9"
	"reflect"
	"testing"
	"time"
)

func TestLazyScheduler_manage_timeouts_with_context(t *testing.T) {

	ls := NewScheduler()

	f := func(args ...int) int {
		time.Sleep(time.Second * 5)
		return 0
	}

	f1 := func(args ...int) int {
		return 0
	}

	ls.Add(f, []int{1, 2, 3})
	ls.Add(f1, []int{1, 2, 3})

	want := []Job{{0, false, "time out"}, {0, true, "success"}}

	if got, _ := ls.Run(); reflect.DeepEqual(got, want) == false {
		t.Errorf("Scheduler.Run() = %v; want {{0, false, \"\"}, {6, true, \"success\"}}", got)
	}
	if _, err := ls.Run(); err == nil {
		t.Errorf("Scheduler.Run() = %v ; want one timeout error", err)
	}

}
