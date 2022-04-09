package simple

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortInts(t *testing.T) {
	var integers = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	sort.Ints(integers)

	sortedInts := integers
	got := sortedInts
	want := []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("sortStrings() = %v, want %v", got, want)
	}
}