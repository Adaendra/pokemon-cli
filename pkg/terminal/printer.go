package terminal

import (
	"fmt"
	"os"
)

func Println(message string) {
	fmt.Fprintln(os.Stderr, message)
}

func ErrorPrintln(message string) {
	Println(fmt.Sprintf("\033[31m%s\033[0m", message))
}
