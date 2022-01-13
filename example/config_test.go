package example

import "testing"

func TestNewConfig(t *testing.T) {
	tc := NewFuncNameSpecified(false, "", WithTestMapIntInt(map[int]int{2: 4}))
	if tc == nil {
		t.Fatal("new config error")
	}
	tc.GetFOO()
	if tc.GetTestMapIntInt()[2] != 4 {
		t.Fatal("map get val error")
	}
	previousValue := tc.GetTestInt()
	changeTo := 1232323232323232
	previous := tc.ApplyOption(WithTestInt(changeTo))
	if tc.GetTestInt() != changeTo {
		t.Fatal("ApplyOption failed")
	}
	tc.ApplyOption(previous...)
	if tc.GetTestInt() != previousValue {
		t.Fatal("ApplyOption Restore failed")
	}

	WithXXXXXXRedis(nil)
}
