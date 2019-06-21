package main

import (
	log "github.com/shevilangle/log4go"
)

func main() {
	if err := log.SetupLogWithConf("log.json"); err != nil {
		panic(err)
	}
	defer log.Close()

	var name = "skoo"
	index := 0
	log.Debug("log4go %d by %s", index, name)
	log.Info("log4go %d by %s", index, name)
	log.Warn("log4go %d by %s", index, name)
	log.Error("log4go %d by %s", index, name)
	log.Fatal("log4go %d by %s", index, name)
}
