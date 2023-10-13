package controllers
import(
  "github.com/timothycates/arch-install-tui/internal/models/menu"
  "github.com/timothycates/arch-install-tui/internal/views"
)
type mainMenu struct{
  Update func()
  View func() string
}



var model = menu.New([]string{"start", "exit"}, views.CenteredMenu)
var MainMenu mainMenu = mainMenu{

  Update: func(){
    
  },
  View: func() string{
   return model.View()
  },
}
