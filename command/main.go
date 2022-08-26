package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	status int
	err    error
}

const url = "https://charm.sh/"

func checkServer() tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)

	if err != nil {
		return errMsg{err}
	}
	go func() {
		time.Sleep(2 * time.Second)
		
	}
	return statusMsg(res.StatusCode)
}

type statusMsg int

type errMsg struct {
	err error
}

func (e errMsg) Error() string { return e.err.Error() }

func (m model) Init() tea.Cmd {
	return checkServer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		m.status = int(msg)
		return m, tea.Quit
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err)
	}

	s := fmt.Sprintf("Checking %s ...", url)

	if m.status > 0 {
		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
	}

	return "\n" + s + "\n\n"
}

func main() {
	if err := tea.NewProgram(model{}).Start(); err != nil {
		fmt.Printf("error!: %v\n", err)
		os.Exit(1)
	}
}
