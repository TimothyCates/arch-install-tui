package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/timothycates/arch-install-tui/internal/models/menu"
)

func CenteredMenu(m tea.Model) string{
  if custom, ok := m.(menu.Model); ok{
    return custom.Choices[0]
  }
  return ""
}


