package example

import "testing"

func TestNewConfig(t *testing.T) {
	tc := NewConfig(WithTestMapIntInt(map[int]int{2:4}))
	if tc == nil {
		t.Fatal("new config error")
	}
	if tc.TestMapIntInt[2] !=4 {
		t.Fatal("map get val error")
	}
}