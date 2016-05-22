package main

import (
	"fmt"
	"os"

	"github.com/goodjoon/lang-test-go/pyrasis/utils"
)

func main() {
	utils.GetPwd()
	file, err := os.Open("pyrasis/26_defer_file/hello.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return // file.Close() called
	}

	buf := make([]byte, 4096)
	if _, err := file.Read(buf); err != nil {
		fmt.Println(err)
		return // file.Close() called
	}

	fmt.Println(string(buf))
}
