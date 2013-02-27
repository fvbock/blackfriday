//
// Blackfriday Markdown Processor
// Available at http://github.com/russross/blackfriday
//
// Copyright © 2011 Russ Ross <russ@russross.com>.
// Distributed under the Simplified BSD License.
// See README.md for details.
//

// Rendering backend to strip markdown from a text
//
// This is a quick hack that needs some cleaning - but works for what
// i need now...
package blackfriday

import (
	"bytes"
	"fmt"
)

// Stripped is a type that implements the Renderer interface for output
// stripped of markdown.
//
// Do not create this directly, instead use the StrippedRenderer function.
type Stripped struct {
}

// StrippedRenderer creates and configures a Stripped object, which
// satisfies the Renderer interface.
//
// flags is a set of STRIPPED_* options ORed together (currently no such
// options are defined).
func StrippedRenderer(flags int) Renderer {
	fmt.Println("!!!!!!!!!!!!!")
	return &Stripped{}
}

// render code chunks using verbatim, or listings if we have a language
func (options *Stripped) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	out.Write(text)
}

func (options *Stripped) BlockQuote(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (options *Stripped) BlockHtml(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (options *Stripped) Header(out *bytes.Buffer, text func() bool, level int) {
	marker := out.Len()
	if !text() {
		out.Truncate(marker)
		return
	}
	out.WriteString(" ")
}

func (options *Stripped) HRule(out *bytes.Buffer) {
	out.WriteString(" ")
}

func (options *Stripped) List(out *bytes.Buffer, text func() bool, flags int) {
	marker := out.Len()
	if !text() {
		out.Truncate(marker)
		return
	}
	out.WriteString(" ")
}

func (options *Stripped) ListItem(out *bytes.Buffer, text []byte, flags int) {
	out.Write([]byte("\n"))
	out.Write(text)
}

func (options *Stripped) Paragraph(out *bytes.Buffer, text func() bool) {
	marker := out.Len()
	out.WriteString("\n")
	if !text() {
		out.Truncate(marker)
		return
	}
}

func (options *Stripped) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	// TODO: this still outputs table delimiters | ...
	out.Write(header)
	out.Write([]byte("\n"))
	out.Write(body)
	out.Write([]byte("\n"))
}

func (options *Stripped) TableRow(out *bytes.Buffer, text []byte) {
	if out.Len() > 0 {
		out.WriteString(" \n")
	}
	out.Write(text)
}

func (options *Stripped) TableCell(out *bytes.Buffer, text []byte, align int) {
	out.WriteString(" ")
	out.Write(text)
}

func (options *Stripped) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	// out.WriteString("\\href{")
	// if kind == LINK_TYPE_EMAIL {
	// 	out.WriteString("mailto:")
	// }
	out.Write([]byte(" "))
	out.Write(link)
	out.Write([]byte(" "))
	// out.WriteString("}{")
	// out.Write(link)
	// out.WriteString("}")
}

func (options *Stripped) CodeSpan(out *bytes.Buffer, text []byte) {
	out.WriteString(" ")
	escapeSpecialCharsStripped(out, text)
	out.WriteString(" ")
}

func (options *Stripped) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	out.WriteString(" ")
	out.Write(text)
	out.WriteString("}")
	out.WriteString(" ")
}

func (options *Stripped) Emphasis(out *bytes.Buffer, text []byte) {
	out.WriteString("\\textit{")
	out.Write(text)
	out.WriteString("}")
}

func (options *Stripped) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	// if bytes.HasPrefix(link, []byte("http://")) || bytes.HasPrefix(link, []byte("https://")) {
	// 	// treat it like a link
	// 	out.WriteString("\\href{")
	// 	out.Write(link)
	// 	out.WriteString("}{")
	// 	out.Write(alt)
	// 	out.WriteString("}")
	// } else {
	// 	out.WriteString("\\includegraphics{")
	// 	out.Write(link)
	// 	out.WriteString("}")
	// }
	out.WriteString(" ")
}

func (options *Stripped) LineBreak(out *bytes.Buffer) {
	out.WriteString(" \n")
}

func (options *Stripped) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	// out.WriteString("\\href{")
	out.WriteString(" ")
	out.Write(link)
	out.WriteString(" ")
	// out.WriteString("}{")
	// out.Write(content)
	// out.WriteString("}")
}

func (options *Stripped) RawHtmlTag(out *bytes.Buffer, tag []byte) {
}

func (options *Stripped) TripleEmphasis(out *bytes.Buffer, text []byte) {
	out.WriteString(" ")
	out.Write(text)
	out.WriteString(" ")
}

func (options *Stripped) StrikeThrough(out *bytes.Buffer, text []byte) {
	out.WriteString(" ")
	out.Write(text)
	out.WriteString(" ")
}


func needsBackslashStripped(c byte) bool {
	for _, r := range []byte("_{}%$&\\~") {
		if c == r {
			return true
		}
	}
	return false
}


func escapeSpecialCharsStripped(out *bytes.Buffer, text []byte) {
	for i := 0; i < len(text); i++ {
		// directly copy normal characters
		org := i

		for i < len(text) && !needsBackslash(text[i]) {
			i++
		}
		if i > org {
			out.Write(text[org:i])
		}

		// escape a character
		if i >= len(text) {
			break
		}
		out.WriteByte('\\')
		out.WriteByte(text[i])
	}
}

func (options *Stripped) Entity(out *bytes.Buffer, entity []byte) {
	// TODO: convert this into a unicode character or something
	out.Write(entity)
}

func (options *Stripped) NormalText(out *bytes.Buffer, text []byte) {
	escapeSpecialChars(out, text)
}

// header and footer
func (options *Stripped) DocumentHeader(out *bytes.Buffer) {
	out.WriteString(" ")
}

func (options *Stripped) DocumentFooter(out *bytes.Buffer) {
	out.WriteString("\n")
}
