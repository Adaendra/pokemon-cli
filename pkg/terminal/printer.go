package terminal

import (
	"fmt"
	"os"
)

func Println(message string) {
	fmt.Fprintln(os.Stderr, message)
}
