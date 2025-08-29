package main

import (
	"os"
	"strings"
)

var (
	app      Application
	gui      GUI
	watcher  Watcher
	database Database
	ipc      IPC
)

func main() {
	app.init()
	osArgs := strings.Join(os.Args, " ")
	switch {
	// No need to use "flag" package. It's overkill.
	case strings.Contains(osArgs, " --watch"):
		watcher.init()
	// Deprecated. Will be removed in future releases.
	case strings.Contains(osArgs, " watch"):
		watcher.init()
	default:
		gui.init()
	}
}
