package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer A") // called at last even there is PANIC
	}()

	defer func() {
		fmt.Println("defer B") // called secondly in PANIC

		e := recover()
		if (e != nil) {
			fmt.Println(e)
		}
	}()

	defer func() {
		fmt.Println("defer C") // called first in PANIC
	}()

	panic("PANIC!!")
	fmt.Println("Good Joon") // this code is never executed even though recover() is called within defer function
}
