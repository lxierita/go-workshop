package lazys_test

import (
	. "lazys"
	"testing"
)

func TestLazyScheduler_should_extract_name_from_function_argument(t *testing.T) {
	want := "foo"

	if got := GetName(foo); want != got {
		t.Errorf("GetName(foo) = %s; want \"foo\"", got)
	}
}

func foo(...int) int {
	return 0
}
