package main

import cmd "solana/cmd/commands"

func main() {
	root := cmd.NewRootCmd()

	root.Execute()
}
