package menu

import tea "github.com/charmbracelet/bubbletea"

type Model struct{
  Choices []string
  Cursor int
  Selected map[int]struct{}
  view func(m tea.Model) string
}

func New(options []string, view func(m tea.Model) string) Model{
  return  Model{
    Choices: options,
    view: view,
  }
}

func (m Model) Init() tea.Cmd{
  return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
  return m, nil
}

func (m Model) View() string{
  return m.view(m)
}

