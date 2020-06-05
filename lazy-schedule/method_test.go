package lazys_test

import (
	"fmt"
	. "lazys"
	"testing"
)

// func init() {
// 	Scheduled = Scheduled[:0]
// }

func TestLazyScheduler_should_execute_functions_in_order(t *testing.T) {
	// setup
	ls := Scheduler{}
	Scheduled = Scheduled[:0]

	want := []int{6, 2}

	f1 := func(...int) int {
		return 6
	}
	f2 := func(...int) int {
		return 2
	}
	ls.Add(f1, []int{1, 2, 3})
	ls.Add(f2, []int{1, 2})

	fmt.Println(len(Scheduled))

	ls.Run()
	got := ls.Result

	if diff := compare(got, want); diff != 0 {
		t.Errorf("Scheduler.run() = %d; want [6, 2]", got)
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
