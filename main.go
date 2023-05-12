package main

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	t.AppendRows([]table.Row{
		{1, "Darth", "Vader", 100000},
		{2, "Leia", "Organa", 200000},
		{3, "Luke", "Skywalker", 300000},
	})
	t.Render()
}
