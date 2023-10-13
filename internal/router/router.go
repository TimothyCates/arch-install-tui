package router
import(
  ctrls "github.com/timothycates/arch-install-tui/internal/controllers"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/timothycates/arch-install-tui/internal/models/installOptions"
)

type RouterModel struct{
  activeController int
  options installOptions.Options
  controllers *OrderedMap
}

func (r RouterModel)  setController(controller string){
  ctrlerIndex, error := r.controllers.GetIndex(controller)
  if error != nil{
    return
  }
  r.activeController = ctrlerIndex
}

func New() RouterModel{
  var r RouterModel = RouterModel{
    controllers: NewOrderedMap(),
  }
  r.controllers.Add("MainMenu", ctrls.MainMenu)
  return r
}

func (r RouterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd){
  switch msg := msg.(type){
  case tea.KeyMsg:
    switch msg.String(){
    case "ctrl+c":
      return r, tea.Quit
    }
  }
  _, ctrlValue, error := r.controllers.GetByIndex(r.activeController)
  if error != nil{
    return r, nil
  }

  controller, ok := ctrlValue.(ctrls.Controller)
  if !ok{
    return r, nil
  }

  controller.Update(msg)

  return r, nil
}

func (r RouterModel) View() string{
  _, ctrlValue, error := r.controllers.GetByIndex(r.activeController)
  if error != nil {
    return "activeView set to impossible state"
  }
  controller, ok := ctrlValue.(ctrls.Controller)
  if !ok{
    return ""
  }

  return controller.View()
}

func (r RouterModel) Init() tea.Cmd{
  return nil
}
