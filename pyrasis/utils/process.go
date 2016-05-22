package utils

import (
	"fmt"
	"os"
)

func GetPwd() string{
	pwd, err := os.Getwd()
	if (err != nil) {
		fmt.Println("Error : ", err)
	}
	fmt.Println(pwd)
	return pwd
}
