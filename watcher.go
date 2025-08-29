package main

import (
	"os"

	_ "github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Watcher struct{}

func (watcher *Watcher) init() {
	gtkWatcherApp := gtk.NewApplication(app.id+"_watcher", gio.ApplicationDefaultFlags)
	gtkWatcherApp.ConnectActivate(func() { watcher.activate(gtkWatcherApp) })
	gtkWatcherApp.ConnectShutdown(func() {
		database.vacuum()
		database.close()
	})
	if code := gtkWatcherApp.Run(nil); code > 0 {
		os.Exit(code)
	}
}

func (watcher *Watcher) activate(gtkWatcherApp *gtk.Application) {
	database.vacuum()
	clipboard.updateRecentContentFromDatabase()
	clipboard.watch()
	gtkWatcherApp.Hold()
}
