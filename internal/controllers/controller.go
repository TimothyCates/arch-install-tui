package controllers

import tea "github.com/charmbracelet/bubbletea"

type Controller struct{
  Update func(tea.Msg) (tea.Model, tea.Cmd)
  View func() string
}
