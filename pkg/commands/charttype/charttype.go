package charttype

import (
	"fmt"
	"github.com/Adaendra/pokemon-cli/pkg/pokemon"
	"github.com/Adaendra/pokemon-cli/pkg/terminal"
	"os"
)

func Run(types []string, isAttack bool, isDefense bool) {
	if len(types) == 0 {
		terminal.Println("No types selected")
		os.Exit(0)
	}
	if len(types) > 2 && isDefense || len(types) > 1 && isAttack {
		terminal.Println("Too much types selected")
		os.Exit(0)
	}

	for _, t := range types {
		if !pokemon.IsExistingType(t) {
			terminal.ErrorPrintln(fmt.Sprintf("Incorrect type : %s", t))
		}
	}

}
