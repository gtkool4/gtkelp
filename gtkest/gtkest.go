// Package gtkest defines a widget maker to run gtk tests.
package gtkest

import (
	"fmt"
	"testing"
	"time"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"

	"github.com/gtkool4/grun"
	"github.com/gtkool4/gtkelp/gtknew"
)

// Maker defines a widget maker to run gtk tests.
type Maker struct {
	newW    func() gtk.Widgetter
	app     *grun.App
	baseID  string
	counter int
}

// New creates a widget maker to run gtk tests.
func New(app *grun.App, newW func() gtk.Widgetter) *Maker {
	baseID := app.ID
	if baseID == "" {
		baseID = "com.github.gotk4.gtkest.default"
	}
	return &Maker{app: app, newW: newW, baseID: baseID}
}

// Run launches the gtk application to run tests on the widget.
func (m *Maker) Run(t *testing.T, userCalls ...func(*testing.T, gtk.Widgetter)) {
	var w gtk.Widgetter // stores the widget to provide for all calls.
	calls := make([]interface{}, len(userCalls)+1)
	calls[0] = func() gtk.Widgetter {
		w = m.newW()
		return w
	}
	for i, call := range userCalls {
		calls[i+1] = func() { call(t, w) }
	}
	m.app.ID = fmt.Sprintf("%s%d", m.baseID, m.counter)
	m.counter++
	m.app.Run(calls...)
}

//
//--------------------------------------------------------------------[ EXIT ]--

// Exit creates an Action that closes the application.
func (m *Maker) Exit(exitCode int) func(*testing.T, gtk.Widgetter) {
	fmt.Println("exit")
	return func(*testing.T, gtk.Widgetter) { gtknew.Idle(func() { m.app.Exit(exitCode) }) }
}

// ExitAfter creates an Action that closes the application after duration.
func (m *Maker) ExitAfter(d time.Duration, exitCode int) func() {
	return func() { go time.AfterFunc(d, func() { m.app.Exit(exitCode) }) }
}
