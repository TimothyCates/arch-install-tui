package controllers

import tea "github.com/charmbracelet/bubbletea"

type Controller interface{
  Update(tea.Msg) (tea.Model, tea.Cmd)
  View() string
}

