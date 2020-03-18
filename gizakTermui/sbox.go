package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	table := widgets.NewTable()
	table.SetRect(10, 0, 53, 10)
	table.Rows = [][]string{
		[]string{"", "", ""},
		[]string{"Col1", "", ""},
		[]string{"", "Col2", ""},
		[]string{"", "", ""},
		[]string{"", "", ""},
		[]string{"", "", "    Col3"},
		[]string{"", "", ""},
		[]string{"", "", ""},
	}

	table.TextStyle = ui.NewStyle(ui.ColorCyan)
	table.RowSeparator = false
	table.BorderStyle = ui.NewStyle(ui.ColorBlack)

	table.FillRow = true

	for i := 0; i < 3; i++ {
		table.RowStyles[i] = ui.NewStyle(ui.ColorYellow, ui.ColorBlack)
	}
	for i := 3; i < 8; i++ {
		table.RowStyles[i] = ui.NewStyle(ui.ColorRed, ui.ColorGreen)
	}

	ui.Render(table)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
