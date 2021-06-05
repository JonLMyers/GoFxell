package client

import (
	"fmt"
	"os"
	"strings"
)

func Executor(cmd string) {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return
	} else if cmd == "quit" || cmd == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

}
