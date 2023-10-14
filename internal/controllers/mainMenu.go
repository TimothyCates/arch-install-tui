package controllers

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/timothycates/arch-install-tui/internal/models/menu"
	"github.com/timothycates/arch-install-tui/internal/views"
)

type MainMenu struct{
  model tea.Model
  ChangeView func(string)
}

func (m MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd){
  switch msg := msg.(type){
  case tea.KeyMsg:
    switch msg.String(){
    }
  }
  return m.model, nil
}

func (m MainMenu) View() string{
  return m.model.View()
}

func NewMainMenu(routerCallback func(string)) (Controller){
  var model = menu.New([]string{"Test1", "Exit"}, views.CenteredMenu)
  var menu MainMenu = MainMenu{
    model: model,
    ChangeView: routerCallback,
  }
  return menu
}
