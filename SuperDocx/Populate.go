package docx

import (
	"strconv"

	"github.com/PedroReboli/go-xmldom"
)

func (w *wDocument) populate(Root *xmldom.Node) {
	w.Document = Root
	w.Body = body{}
	w.Body.populate(Root.GetChild("body"))
}
func (b *body) populate(Root *xmldom.Node) {
	par := Root.GetChildren("p")
	b.Document = Root
	para := []*xmldom.Node{}
	for _, p := range par {
		para = append(para, p)

	}
	b.Paragraph = make([]*paragraph, len(para))
	if len(para) == 0 {
		//fmt.Println("no Paragraphs")
		return
	}
	for i := 0; i < len(para); i++ {
		b.Paragraph[i] = &paragraph{}
		b.Paragraph[i].populate(para[i])
	}
	tab := Root.GetChildren("tbl")
	b.Tables = make([]*table, len(tab))
	for i := 0; i < len(tab); i++ {
		b.Tables[i] = &table{}
		b.Tables[i].populate(tab[i])
	}
}
func (p *paragraph) populate(Root *xmldom.Node) {
	p.Document = Root
	runs := Root.GetChildren("r")

	if len(runs) == 0 {
		//fmt.Println("No Runs")
		return
	}

	for i := 0; i < len(runs); i++ {
		x := p.CreateRun()
		x.populate(runs[i])
	}

}
func (r *run) populate(Root *xmldom.Node) {
	r.Document = Root
	if len(r.Document.GetChildren("rPr")) != 0 {
		r.Style.populate(r.Document.GetChild("rPr"))
	}
	if len(r.Document.GetChildren("t")) == 0 {
		return

	} else {
		if len(r.Document.GetChild("t").Text) == 0 {
			Root.Parent.RemoveChild(Root)
		}
	}
	x := r.Document.GetChild("t")
	for _, A := range x.Attributes {
		if A.Name == "space" {
			A.NS.Name = "xml"
		}
	}
}
func (t *TextStyle) populate(Root *xmldom.Node) {
	t.Document = Root
	for _, Element := range t.Document.Children {
		switch Name := Element.Name; Name {
		case "b":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.bold = true
		case "i":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.italic = true

		case "caps":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.caps = true
		case "color":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.color = Element.GetAttributeValue("val")
		case "dstrike":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.dstrike = true

		case "emboss":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.emboss = true
		case "imprint":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.imprint = true
		case "outline":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.outline = true
		case "shadow":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.shadow = true
		case "smallCaps":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.smallCaps = true
		case "strike":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.strike = true
		case "sz":
			i, err := strconv.Atoi(Element.GetAttributeValue("val"))
			if err != nil {
				t.Document = t.Document.RemoveChild(Element)
			} else {
				t.size = i
			}
		case "vanish":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.vanish = true
		case "vertAlign":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.vertAlign = Element.GetAttributeValue("val")
		case "u":
			if Element.GetAttributeValue("val") == "0" {
				t.Document.RemoveChild(Element)
				break
			}
			t.underline = true
			t.structUnderline = sUnderline{value: Element.GetAttributeValue("val"), color: Element.GetAttributeValue("color")}
		case "rFonts":
			t.font = ssFonts{
				ascii: Element.GetAttributeValue("ascii"),
				hAnsi: Element.GetAttributeValue("hAnsi"),
				cs:    Element.GetAttributeValue("cs")}
		}
	}
}
func (t *table) populate(Root *xmldom.Node) {
	t.Document = Root
	if len(t.Document.GetChildren("tblPr")) != 0 {
		t.Style.populate(t.Document.GetChild("tblPr"))
	}
	if len(t.Document.GetChildren("tblGrid")) != 0 {
		t.Grid.populate(t.Document.GetChild("tblPr"))
	}
	if len(t.Document.GetChildren("tr")) != 0 {
		rows := t.Document.GetChildren("tr")
		t.Rows = make([]*tableRow, len(rows))
		for i := 0; i < len(rows); i++ {
			t.Rows[i] = &tableRow{}
			t.Rows[i].populate(rows[i])
		}
	}
}
func (t *tableStyle) populate(Root *xmldom.Node) {
	t.Document = Root
}
func (t *tableGrid) populate(Root *xmldom.Node) {
	t.Document = Root
}
func (t *tableRow) populate(Root *xmldom.Node) {
	t.Document = Root
	if len(t.Document.GetChildren("tc")) != 0 {
		cols := t.Document.GetChildren("tc")
		t.Cols = make([]*tableColumn, len(cols))
		for i := 0; i < len(cols); i++ {
			t.Cols[i] = &tableColumn{}
			t.Cols[i].populate(cols[i])
		}
	}
}
func (t *tableColumn) populate(Root *xmldom.Node) {
	t.Document = Root
	if len(t.Document.GetChildren("tcPr")) != 0 {
		t.Style.populate(t.Document.GetChild("tcPr"))
	}
	if len(t.Document.GetChildren("p")) != 0 {
		para := t.Document.GetChildren("p")
		t.Paragraphs = make([]*paragraph, len(para))
		for i := 0; i < len(para); i++ {
			t.Paragraphs[i] = &paragraph{}
			t.Paragraphs[i].populate(para[i])
		}
	}
}
func (c *columnStyle) populate(Root *xmldom.Node) {
	c.Document = Root
}
