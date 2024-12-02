package domain

import "testing"

func TestRandom(t *testing.T) {
	t1, t2, t3, t4 := Random()
	if t1 == 0 {
		t.Error("t1 equals 0")
		t.FailNow()
	}
	if t2 == t1 {
		t.Error("t2 equals t1")
		t.FailNow()
	}
	if t3 == t1 || t3 == t2 {
		t.Error("t3 equals t1 or t2")
		t.FailNow()
	}
	if t4 == t1 || t4 == t2 || t4 == t3 {
		t.Error("t4 equals t1 or t2 or t3")
		t.FailNow()
	}
}
