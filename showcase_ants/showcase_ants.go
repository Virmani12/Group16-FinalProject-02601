package main

import (
	"flag"

	//"github.com/icza/gowut/_examples/showcase/showcasecore"
	// double check the functions and make sure they are package showcasecore not package main
	"Group16-FinalProject-02601/showcase_ants/showcasecore_ants"
)

var (
	addr     = flag.String("addr", "", "address to start the server on")
	appName  = flag.String("appName", "showcase", "Gowut app name")
	autoOpen = flag.Bool("autoOpen", true, "auto-open the demo in default browser")
)

func main() {
	flag.Parse()

	showcasecore_ants.StartServer(*appName, *addr, *autoOpen)
}
