package main

import "github.com/showyquasar88/proj-combine/cobra-demo/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
