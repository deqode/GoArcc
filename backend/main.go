package  main

import "alfred/cmd"

func main() {
	if err := cmd.RunServer() ; err != nil {
		panic(err)
	}
}
