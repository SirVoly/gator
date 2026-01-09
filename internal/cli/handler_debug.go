package cli

import (
	"fmt"
)

func handlerDebug(s *state, cmd command) error {
	fmt.Printf("Debugging\n")
	fmt.Printf("---------\n")
	return nil
}
