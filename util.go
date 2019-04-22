package pentimento

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func IsTerminal() bool {
	return terminal.IsTerminal(int(os.Stdout.Fd()))
}
