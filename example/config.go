package main

import (
	"log"
)

//go:generate optionGen

func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"Sounds": string("Meow"),
		"Food":   (*string)(nil),
		"Walk": func() {
			log.Println("Walking")
		},
	}
}
