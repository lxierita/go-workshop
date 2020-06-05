package lazys_test

import (
	. "lazys"
	"testing"
)

func TestLazyScheduler_should_add_function(t *testing.T) {
	want := 1

	ls := Scheduler{}
	f := func(...int) int {
		return 0
	}

	ls.Add(f, []int{1, 2, 3})
	ls.Len = 1
	if got := ls.Len; want != got {
		t.Errorf("ls.Len = %d; want 1", got)
	}
}
