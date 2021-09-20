// Package buildhelp helps loading interfaces from gtk.Builder.
package buildhelp

import (
	"fmt"

	"github.com/gtkool4/grun"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// Errors formating.
var (
	FmtErrBadType  = "builder has bad type for key %s not a %s"
	FmtErrNotFound = "builder found no object for key %s (%s)"
)

//
//------------------------------------------------------------[ BUILD HELPER ]--

// BuildHelp is a small wrapper around gtk.Builder to load interfaces easily.
type BuildHelp struct {
	gtk.Builder
	errors grun.Errors
}

// New creates a *BuildHelp to load gtk.Builder interfaces easily.
func New() *BuildHelp {
	return &BuildHelp{Builder: *gtk.NewBuilder()}
}

// NewFromString creates a *BuildHelp to load gtk.Builder interfaces easily from a string.
func NewFromString(str string) *BuildHelp {
	return &BuildHelp{Builder: *gtk.NewBuilderFromString(str, -1)}
}

// NewFromFile creates a *BuildHelp to load gtk.Builder interfaces easily from a file.
func NewFromFile(file string) *BuildHelp {
	return &BuildHelp{Builder: *gtk.NewBuilderFromFile(file)}
}

// Errors returns builder errors: bad types or not found.
func (b *BuildHelp) Errors() grun.Errors { return b.errors }

func (b *BuildHelp) getO(typ, name string, call func(interface{}) bool) {
	obj := b.GetObject(name)
	if obj == nil {
		b.errors.Append(fmt.Errorf(FmtErrNotFound, name, typ))
		return
	}

	if !call(obj.Cast()) {
		b.errors.Append(fmt.Errorf(FmtErrBadType, name, typ))
	}
}

//
//-----------------------------------------------------------------[ WIDGETS ]--

// Adjustment get the named object as Adjustment.
func (b *BuildHelp) Adjustment(name string) (w *gtk.Adjustment) {
	b.getO("Adjustment", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.Adjustment); return })
	return w
}

// ApplicationWindow get the named object as ApplicationWindow.
func (b *BuildHelp) ApplicationWindow(name string) (w *gtk.ApplicationWindow) {
	b.getO("ApplicationWindow", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.ApplicationWindow); return })
	return w
}

// Box get the named object as Box.
func (b *BuildHelp) Box(name string) (w *gtk.Box) {
	b.getO("Box", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.Box); return })
	return w
}

// Button get the named object as Button.
func (b *BuildHelp) Button(name string) (w *gtk.Button) {
	b.getO("Button", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.Button); return })
	return w
}

// CellRendererText get the named object as CellRendererText.
func (b *BuildHelp) CellRendererText(name string) (w *gtk.CellRendererText) {
	b.getO("CellRendererText", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.CellRendererText); return })
	return w
}

// CheckButton get the named object as CheckButton.
func (b *BuildHelp) CheckButton(name string) (w *gtk.CheckButton) {
	b.getO("CheckButton", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.CheckButton); return })
	return w
}

// ComboBox get the named object as ComboBox.
func (b *BuildHelp) ComboBox(name string) (w *gtk.ComboBox) {
	b.getO("ComboBox", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.ComboBox); return })
	return w
}

// Dialog get the named object as Dialog.
func (b *BuildHelp) Dialog(name string) (w *gtk.Dialog) {
	b.getO("Dialog", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.Dialog); return })
	return w
}

// Frame get the named object as Frame.
func (b *BuildHelp) Frame(name string) (w *gtk.Frame) {
	b.getO("Frame", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.Frame); return })
	return w
}

// HeaderBar get the named object as HeaderBar.
func (b *BuildHelp) HeaderBar(name string) (w *gtk.HeaderBar) {
	b.getO("HeaderBar", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.HeaderBar); return })
	return w
}

