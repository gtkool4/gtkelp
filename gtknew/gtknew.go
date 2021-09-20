// Package gtknew creates gtk widgets easier.
package gtknew

import (
	"io"
	"sync"

	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

//
//--------------------------------------------------------------[ CONTAINERS ]--

// HBox creates an horizontal *gtk.Box with optional widgets.
func HBox(spacing int, list ...gtk.Widgetter) *gtk.Box {
	box := gtk.NewBox(gtk.OrientationHorizontal, spacing)
	BoxAppend(box, list...)
	return box
}

// VBox creates a vertical *gtk.Box with optional widgets.
func VBox(spacing int, list ...gtk.Widgetter) *gtk.Box {
	box := gtk.NewBox(gtk.OrientationVertical, spacing)
	BoxAppend(box, list...)
	return box
}

// Appender defines a widget with append method.
type Appender interface {
	gtk.Widgetter
	Append(gtk.Widgetter)
}

// BoxAppend appends widgets to a box, or other "Appender".
func BoxAppend(box Appender, list ...gtk.Widgetter) gtk.Widgetter {
	for _, w := range list {
		box.Append(w)
	}
	return box
}

// CenterBox creates a *gtk.CenterBox with its child widgets.
func CenterBox(start, center, end gtk.Widgetter) *gtk.CenterBox {
	centerBox := gtk.NewCenterBox()
	if start != nil {
		centerBox.SetStartWidget(start)
	}
	if center != nil {
		centerBox.SetCenterWidget(center)
	}
	if end != nil {
		centerBox.SetEndWidget(end)
	}
	return centerBox
}

// HPaned creates an horizontal *gtk.Paned with child widgets.
func HPaned(left, right gtk.Widgetter) *gtk.Paned {
	box := gtk.NewPaned(gtk.OrientationHorizontal)
	if left != nil {
		box.SetStartChild(left)
	}
	if right != nil {
		box.SetEndChild(right)
	}
	return box
}

// VPaned creates a vertical *gtk.Paned with child widgets.
func VPaned(left, right gtk.Widgetter) *gtk.Paned {
	box := gtk.NewPaned(gtk.OrientationVertical)
	if left != nil {
		box.SetStartChild(left)
	}
	if right != nil {
		box.SetEndChild(right)
	}
	return box
}

// Expander creates a *gtk.Expander with an optional child widget.
func Expander(name string, w ...gtk.Widgetter) gtk.Widgetter {
	exp := gtk.NewExpander(name)
	if len(w) > 0 && w[0] != nil {
		exp.SetChild(w[0])
	}
	return exp
}

// Frame creates a *gtk.Frame with an optional child widget.
func Frame(title string, w ...gtk.Widgetter) gtk.Widgetter {
	fram := gtk.NewFrame(title)
	if len(w) > 0 && w[0] != nil {
		fram.SetChild(w[0])
	}
	return fram
}

// ScrolledWindow packs a widget in a *gtk.ScrolledWindow.
func ScrolledWindow(child gtk.Widgetter) *gtk.ScrolledWindow {
	sw := gtk.NewScrolledWindow()
	if child != nil {
		sw.SetChild(child)
	}
	return sw
}

//
//------------------------------------------------------------------[ OTHERS ]--

// HSep creates an horizontal *gtk.Separator.
func HSep() *gtk.Separator {
	return gtk.NewSeparator(gtk.OrientationHorizontal)
}

// VSep creates a vertical *gtk.Separator.
func VSep() *gtk.Separator {
	return gtk.NewSeparator(gtk.OrientationVertical)
}

// LabelWithMarkup creates a *gtk.Label with markup text.
func LabelWithMarkup(str string) *gtk.Label {
	label := gtk.NewLabel("")
	label.SetMarkup(str)
	return label
}

//
//------------------------------------------------------------------[ PIXBUF ]--

// PixbufReader creates a *gdkpixbuf.Pixbuf from a reader.
func PixbufReader(reader io.Reader) (*gdkpixbuf.Pixbuf, error) {
	load := gdkpixbuf.NewPixbufLoader()
	load.Connect("size-prepared", func(m *gdkpixbuf.PixbufLoader, w, h int) {
		m.SetSize(48, 48)
	})
	io.Copy(writer{load}, reader)
	pix := load.Pixbuf()
	e := load.Close()
	return pix, e
}

type writer struct{ loader *gdkpixbuf.PixbufLoader }

func (w writer) Write(data []byte) (n int, err error) {
	w.loader.Write(data)
	return len(data), nil
}

// PixbufReaderError creates a *gdkpixbuf.Pixbuf from a failable reader source
// like http.Get.
func PixbufReaderError(reader io.Reader, e error) (*gdkpixbuf.Pixbuf, error) {
	if e != nil {
		return nil, e
	}
	return PixbufReader(reader)
}

//
//------------------------------------------------------------[ IDLE ACTIONS ]--

// Idle adds a function to call on the next gtk idle cycle, to safely use the
// GTK backend with our goroutines.
//
// This version has a call buffer, useful when you're adding a lot of idle calls.
// Idle is stacking callbacks on the go side and flush them all at once when
// called by the C side, to reduce roundtrips between go and C.
//
func Idle(calls ...func()) {
	if len(calls) == 0 {
		return
	}
	idleMu.Lock()
	idleStack = append(idleStack, calls...)
	if !idleRun {
		idleRun = true
		externglib.IdleAdd(callIdle)
	}
	idleMu.Unlock()
}

var idleMu = &sync.Mutex{} // Protects idleStack and idleRun.
var idleStack []func()     // List of functions to run in the glib main loop.
var idleRun bool           // Tells if the idle flusher is running or not.

// callIdle flushes the idleStack list by calling them all.
// It will also process calls received while running, and stops only when
// idleStack is really empty.
//
func callIdle() {
	// fmt.Println("--gtk loop", len(idleStack))
	var toRun []func()
	for idleStack != nil {
		idleMu.Lock()
		toRun, idleStack = idleStack, nil
		idleMu.Unlock()

		for _, call := range toRun {
			call()
		}
	}

	idleMu.Lock()
	idleRun = false
	idleMu.Unlock()
}
