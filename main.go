package main

import "github.com/cita-cloud/executor_noop_go/cmd"

var app = cmd.NewApp("the controller mico-server binary")

func main() {
	println("hello world")
}
