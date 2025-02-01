package main

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/yiffyi/autope"
)

func main() {
	autope.SetupDefaultLogger()
	fd, err := os.Open("playbook.toml")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	dec := toml.NewDecoder(fd)

	var playbook autope.Playbook
	err = dec.Decode(&playbook)
	if err != nil {
		panic(err)
	}

	err = playbook.PEStage.Run()
	if err != nil {
		panic(err)
	}
}
