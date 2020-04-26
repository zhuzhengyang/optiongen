package example

import (
	"log"
)

//go:generate optionGen --option_with_struct_name=false

func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"Sounds": string("Meow"),
		"Food":   (*string)(nil),
		"Walk": func() {
			log.Println("Walking")
		},
		"TestNilFunc":(func())(nil),
	}
}
