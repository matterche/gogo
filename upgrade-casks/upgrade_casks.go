// stupid bulk upgarde of all installed brew cask packages
package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	out, err := exec.Command("brew", "cask", "list").Output()

	if err != nil {
		log.Fatal(err)
	}

	casks := strings.Split(string(out), "\n")

	for _, cask := range casks {
		log.Printf("\n Upgrading cask %v", cask)
		cmd := exec.Command("brew", "cask", "upgrade", cask)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()

		if err != nil {
			log.Print(err)
		} else {
			log.Printf("\n Succesfully upgraded cask %v", cask)
		}

	}

	os.Exit(0)
}
