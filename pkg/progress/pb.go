package progress

import (
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

func NewProgress() *tea.Program {
	p1 := progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))
	p2 := progress.New(progress.WithScaledGradient("#DB9831", "#FDFF8C"))
	p3 := progress.New(progress.WithScaledGradient("#FF83E9", "#FDFF8C"))

	p1.Width = 80
	p2.Width = 80
	p3.Width = 80
	prog := make(map[string]progress.Model)
	prog["assessments"] = p1
	prog["symbols"] = p2
	prog["deletes"] = p3

	p := tea.NewProgram(model{progress: prog})
	return p

}

type model struct {
	percent  float64
	progress map[string]progress.Model
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)
	s := ""
	for k, v := range m.progress {
		s += "\n" + pad + k + pad + v.ViewAs(m.percent) + "\n"
	}

	s += "\n\n" + pad + helpStyle("press any key to quit")
	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {

	case tea.KeyMsg:
		return m, tea.Quit

	// case tea.WindowSizeMsg:
	// 	m.progress["symbols"].Width = msg.Width - padding*2 - 4
	// 	if m.progress["symbols"].Width > maxWidth {
	// 		m.progress["symbols"].Width = maxWidth
	// 	}
	// 	return m, nil

	case int:
		m.percent += .01
		return m, nil

	default:
		return m, nil
	}
}

// func returnData(ch chan (int)) {
// 	for i := 0; i < 50; i++ {
// 		go func(i int) {

// 			r := rand.Intn(10)
// 			time.Sleep(time.Duration(r) * time.Second)
// 			p.Send(i)
// 		}(i)

// 	}
// }
