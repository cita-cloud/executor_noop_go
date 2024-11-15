package main

import (
	"fmt"
	"github.com/cita-cloud/executor_noop_go/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var app = cmd.NewApp("the controller mico-server binary")

func init() {
	app.Name = "executor"
	app.Commands = []*cli.Command{
		runCommand,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}
