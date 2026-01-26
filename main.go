package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// flyweight.Main()

	u := User{
		Name: "John",
		Age:  30,
	}

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	fmt.Println("Name: ", t.Name(), " || Kind: ", t.Kind(), " || NumField: ", t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		tag := field.Tag.Get("json")
		fmt.Println("Field: ", field.Name)
		fmt.Println("Tag: ", tag)
		fmt.Println("Value: ", value.Interface())
		fmt.Println("--------------------------------")
	}
}
