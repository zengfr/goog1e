package cmd

import (
	"log"
	"os/exec"
)

func main() {
	cmdCat := exec.Command("cat", "main.go")
	cmdWC := exec.Command("wc", "-l")
	data, err := pipeCommands(cmdCat, cmdWC)
	if err != nil {
		log.Fatalf("failed to call pipeCommands(): %v", err)
	}
	log.Printf("output: %s", data)
}

func pipeCommands(commands ...*exec.Cmd) ([]byte, error) {
	for i, command := range commands[:len(commands)-1] {
		out, err := command.StdoutPipe()
		if err != nil {
			return nil, err
		}
		command.Start()
		commands[i+1].Stdin = out
	}
	final, err := commands[len(commands)-1].Output()
	if err != nil {
		return nil, err
	}
	return final, nil
}
