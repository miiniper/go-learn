package main

import (
	"errors"
	"fmt"
	"reflect"
)

func typeof(v interface{}) string {
	fmt.Println("#########This is reflect result########################")
	return reflect.TypeOf(v).String()
}

type item struct {
	Name string
}

func (i item) String() string {
	return fmt.Sprintf("1item name: %v", i.Name)
}

type person struct {
	Name string
	Sex  string
}

func (p person) String() string {
	return fmt.Sprintf("2person name: %v sex: %v", p.Name, p.Sex)
}

func Parse(i interface{}) interface{} {
	fmt.Println(typeof(i))
	switch i.(type) {
	case string:
		return &item{
			Name: i.(string),
		}
	case []string:
		data := i.([]string)
		length := len(data)
		if length == 2 {
			return &person{
				Name: data[0],
				Sex:  data[1],
			}
		} else {
			return nil
		}
	default:
		panic(errors.New("type match miss"))
	}
	return nil
}

func main() {
	fmt.Println("server starting ... ")

}
