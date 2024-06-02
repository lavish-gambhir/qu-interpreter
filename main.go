package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/lavish-gambhir/qu-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is qu programming lang!\n", user.Username)
	fmt.Printf("Feel free to type in the commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
