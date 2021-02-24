package docx

import "github.com/PedroReboli/go-xmldom"

func (r *run) setup() {
	r.Document.CreateNode("rPr").NS = &xmldom.Namespace{Name: "w"}
	r.Document.CreateNode("t").NS = &xmldom.Namespace{Name: "w"}
	r.Style = TextStyle{}
	r.Style.populate(r.Document.GetChild("rPr"))
}
