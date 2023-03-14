package vespera

import (
  "context"
  "fmt"
  "log"

	"github.com/darcys22/vespera/vespera/gamepad"
)


// State struct
type State struct {
	ctx context.Context
  light bool
  xboxready bool
  xbox *gamepad.GamePad
}

// NewState creates a new State application struct
func NewState() *State {
	return &State{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (s *State) Startup(ctx context.Context) {
  s.xboxready = false
	s.ctx = ctx
  xbox, err := gamepad.NewGamepad(0)
  if err != nil {
    log.Println(err)
  } else {
    s.xbox = xbox
    s.xboxready = true
  }
}

func (s *State) Shutdown(ctx context.Context) {}

// Toggles the light on the mainboard
func (s *State) ToggleLight() string {
  s.light = !s.light
  return fmt.Sprintf("Toggled value of flag: %v", s.light)
}

// Prints the xbox controller status
func (s *State) ControllerStatus() string {

  if !s.xboxready {
    return "Xbox controller not ready"
  } 

  state, err := s.xbox.State()
  if err != nil {
    log.Fatal(err)
  }
  return fmt.Sprintf("State: %#v", state)

}

