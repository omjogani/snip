package decoder

import (
	"github.com/fatih/color"
	"github.com/omjogani/snip/controllers"
)

// Enums to hold command values
const (
	Get    string = "get"
	Add    string = "add"
	Remove string = "remove"
	Update string = "update"
	Append string = "Append"
	Help   string = "help"
)

/*
* Command Decoder: Decodes commands and call action accordingly.
* Take use of Above Enum and call respective function from helpers.
 */
func CommandDecoder(cmd string) {
	if cmd == Get {
		controllers.GetSnip()
	} else if cmd == Add {
		controllers.AddSnip()
	} else if cmd == Remove {
		controllers.RemoveSnip()
	} else if cmd == Update {
		controllers.UpdateSnip()
	} else if cmd == Append {
		controllers.AppendSnip()
	} else if cmd == Help {
		// TODO: Write HELP Command and About
	} else {
		color.Red("Invalid Command (type 'help' for help)")
	}
}
