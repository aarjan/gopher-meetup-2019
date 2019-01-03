package main

import (
	"os"

	myLogger "github.com/aarjan/devfest-2017/examples/log"
)

func main() {
	log := myLogger.NewLogger(myLogger.SetOutput(os.Stdout), myLogger.SetPrefix("gopher-meetup"))
	log.Info("Hey All!!")
}
