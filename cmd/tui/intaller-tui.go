package main

import (
	"fmt"
	"os"
	"github.com/timothycates/arch-install-tui/internal/router"
	tea "github.com/charmbracelet/bubbletea"
)
func main(){
  var router *router.RouterModel = router.New()
  var program *tea.Program = tea.NewProgram(router)
  if _, err := program.Run(); err != nil{
		fmt.Fprintf(os.Stderr, "Error running arch-tui: %v\n", err)
		os.Exit(1)
  }

}
