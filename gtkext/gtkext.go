// Package gtkext formats strings for Pango / gtk.
//
package gtkext

import "strings"

// Escape escapes a string to use as gtk text. Symbols affected: &<>.
//
func Escape(msg string) string {
	for _, chr := range []struct{ from, to string }{
		{"&", "&amp;"}, // Escape ampersand MUST be first.
		{"<", "&lt;"},
		{">", "&gt;"},
	} {
		msg = strings.Replace(msg, chr.from, chr.to, -1)
	}
	return msg
}

//
//--------------------------------------------------------------------[ TAGS ]--

// Big formats the text with the big size.
func Big(text string) string { return tag("big", text) }

// Small formats the text with the small size.
func Small(text string) string { return tag("small", text) }

// Bold formats the text with the bold tag.
func Bold(text string) string { return tag("b", text) }

// Mono formats the text with the monospace font.
func Mono(text string) string { return tag("tt", text) }

// Italic formats the text with the italic tag.
func Italic(text string) string { return tag("i", text) }

// Strike formats the text with the strike tag.
func Strike(text string) string { return tag("s", text) }

// Sub formats the text with the Subscript tag (underline).
func Sub(text string) string { return tag("sub", text) }

// Sup formats the text with the Superscript tag (overline).
func Sup(text string) string { return tag("sup", text) }

// Underline formats the text with the underline tag.
func Underline(text string) string { return tag("u", text) }

func tag(tagl, text string) string { return "<" + tagl + ">" + text + "</" + tagl + ">" }

//
//------------------------------------------------------------------[ OTHERS ]--

// URI formats a link with its text.
//
func URI(uri, text string) string { return "<a href=\"" + uri + "\">" + text + "</a>" }

// List converts a list of lines to a bullet list (prepend " *" to lines )
func List(list ...string) string { sep := " * "; return sep + strings.Join(list, "\n"+sep) }

//
//--------------------------------------------------------------------[ SPAN ]--

// Span is the most general markup tag. See SpanAttribute for its defined attributes.
func Span(text string, attrs map[SpanAttribute]string) string {
	out := []string{}
	for attr, text := range attrs {
		out = append(out, string(attr)+"=\""+text+"\"")
	}
	return tagA("span", strings.Join(out, " "), text)
	// return "<span " + strings.Join(out, " ") + ">" + text + "</span>"
}

func tagA(t, attrs, text string) string { return "<" + t + " " + attrs + ">" + text + "</" + t + ">" }

// SpanAttribute defines text attributes usable in Pango / gtk
type SpanAttribute string

// Span attributes list.
const (
	AttrFontDesc      SpanAttribute = "font_desc"     // A font description string, such as "Sans Italic 12"; note that any other span attributes will override this description. So if you have "Sans Italic" and also a style="normal" attribute, you will get Sans normal, not italic.
	AttrFontFamily                  = "font_family"   // A font family name such as "normal", "sans", "serif" or "monospace".
	AttrFace                        = "face"          // A synonym for font_family
	AttrSize                        = "size"          // The font size in thousandths of a point, or one of the absolute sizes 'xx-small', 'x-small', 'small', 'medium', 'large', 'x-large', 'xx-large', or one of the relative sizes 'smaller' or 'larger'.
	AttrStyle                       = "style"         // The slant style - one of 'normal', 'oblique', or 'italic'
	AttrWeight                      = "weight"        // The font weight - one of 'ultralight', 'light', 'normal', 'bold', 'ultrabold', 'heavy', or a numeric weight.
	AttrVariant                     = "variant"       // The font variant - either 'normal' or 'smallcaps'.
	AttrStretch                     = "stretch"       // The font width - one of 'ultracondensed', 'extracondensed', 'condensed', 'semicondensed', 'normal', 'semiexpanded', 'expanded', 'extraexpanded', 'ultraexpanded'.
	AttrForeground                  = "foreground"    // An RGB color specification such as '#00FF00' or a color name such as 'red'.
	AttrBackground                  = "background"    // An RGB color specification such as '#00FF00' or a color name such as 'red'.
	AttrUnderline                   = "underline"     // The underline style - one of 'single', 'double', 'low', or 'none'.
	AttrRise                        = "rise"          // The vertical displacement from the baseline, in ten thousandths of an em. Can be negative for subscript, positive for superscript.
	AttrStrikethrough               = "strikethrough" // 'true' or 'false' whether to strike through the text.
	AttrFallback                    = "fallback"      // If True enable fallback to other fonts of characters are missing from the current font. If disabled, then characters will only be used from the closest matching font on the system. No fallback will be done to other fonts on the system that might contain the characters in the text. Fallback is enabled by default. Most applications should not disable fallback.
	AttrLang                        = "lang"          // A language code, indicating the text language.
)
