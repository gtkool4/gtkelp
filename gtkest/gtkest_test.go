package gtkest_test

import (
	"testing"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"

	"github.com/gtkool4/grun"
	"github.com/gtkool4/gtkelp/gtkest"
)

// Copy and edit this example to launch chains of tests on your widgets.

// Run a test on a widget with a simple callback.
//
func testLabel(t *testing.T) {
	runT(t, func(t *testing.T, w *gtk.Label) {
		w.SetLabel("updated")
		if w.Label() != "updated" {
			t.Error("label content is not 'updated': ", w.Text())
		}
	}) // Can run more tests on the same widget.
}

//
//-------------------------------------------[ PREPARE CREATOR AND CONVERTER ]--

// Set the creator with your make widget function and either:
//    set the type alias to your widget type.
//      or
//    replace *gtk.Label with your widget type and remove the type alias.
//
type WidgetT = *gtk.Label

var thisW = gtkest.New(gapp, func() gtk.Widgetter { return gtk.NewLabel("label") })

var gapp = grun.NewSized(800, 700,
	grun.SetFmtTitleTest(),
	grun.SetHeadless(), // Comment to display windows.
)

func runT(t *testing.T, call func(*testing.T, WidgetT)) {
	thisW.Run(t,
		func(t *testing.T, w gtk.Widgetter) { call(t, w.(WidgetT)) },
		thisW.Exit(0), // Comment to keep applications alive.
	)
	// t.Fail() // Uncomment to fail all tests (to see output).
}

//--

func Example() {}
