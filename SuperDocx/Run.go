package docx

import (
	"strings"

	"github.com/PedroReboli/go-xmldom"
)

//-----utils----//
func (r *run) GetText() string {
	if len(r.Document.GetChildren("t")) == 0 {
		return ""
	}
	return r.Document.GetChild("t").Text
}
func (r *run) SetText(text string) {
	if len(r.Document.GetChildren("t")) == 0 {
		return
	}

	splited := strings.Split(text, "")
	if len(splited) != 0 && (splited[0] == " " || splited[len(splited)-1] == " ") {
		r.Document.GetChild("t").SetAttributeValue("space", "preserve")
		r.Document.GetChild("t").GetAttribute("space").NS = &xmldom.Namespace{Name: "xml"}
	} else {
		r.Document.GetChild("t").RemoveAttribute("space")
	}
	r.Document.GetChild("t").Text = text
}
