package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	err error
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "e":
			c := exec.Command(os.Getenv("EDITOR")) //nolint:gosec
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			return m, tea.Exec(c)

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.ExecCompletedMsg:
		m.err = msg.Err
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}
	return "Press e to open your EDITOR. Press q to quit."
}

func main() {
	if err := tea.NewProgram(model{}).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
