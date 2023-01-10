package main

import (
	"appdaemon"
	"time"
)

func main() {
	new(appdaemon.DualProcessGuard).SetInterval(7*time.Second).Start(appdaemon.Start)
}
