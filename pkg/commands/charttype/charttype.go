package charttype

import (
	"fmt"
	"github.com/Adaendra/pokemon-cli/pkg/pokemon"
	"github.com/Adaendra/pokemon-cli/pkg/terminal"
	"os"
)

func Run(types []string, isAttack bool, isDefense bool) {
	if !isAttack && !isDefense {
		isAttack = true
	}

	if len(types) == 0 {
		terminal.ErrorPrintln("No types selected")
		os.Exit(0)
	}
	if len(types) > 2 && isDefense || len(types) > 1 && isAttack {
		terminal.ErrorPrintln("Too much types selected")
		os.Exit(0)
	}

	for _, t := range types {
		if !pokemon.IsExistingType(t) {
			terminal.ErrorPrintln(fmt.Sprintf("Incorrect type : %s", t))
		}
	}

	// Attack
	if isAttack {
		terminal.Println("ATTACK")
	}

	// Defense
	if isDefense {
		terminal.Println("DEFENSE")
	}
}
