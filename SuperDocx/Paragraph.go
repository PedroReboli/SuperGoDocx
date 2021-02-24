package docx

import (
	"strings"

	"github.com/PedroReboli/go-xmldom"
)

/*func (p *paragraph) CreateAddRun() {
	p.Runs = append(p.Runs, &run{})

}
func (p *paragraph) CreateInsertRun(index int) {
	p.Runs = append(p.Runs, &run{})
	copy(p.Runs[index+1:], p.Runs[index:])
	p.Runs[index] = &run{}
}

//----------CreateMultiple-------//
func (p *paragraph) CreateAddRuns(n int) {
	r := make([]*run, n)
	p.Runs = append(p.Runs, r...)

}
func (p *paragraph) CreateInsertRuns(index int, n int) {
	r := make([]*run, n)
	p.Runs = append(p.Runs, r...)
	copy(p.Runs[index+n:], p.Runs[index:])
	copy(p.Runs[index:index+n], r)
}

//--------------Append-----------//
func (p *paragraph) AddRun(r run) {
	p.Runs = append(p.Runs, &r)

}
func (p *paragraph) InsertRun(index int, r run) {
	p.Runs = append(p.Runs, &run{})
	copy(p.Runs[index+1:], p.Runs[index:])
	p.Runs[index] = &run{}
}

//---------AppendMultiple--------//
func (p *paragraph) AddRuns(r []*run) {
	p.Runs = append(p.Runs, r...)

}
func (p *paragraph) InsertRuns(index int, r []*run) {
	runlenth := len(r)
	p.Runs = append(p.Runs, r...)                 //increase slice size
	copy(p.Runs[index+runlenth:], p.Runs[index:]) //alocate space for runs, copy from (index-end) to index plus lenth of run slice
	copy(p.Runs[index:index+runlenth], r)         //put slice of run in alocated space
}

//--------------Styles------------//
func (p *paragraph) CopyStyle(Para paragraph) {
	p.Style = Para.Style
}
func (p *paragraph) SameSyle(Para paragraph) bool {
	return p.Style == Para.Style
}*/
//--------------Create-----------//
func (p *paragraph) createStyle() {
	n := p.Document.CreateNode("pPr")
	n.NS = &xmldom.Namespace{Name: "w"}
	p.Style = &paragraphStyle{Document: n}
}
func (p *paragraph) CreateRun() *run {
	n := p.Document.CreateNode("r")
	n.NS = &xmldom.Namespace{Name: "w"}
	r := run{Document: n}
	r.setup()
	p.Runs = append(p.Runs, &r)
	return &r
}
func (p *paragraph) CreateRuns(n int) {
	for x := 0; x <= n; x++ {
		n := p.Document.CreateNode("r")
		n.NS = &xmldom.Namespace{Name: "w"}
		r := run{Document: n}
		r.setup()
		p.Runs = append(p.Runs, &r)
	}
}

//-------------Utils-------------//
func (p *paragraph) ClearRuns() {
	for _, r := range p.Document.GetChildren("r") {
		p.Document.RemoveChild(r)
	}
	p.Runs = []*run{}
}
func (p *paragraph) GetText() string {
	text := ""
	for _, r := range p.Runs {
		text += r.GetText()
	}
	return text
}

type character struct {
	char     string
	runStyle TextStyle
}

//Replace search paragraph text and replace old by new and return true if find
func (p *paragraph) Replace(old string, new string) bool {
	text := p.GetText()

	index := index(text, old)

	if index == -1 { //verify if paragraph contains old string
		return false
	}
	Runchars := make([]character, len(text)) //create slice of characters to store char and style of char

	NewRunchars := make([]character, len(new)+(len(text)-len(old))) //crate slice to with size to put new string into
	char := 0
	for _, runn := range p.Runs {

		for _, c := range strings.Split(runn.GetText(), "") {

			Runchars[char] = character{char: c, runStyle: runn.Style}

			char++
		}
	}
	StyleBase := Runchars[index].runStyle //use the index char as style base

	copy(NewRunchars[:index], Runchars[:index])
	copy(NewRunchars[index+len(new):], Runchars[index+len(old):])

	for i, c := range strings.Split(new, "") {
		NewRunchars[i+index] = character{char: c, runStyle: StyleBase} //put new string into NewRunChars
	}
	//Rebuild Runs
	//p.ClearRuns()
	style := StyleBase
	p.ClearRuns()
	RunText := ""
	for _, t := range NewRunchars {
		if t.runStyle.IsEqual(style) {
			RunText += t.char
		} else {
			r := p.CreateRun()
			r.Style.Clone(style)
			r.SetText(RunText)
			RunText = t.char
			style = t.runStyle

		}
	}
	if RunText != "" {
		r := p.CreateRun()
		r.Style.Clone(style)
		r.SetText(RunText)
	}

	return true
}
