package main

import (
	"github.com/Adaendra/pokemon-cli/pkg/commands/charttype"
	"github.com/alecthomas/kong"
)

var CLI struct {
	ChartType struct {
		Types []string `arg:"" name:"types" help:"Types to check" type:"string"`
	} `cmd:"" help:"Give the chart type values for the type (or double type) given in parameter."`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "chart-type <types>":
		charttype.Run(CLI.ChartType.Types)
	}
}
