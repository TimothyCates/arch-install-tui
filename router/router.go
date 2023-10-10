package router
import(
  tea "github.com/charmbracelet/bubbletea"
)

type RouterModel struct{
  activeController int
}

func New() RouterModel{
  return RouterModel{}
}

func (r RouterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd){
  return r, nil
}

func (r RouterModel) View() string{
  return "Router"
}

func (r RouterModel) Init() tea.Cmd{
  return nil
}