// IconView get the named object as IconView.
func (b *BuildHelp) IconView(name string) (w *gtk.IconView) {
	b.getO("IconView", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.IconView); return })
	return w
}

// Image get the named object as Image.
func (b *BuildHelp) Image(name string) (w *gtk.Image) {
	b.getO("Image", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.Image); return })
	return w
}

// Label get the named object as Label.
func (b *BuildHelp) Label(name string) (w *gtk.Label) {
	b.getO("Label", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.Label); return })
	return w
}

// ListStore get the named object as ListStore.
func (b *BuildHelp) ListStore(name string) (w *gtk.ListStore) {
	b.getO("ListStore", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.ListStore); return })
	return w
}

// ScrolledWindow get the named object as ScrolledWindow.
func (b *BuildHelp) ScrolledWindow(name string) (w *gtk.ScrolledWindow) {
	b.getO("ScrolledWindow", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.ScrolledWindow); return })
	return w
}

// Scale get the named object as Scale.
func (b *BuildHelp) Scale(name string) (w *gtk.Scale) {
	b.getO("Scale", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.Scale); return })
	return w
}

// ShortcutsWindow get the named object as ShortcutsWindow.
func (b *BuildHelp) ShortcutsWindow(name string) (w *gtk.ShortcutsWindow) {
	b.getO("ShortcutsWindow", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.ShortcutsWindow); return })
	return w
}

// Stack get the named object as Stack.
func (b *BuildHelp) Stack(name string) (w *gtk.Stack) {
	b.getO("Stack", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.Stack); return })
	return w
}

// StackSwitcher get the named object as StackSwitcher.
func (b *BuildHelp) StackSwitcher(name string) (w *gtk.StackSwitcher) {
	b.getO("StackSwitcher", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.StackSwitcher); return })
	return w
}

// Switch get the named object as Switch.
func (b *BuildHelp) Switch(name string) (w *gtk.Switch) {
	b.getO("Switch", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.Switch); return })
	return w
}

// TextView get the named object as TextView.
func (b *BuildHelp) TextView(name string) (w *gtk.TextView) {
	b.getO("TextView", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.TextView); return })
	return w
}

// ToggleButton get the named object as ToggleButton.
func (b *BuildHelp) ToggleButton(name string) (w *gtk.ToggleButton) {
	b.getO("ToggleButton", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.ToggleButton); return })
	return w
}

// TreeModelFilter get the named object as TreeModelFilter.
func (b *BuildHelp) TreeModelFilter(name string) (w *gtk.TreeModelFilter) {
	b.getO("TreeModelFilter", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.TreeModelFilter); return })
	return w
}

// TreeModelSort get the named object as TreeModelSort.
func (b *BuildHelp) TreeModelSort(name string) (w *gtk.TreeModelSort) {
	b.getO("TreeModelSort", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.TreeModelSort); return })
	return w
}

// TreeSelection get the named object as TreeSelection.
func (b *BuildHelp) TreeSelection(name string) (w *gtk.TreeSelection) {
	b.getO("TreeSelection", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.TreeSelection); return })
	return w
}

// TreeStore get the named object as TreeStore.
func (b *BuildHelp) TreeStore(name string) (w *gtk.TreeStore) {
	b.getO("TreeStore", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.TreeStore); return })
	return w
}

// TreeView get the named object as TreeView.
func (b *BuildHelp) TreeView(name string) (w *gtk.TreeView) {
	b.getO("TreeView", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.TreeView); return })
	return w
}

// TreeViewColumn get the named object as TreeViewColumn.
func (b *BuildHelp) TreeViewColumn(name string) (w *gtk.TreeViewColumn) {
	b.getO("TreeViewColumn", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.TreeViewColumn); return })
	return w
}

// Window get the named object as Window.
func (b *BuildHelp) Window(name string) (w *gtk.Window) {
	b.getO("Window", name, func(o interface{}) (ok bool) { w, ok = o.(*gtk.Window); return })
	return w
}
