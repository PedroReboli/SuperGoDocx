package docx

import (
	"github.com/PedroReboli/go-xmldom"
)

//----------Run---------//
type sUnderline struct {
	value string
	color string
}
type ssFonts struct {
	ascii string
	hAnsi string
	cs    string
}

//TextStyle contains the style of the run or paragraph
type TextStyle struct {
	Document        *xmldom.Node
	bold            bool
	italic          bool
	caps            bool
	color           string
	dstrike         bool
	emboss          bool
	imprint         bool
	outline         bool
	shadow          bool
	smallCaps       bool
	strike          bool
	size            int
	vanish          bool
	vertAlign       string
	underline       bool
	structUnderline sUnderline
	font            ssFonts
}
type run struct {
	Document *xmldom.Node
	Style    TextStyle
}

//--------Paragraph-------//
type pStyle struct {
	Document *xmldom.Node
}
type paragraphStyle struct {
	Document  *xmldom.Node
	StyleName pStyle
	Style     TextStyle
}
type paragraph struct {
	Document *xmldom.Node
	Style    *paragraphStyle
	Runs     []*run
}

//----------Table---------//
type columnStyle struct {
	Document *xmldom.Node
}
type tableColumn struct {
	Document   *xmldom.Node
	Style      columnStyle
	Paragraphs []*paragraph
}
type tableRow struct {
	Document *xmldom.Node
	Cols     []*tableColumn
}
type tableStyle struct {
	Document *xmldom.Node
}
type tableGrid struct {
	Document *xmldom.Node
}
type table struct {
	Document *xmldom.Node
	Style    tableStyle
	Grid     tableGrid
	Rows     []*tableRow
}

//--------Document--------//
type body struct {
	Document  *xmldom.Node
	Paragraph []*paragraph
	Tables    []*table
}
type wDocument struct {
	Document *xmldom.Node
	Body     body
}
