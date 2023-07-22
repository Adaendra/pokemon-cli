package main

import (
	"github.com/Adaendra/pokemon-cli/pkg/commands/charttype"
	"github.com/alecthomas/kong"
)

var CLI struct {
	ChartType struct {
		Attack  bool `help:"Give the chart type entries for an attack with the defined type" short:"a" default:"true"'`
		Defense bool `help:"Give the chart type result for a pokemon with the defines types" short:"d" default:"false"'`

		Types []string `arg:"" name:"types" help:"Types to check" type:"string"`
	} `cmd:"" help:"Give the chart type values for the type (or double type) given in parameter."`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "chart-type <types>":
		charttype.Run(CLI.ChartType.Types, CLI.ChartType.Attack, CLI.ChartType.Defense)
	}
}
