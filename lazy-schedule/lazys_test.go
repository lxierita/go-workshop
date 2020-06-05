package lazys_test

import (
	. "lazys"
	"testing"
)

func TestLazyScheduler_should_schedule_functions_lazily(t *testing.T) {
	ls := Scheduler{}

	f := func(...int) int {
		t.Fatalf("Scheduler is not lazy")
		return 0
	}
	ls.Add(f, []int{1, 2, 3})
	ls.Add(f, []int{4, 5, 6})

}
