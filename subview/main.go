package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// tea.NewProgram()
	m := New()
	err := tea.NewProgram(m).Start()
	if err != nil {
		fmt.Print(err)
	}
}

type state int

const (
	view1 state = iota
	view2
)

// MainModel the main model of the program; holds other models and bubbles
type MainModel struct {
	state state
	v1    tea.Model
	v2    tea.Model
}

// View return the text UI to be output to the terminal
func (m MainModel) View() string {
	switch m.state {
	case view1:
		return m.v1.View()
	default:
		return m.v2.View()
	}
}

// New initialize the main model for your program
func New() MainModel {
	return MainModel{
		state: view1,
		v1:    model1{s: "abc", j: 10},
		v2:    model2{"gef"},
	}
}

// Init run any intial IO on program start
func (m MainModel) Init() tea.Cmd {
	return nil
}

// Update handle IO and commands
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			return m, tea.Quit
		}
		if k == "s" {
			m.state = view1
		}
		if k == "r" {
			m.state = view2
		}
	}

	if m.state == view1 {
		m.v1, cmd = m.v1.Update(msg)
	} else {
		m.v2, cmd = m.v2.Update(msg)
	}
	return m, cmd
}

type model1 struct {
	s string
	j int
}

func (m model1) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model1) Init() tea.Cmd {
	return nil
}

func (m model1) View() string {
	return m.s
}

type model2 struct {
	s string
}

func (m model2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model2) Init() tea.Cmd {
	return nil
}

func (m model2) View() string {
	return m.s
}
