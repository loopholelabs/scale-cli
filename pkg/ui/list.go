package ui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/loopholelabs/scale-go/scalefunc"
)

var (
	listStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
)

var _ tea.Model = (*List)(nil)

type List struct {
	table table.Model
}

func (l List) Init() tea.Cmd { return nil }

func (l List) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return l, tea.Quit
}

func (l List) View() string {
	return listStyle.Render(l.table.View()) + "\n"
}

func NewList(entries []*scalefunc.ScaleFunc, middleware bool) error {
	columns := []table.Column{
		{Title: "Name", Width: 10},
		{Title: "Tag", Width: 10},
		{Title: "Language", Width: 10},
		{Title: "Middleware", Width: 10},
	}

	rows := make([]table.Row, 0, len(entries))
	for _, scaleFunc := range entries {
		if middleware && !scaleFunc.ScaleFile.Middleware {
			continue
		}
		row := table.Row{scaleFunc.ScaleFile.Name, "", scaleFunc.ScaleFile.Build.Language, "false"}
		if scaleFunc.ScaleFile.Middleware {
			row[3] = "true"
		}
		if scaleFunc.Tag != "" {
			row[1] = scaleFunc.Tag
		}
		rows = append(rows, row)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight(len(rows)),
	)

	s := table.DefaultStyles()
	s.Selected.Foreground(lipgloss.Color("#FFFFFF"))
	s.Selected.Bold(false)
	s.Cell.Foreground(lipgloss.Color("#FFFFFF"))
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	t.SetStyles(s)

	l := List{t}

	return tea.NewProgram(l).Start()
}
