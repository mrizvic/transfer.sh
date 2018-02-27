package main

import "github.com/mrizvic/transfer.sh/cmd"

func main() {
	app := cmd.New()
	app.RunAndExitOnError()
}
