package part10_test

import (
	. "part10"
	"reflect"
	"testing"
)

func TestLazyScheduler_can_add_and_run_interface_type_function(t *testing.T) {

	ls := NewScheduler()
	f1 := func() interface{} {
		return 0
	}
	ls.Add(f1)

	want := []Job{{0, true, "success"}}

	if got, _ := ls.Run(); reflect.DeepEqual(got, want) == false {
		t.Errorf("Scheduler.Run() = %v; want {0, true, \"success\"}", got)
	}
}
