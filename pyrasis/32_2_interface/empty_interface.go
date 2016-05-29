package main

import (
	"fmt"
	"strconv"
)

func f1 (arg interface{}) { // interface{} means ANY TYPE
	fmt.Println(arg)
}

type Any interface{}

func f2 (arg Any) {
	fmt.Println("Any : ", arg)
}

type Person struct {
	name string
}

func (p Person) GetName () string {
	return p.name
}

func formatString(arg interface{}) string {
	switch arg.(type) { // check it arg's type. .(type) only can be used in switch's condition
	case int:
		i := arg.(int) // get `arg` value as int (called TYPE ASSERTION)
		return strconv.Itoa(i)
	case float32:
		f := arg.(float32) // get `arg` value as float32
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
	case float64:
		f := arg.(float64) // get `arg` value as float64
		return strconv.FormatFloat(f, 'f', -1, 64)
	case string:
		s := arg.(string) // get `arg` value as string
		return s
	case Person:
		return arg.(Person).GetName()
	case *Person:
		return "[Pointer]" + arg.(*Person).GetName()
	default:
		return "Error"
	}
}

func main() {
	f1(100)
	var str string = "ABC"
	f2(str)

	person := Person{"Good Joon"}
	person2 := new(Person) // declare and assign pointer to Person object
	person2.name = "Bad Joon"
	var t interface{}
	t = Person{"Fair Joon"}

	fmt.Println(formatString(1))
	fmt.Println(formatString(2.555))
	fmt.Println(formatString("This is Joon"))
	fmt.Println(formatString(person))
	fmt.Println(formatString(&person))
	fmt.Println(formatString(person2))

	if v, ok := t.(Person); ok {
		fmt.Println(v, ok)
	}
}
