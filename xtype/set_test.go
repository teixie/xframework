package xtype

import "testing"

func TestSet(t *testing.T) {
	intCollection := NewIntCollection([]int{1, 2, 3, 1}...)
	if got := intCollection.Size(); got != 4 {
		t.Errorf("intCollection.Size() = %v, want %v", got, 4)
	}

	intSet := NewIntSet([]int{1, 2, 1, 2}...)
	if got := intSet.Join(","); got != "1,2" {
		t.Errorf("intSet.Join(%s) = %v, want %v", ",", got, "1,2")
	}
}
