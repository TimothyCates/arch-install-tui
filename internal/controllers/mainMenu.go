package controllers

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/timothycates/arch-install-tui/internal/models/menu"
	"github.com/timothycates/arch-install-tui/internal/views"
)

var model = menu.New([]string{"start", "exit"}, views.CenteredMenu)
var MainMenu Controller = Controller{

  Update: func(msg tea.Msg) (tea.Model, tea.Cmd){
    return model, nil 
  },

  View: func() string{
    return model.View()
  },
}
