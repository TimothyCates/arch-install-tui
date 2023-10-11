package router
import(
  tea "github.com/charmbracelet/bubbletea"
  "github.com/timothycates/arch-install-tui/internal/models/installOptions"
)

type RouterModel struct{
  activeController int
  options installOptions.Options
}

func New() RouterModel{
  return RouterModel{}
}

func (r RouterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd){
  switch msg := msg.(type){
  case tea.KeyMsg:
    switch msg.String(){
    case "ctrl+c":
      return r, tea.Quit
    }
  }
  return r, nil
}

func (r RouterModel) View() string{
  return "Router"
}

func (r RouterModel) Init() tea.Cmd{
  return nil
}
