package example

import (
	"log"
)

//go:generate optiongen --option_with_struct_name=false --v=true

func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"TestNil":           nil,
		"TestBool":          false,
		"TestInt":           32,
		"TestInt64":         int64(32),
		"TestSliceInt":      []int{1, 2, 3},
		"TestSliceInt64":    []int64{1, 2, 3},
		"TestSliceString":   []string{"test1", "test2"},
		"TestSliceBool":     []bool{false, true},
		"TestSliceIntNil":   []int(nil),
		"TestSliceIntEmpty": []int{},

		"TestMapIntInt":       map[int]int{1: 1, 2: 2, 3: 3},
		"TestMapIntString":    map[int]string{1: "test"},
		"TestMapStringInt":    map[string]int{"test": 1},
		"TestMapStringString": map[string]string{"test": "test"},

		"TestString": "Meow",
		"Food":       (*string)(nil),
		"Walk": func() {
			log.Println("Walking")
		},
		"TestNilFunc":        (func())(nil),
		"TestReserved1_":     []byte(nil),
		"TestReserved2Inner": 1,
	}
}
