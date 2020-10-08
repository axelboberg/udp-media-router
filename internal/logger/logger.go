package logger

import "fmt"

func Info (realm string, msg string) {
	str := "[" + realm + "] " + msg
	fmt.Println(str)
}