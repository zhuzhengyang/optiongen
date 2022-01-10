package example

import "testing"

func TestNewConfig(t *testing.T) {
	tc := NewFuncNameSpecified(false, "", WithTestMapIntInt(map[int]int{2: 4}))
	if tc == nil {
		t.Fatal("new config error")
	}
	if tc.TestMapIntInt[2] != 4 {
		t.Fatal("map get val error")
	}
	previousValue := tc.TestInt
	changeTo := 1232323232323232
	previous := tc.ApplyOption(WithTestInt(changeTo))
	if tc.TestInt != changeTo {
		t.Fatal("ApplyOption failed")
	}
	tc.ApplyOption(previous...)
	if tc.TestInt != previousValue {
		t.Fatal("ApplyOption Restore failed")
	}
}
