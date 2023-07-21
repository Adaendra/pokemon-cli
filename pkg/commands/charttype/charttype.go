package charttype

import (
	"fmt"
	"github.com/Adaendra/pokemon-cli/pkg/pokemon"
	"github.com/Adaendra/pokemon-cli/pkg/terminal"
	"os"
)

func Run(types []string) {
	if len(types) == 0 {
		terminal.Println("No types selected")
		os.Exit(0)
	}
	if len(types) > 2 {
		terminal.Println("Too much types selected")
		os.Exit(0)
	}

	for _, t := range types {
		if !pokemon.CheckType(t) {
			terminal.Println(fmt.Sprintf("Incorrect type : %s", t))
		}
	}
}
