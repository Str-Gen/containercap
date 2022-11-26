package tui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func initUi() {
	ui_scenarios := widgets.NewList()
	ui_scenarios.Title = "Welcome to containercap!"
	ui_scenarios.TextStyle = ui.NewStyle(ui.ColorYellow)

	ui_scenarios_bar := widgets.NewBarChart()
	ui_scenarios_bar.Title = "Scenarios ?/?"
	ui_scenarios_bar.Labels = []string{"", "Unfinished", "Finished", "Bundled"}
	ui_scenarios_bar.BarColors = []ui.Color{ui.ColorBlack, ui.ColorRed, ui.ColorYellow, ui.ColorGreen}
	ui_scenarios_bar.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
	ui_scenarios_bar.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlack), ui.NewStyle(ui.ColorYellow), ui.NewStyle(ui.ColorYellow), ui.NewStyle(ui.ColorYellow)}
	ui_scenarios_bar.BarWidth = 10

	ui_info := widgets.NewParagraph()
	ui_info.Title = "Info"
	ui_info.WrapText = true

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(ui.NewRow(1.0/2,
		ui.NewCol(1.4/2, ui_scenarios),
		ui.NewCol(0.6/2, ui_scenarios_bar)),
		ui.NewRow(1.0/2, ui_info))

	ui.Render(grid)
}
