package cli

import (
	"fmt"
)

func handlerHelp(s *state, cmd command) error {
	fmt.Printf("Gator is a tool for gathering RSS Feeds.\n\n")
	fmt.Printf("Usage:\n\n")
	fmt.Printf("\tgator <command> [arguments]\n\n")
	fmt.Printf("The commands are:\n\n")
	fmt.Printf("\thelp\t\tshows this informative text\n")
	fmt.Printf("\tlogin\t\tlogin using your username\n")
	fmt.Printf("\tregister\tregister your username\n")
	fmt.Printf("\tusers\t\tview all registered users\n")
	fmt.Printf("\nThe following commands are for development only:\n\n")
	fmt.Printf("\tdebug\t\ta special command to test specific parts\n")
	fmt.Printf("\treset\t\treset the full database.\n")
	return nil
}
