package main

import (
	"os"

	myLogger "github.com/aarjan/gopher-meetup-2019/examples/log"
)

func main() {
	log := myLogger.NewLogger(myLogger.SetOutput(os.Stdout), myLogger.SetPrefix("gopher-meetup"))
	log.Info("Hey All!!")
}
