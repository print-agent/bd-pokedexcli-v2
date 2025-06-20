package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func REPL() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("😼 ▶️ ")

		if !scanner.Scan() {
			break
		}

		cmdText := cleanInput(scanner.Text())

		// TODO: checker if text matched defined commands

		fmt.Printf("💀 You typed command: %v\n", cmdText[0])
	}
}
