package goqueue

import "testing"

func TestQueue(t *testing.T) {
	v := New()
	v.Push(1)
	vtst := v.Pop()
	if vtst != 1 {
		t.Error("Should have", 1, "got", vtst)
	}
}
