package gtkext_test

import (
	"fmt"

	"github.com/gtkool4/gtkelp/gtkext"
)

func Example() {
	// Escape your unvalidated input before adding tags, or the display can be broken.
	fmt.Println("Escape:", gtkext.Escape("&<>"))

	// Tags
	fmt.Println(gtkext.Big("Big"))
	fmt.Println(gtkext.Small("Small"))
	fmt.Println(gtkext.Bold("Bold"))
	fmt.Println(gtkext.Mono("Mono"))
	fmt.Println(gtkext.Italic("Italic"))
	fmt.Println(gtkext.Strike("Strike"))
	fmt.Println(gtkext.Sub("Sub"))
	fmt.Println(gtkext.Sup("Sup"))
	fmt.Println(gtkext.Underline("Underline"))

	// Span
	fmt.Println(gtkext.Span("Span", map[gtkext.SpanAttribute]string{
		gtkext.AttrBackground: "red",
	}))

	// Others
	fmt.Println(gtkext.URI("URI", "text"))
	fmt.Println(gtkext.List("List", "item 2"))

	// Output:
	// Escape: &amp;&lt;&gt;
	// <big>Big</big>
	// <small>Small</small>
	// <b>Bold</b>
	// <tt>Mono</tt>
	// <i>Italic</i>
	// <s>Strike</s>
	// <sub>Sub</sub>
	// <sup>Sup</sup>
	// <u>Underline</u>
	// <span background="red">Span</span>
	// <a href="URI">text</a>
	//  * List
	//  * item 2
}
