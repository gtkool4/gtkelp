package buildhelp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gtkool4/grun"
	"github.com/gtkool4/gtkelp/buildhelp"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

//
//-----------------------------------------------------------------[ TESTS ]--

func TestAll(t *testing.T) {
	for _, call := range tests {
		var realCall func()
		gapp.Run(func(gapp *grun.App) gtk.Widgetter {
			b := buildhelp.NewFromString(uiBasic)
			w, _ := NewCustomWidget(gapp.Win, b)
			realCall = func() { call(t, b); gapp.Win.Close() }
			return w
		}, func() { realCall() })
	}

}

var tests = map[string]func(t *testing.T, b *buildhelp.BuildHelp){
	"NotFound": func(t *testing.T, b *buildhelp.BuildHelp) {
		b.ApplicationWindow("missing")
		testErr(t, b.Errors(), "builder found no object for key missing (ApplicationWindow)")
	},
	"BadType": func(t *testing.T, b *buildhelp.BuildHelp) {
		q := b.Adjustment("quit")
		if q != nil {
			t.Error("widget quit should NOT be an Adjustment")
		}
		testErr(t, b.Errors(), "builder has bad type for key quit not a Adjustment")
	},
	"MessageEmpty": func(t *testing.T, b *buildhelp.BuildHelp) {
		testExpectedString(t, "", b.Errors().Error())
	},
	"MessageBadType": func(t *testing.T, b *buildhelp.BuildHelp) {
		b.Window("mainbox")
		expected := buildhelp.FmtErrBadType
		expected = strings.Replace(expected, "%s", "mainbox", 1)
		expected = strings.Replace(expected, "%s", "Window", 1)
		testExpectedString(t, expected, b.Errors().Error())
	},
	"MessageNotFound": func(t *testing.T, b *buildhelp.BuildHelp) {
		b.Dialog("fail")
		expected := buildhelp.FmtErrNotFound
		expected = strings.Replace(expected, "%s", "fail", 1)
		expected = strings.Replace(expected, "%s", "Dialog", 1)
		testExpectedString(t, expected, b.Errors().Error())
	},
	"MessageMultiple": func(t *testing.T, b *buildhelp.BuildHelp) {
		b.TreeStore("fail")
		b.TreeView("mainbox")
		expected := fmt.Sprintf(
			buildhelp.FmtErrNotFound+"\n"+buildhelp.FmtErrBadType,
			"fail", "TreeStore", "mainbox", "TreeView",
		)
		testExpectedString(t, expected, b.Errors().Error())
	},
}

//
//-----------------------------------------------------------------[ HELPERS ]--

func testExpectedString(t *testing.T, expected, have string) {
	if expected != have {
		t.Errorf("builder error: text should be \"%s\" but we have \"%s\"", expected, have)
	}
}

func testErr(t *testing.T, errs grun.Errors, expected string) {
	if len(errs) != 1 {
		t.Errorf("builder error: should have 1 error but we found %d", len(errs))
	}
	testExpectedString(t, expected, errs[0].Error())
}
