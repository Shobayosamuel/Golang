package main

import (
	"fmt"
	"strings"
)

func changeName(val *string) {
	*val = strings.ToUpper(*val)
}

func testi() {
	// var  name string
	// var namePointer *string
	// name = "Samuel"
	// namePointer = &name
	// var defName = *namePointer
	// fmt.Println("Name: ", name)
	// fmt.Println("Name *: ", namePointer)
	// fmt.Println("Name: ", defName)
	name := "Elvis"
	changeName(&name)
	fmt.Println(name)
}