package main

import "github.com/JoaoPedr0Maciel/charm/cmd"

var (
	version = "dev"
)

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
