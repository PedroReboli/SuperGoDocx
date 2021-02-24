package docx

import (
	"github.com/PedroReboli/go-xmldom"
)

/*
func (b *body) CreateAddParagraph() {
	b.Paragraph = append(b.Paragraph, &paragraph{})

}
func (b *body) CreateInsertParagraph(index int) {
	b.Paragraph = append(b.Paragraph, paragraph{})
	copy(b.Paragraph[index+1:], b.Paragraph[index:])
	b.Paragraph[index] = paragraph{}
}

func (b *body) AddParagraph(Para paragraph) {
	b.Paragraph = append(b.Paragraph, Para)

}
func (b *body) InsertParagraph(index int, Para paragraph) {
	b.Paragraph = append(b.Paragraph, paragraph{})
	copy(b.Paragraph[index+1:], b.Paragraph[index:])
	b.Paragraph[index] = Para
}
*/
func (b *body) RemoveParagraph(p *paragraph) {

	for Index, para := range b.Paragraph {
		if para == p {
			b.Document.RemoveChild(p.Document)
			b.Paragraph = append(b.Paragraph[:Index], b.Paragraph[Index+1:]...)
			return
		}
	}
}
func (b *body) CreateAddParagraph() *paragraph {
	n := b.Document.CreateNode("p")
	n.NS = &xmldom.Namespace{Name: "w"}
	p := paragraph{Document: n}
	p.createStyle()
	b.Paragraph = append(b.Paragraph, &p)
	return &p
}
func (b *body) InsertCreateParagraph(index int) *paragraph {
	n := b.Document.CreateNodeAt(index, "p")
	n.NS = &xmldom.Namespace{Name: "w"}
	p := paragraph{Document: n}
	p.createStyle()
	b.Paragraph = append(b.Paragraph, &paragraph{})
	copy(b.Paragraph[index+1:], b.Paragraph[index:])
	b.Paragraph[index] = &p
	return &p
}
func (b *body) IndexOf(p *paragraph) int {
	for Index, para := range b.Paragraph {
		if para == p {
			return Index
		}
	}
	return -1
}
func (b *body) CreateParagraphAfter(p *paragraph) *paragraph {
	I := b.IndexOf(p)
	return b.InsertCreateParagraph(I + 1)
}
func (b *body) Search(str string) (*paragraph, bool) {
	for _, p := range b.Paragraph {
		if index(p.GetText(), str) != -1 {
			return p, true
		}
	}
	return nil, false

}

func (b *body) ReplaceAll(old string, new string) {
	for _, p := range b.Paragraph {
		for p.Replace(old, new) == true {

		}
	}
	for _, t := range b.Tables {
		t.ReplaceAll(old, new)
	}
}
