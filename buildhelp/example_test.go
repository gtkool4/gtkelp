package buildhelp_test

import (
	_ "embed"
	"fmt"
	"time"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/gtkool4/grun"
	"github.com/gtkool4/gtkelp/buildhelp"
	"github.com/gtkool4/gtkelp/gtknew"
)

//go:embed example_test.ui
var uiBasic string

//
//-----------------------------------------------------------[ CUSTOM WIDGET ]--

// CustomWidget shows how to create advanced widgets with the builder.
//
// This is a text editor with cut, copy, paste and quit buttons.
//
//   Acts as a gtk.Widgetter to interact with the gtk world
//   Acts as a io.Stringer to interact with the go world
//
type CustomWidget struct {
	gtk.Box // Main container is first level. Act as (at least) a gtk.Widgetter.
	cut     *gtk.Button
	copy    *gtk.Button
	paste   *gtk.Button
	quit    *gtk.Button
	text    *gtk.TextView
}

func NewCustomWidget(closer interface{ Close() }, b *buildhelp.BuildHelp) (*CustomWidget, grun.Errors) {
	//
	// Get references of widgets.
	// Note that you will have to keep a reference of all objects connected to
	// callbacks to prevent the GC from freeing the associated function.
	// Remember that otherwise it will be seen as unused on the go side.
	//
	w := &CustomWidget{
		Box:   *b.Box("mainbox"),
		cut:   b.Button("cut"),
		copy:  b.Button("copy"),
		paste: b.Button("paste"),
		quit:  b.Button("quit"),
		text:  b.TextView("text"),
	}

	// Get other widgets from the builder.
	menu := b.HeaderBar("menubar")

	// Check builder errors.
	if b.Errors().IsError() {
		return nil, b.Errors()
	}

	// Apply last settings and connect events callbacks.
	menu.SetShowTitleButtons(true)
	w.cut.Connect("activate", w.Cut)
	w.copy.Connect("activate", w.Copy)
	w.paste.Connect("activate", w.Paste)
	w.quit.Connect("activate", func() { fmt.Println("closed"); closer.Close() })
	return w, nil
}

// Widget Public API.

func (w *CustomWidget) Copy()  { w.text.Buffer().CopyClipboard(w.text.Clipboard()) }             // Copy copies the selected text.
func (w *CustomWidget) Cut()   { w.text.Buffer().CutClipboard(w.text.Clipboard(), true) }        // Cut cuts the selected text.
func (w *CustomWidget) Paste() { w.text.Buffer().PasteClipboard(w.text.Clipboard(), nil, true) } // Paste pastes the buffer at cursor position.

// String returns the content of the text buffer.
func (w *CustomWidget) String() string {
	start, end := w.text.Buffer().Bounds()
	return w.text.Buffer().Text(&start, &end, true)
}

//
//----------------------------------------------------[ CREATE APP & BUILDER ]--

var gapp = grun.App{
	ID:     "com.github.gtkool4.gtkelp.buildhelp",
	Title:  "Basic Application",
	Width:  600,
	Height: 300,
}

// How to create advanced Gtk4 interfaces in go with gtk builder.
func Example() {
	fmt.Println("exit code :", gapp.Run(func(app *grun.App) gtk.Widgetter {
		build := buildhelp.NewFromString(uiBasic)
		w, errs := NewCustomWidget(app.Win, build)
		if !errs.IsError() {
			gtknew.Idle(func() { time.Sleep(time.Second / 2); w.testButtons() })
		}
		return errs.Widget(w) // returns either error label or the valid widget.
	}))
	// Output:
	// start
	//
	// Hello GTK4!
	//
	// Hello GTK4!
	// Hello GTK4! GTK4!
	// closed
	// exit code : 0
}

//
//-------------------------------------------------------------------[ TESTS ]--

func (w *CustomWidget) testButtons() {
	testStr := "Hello GTK4!"
	fmt.Println("start") // start
	fmt.Println(w)       // empty
	w.text.Buffer().InsertAtCursor(testStr, -1)
	fmt.Println(w) // testStr
	start, end := w.text.Buffer().Bounds()
	w.text.Buffer().SelectRange(&start, &end)
	w.cut.Activate()
	externglib.IdleAdd(func() { // copy/paste are async calls so we're resyncing after them.
		fmt.Println(w) // empty
		w.paste.Activate()
		externglib.IdleAdd(func() {
			fmt.Println(w) // testStr
			start, end = w.text.Buffer().Bounds()
			start.ForwardChars(5)
			w.text.Buffer().SelectRange(&start, &end)
			w.copy.Activate()
			w.paste.Activate()
			w.paste.Activate()
			externglib.IdleAdd(func() {
				fmt.Println(w)    // Hello GTK4! GTK4!
				w.quit.Activate() // test the last button
			})
		})
	})
}

// func ExampleNewFromFile() {
// 	gapp.Run(func(app *grun.App) (gtk.Widgetter, grun.Errors) {
// 		return NewCustomWidget(app.Win, buildhelp.NewFromFile("test_simple.ui"))
// 	}, grun.ExitAfter(time.Second/2, 0))
// 	// Output:
// }
