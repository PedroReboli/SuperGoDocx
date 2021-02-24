package docx

import (
	"errors"
	"strconv"

	"github.com/PedroReboli/go-xmldom"
)

//GetBold Return true if TextStyle has Bold
func (t *TextStyle) GetBold() bool {
	return t.bold
}

//SetBold Set bold value to TextStyle
func (t *TextStyle) SetBold(value bool) {
	if value && len(t.Document.GetChildren("b")) == 0 {
		n := t.Document.CreateNode("b")
		n.NS = &xmldom.Namespace{Name: "w"}

	}
	if value == false && len(t.Document.GetChildren("b")) != 0 {
		t.Document.RemoveChild(t.Document.GetChild("b"))
	}
	t.bold = value

}

//GetItalic Return true if TextStyle has Italic
func (t *TextStyle) GetItalic() bool {
	return t.italic
}

//SetItalic Set Italic value to TextStyle
func (t *TextStyle) SetItalic(value bool) {

	if value && len(t.Document.GetChildren("i")) == 0 {
		n := t.Document.CreateNode("i")
		n.NS = &xmldom.Namespace{Name: "w"}
	}
	if value == false && len(t.Document.GetChildren("i")) != 0 {
		t.Document.RemoveChild(t.Document.GetChild("i"))
	}
	t.italic = value

}

//GetCaps Return true if TextStyle has Caps
func (t *TextStyle) GetCaps() bool {
	return t.caps
}

//SetCaps Set value to Caps
func (t *TextStyle) SetCaps(value bool) {

	if value && len(t.Document.GetChildren("caps")) == 0 {
		n := t.Document.CreateNode("caps")
		n.NS = &xmldom.Namespace{Name: "w"}
	}
	if value == false && len(t.Document.GetChildren("caps")) != 0 {
		t.Document.RemoveChild(t.Document.GetChild("caps"))
	}
	t.caps = value
}

//GetColor if has color return Color in hex else return empty string
func (t *TextStyle) GetColor() string {
	return t.color
}

//SetColor Set string to Color. value must be in hex RRGGBB or auto
func (t *TextStyle) SetColor(value string) {
	if value != "" && len(t.Document.GetChildren("color")) == 0 {
		n := t.Document.CreateNode("color")
		n.SetAttributeValue("val", value)
		n.GetAttribute("val").NS = &xmldom.Namespace{Name: "w"}
		n.NS = &xmldom.Namespace{Name: "w"}
	}
	if value == "" && len(t.Document.GetChildren("color")) != 0 {
		t.Document.RemoveChild(t.Document.GetChild("color"))
	}
	t.color = value
}

//GetDstrike Return true if TextStyle has Double Strike
func (t *TextStyle) GetDstrike() bool {
	return t.dstrike
}

//SetDstrike set value to Dstrike
func (t *TextStyle) SetDstrike(value bool) {
	if value && len(t.Document.GetChildren("dstrike")) == 0 {
		n := t.Document.CreateNode("dstrike")
		n.SetAttributeValue("val", "true")
		n.GetAttribute("val").NS = &xmldom.Namespace{Name: "w"}
		n.NS = &xmldom.Namespace{Name: "w"}
	}
	if value == false && len(t.Document.GetChildren("dstrike")) != 0 {
		t.Document.RemoveChild(t.Document.GetChild("dstrike"))
	}
	t.dstrike = value

}

//GetEmboss Return true if TextStyle has Emboss
func (t *TextStyle) GetEmboss() bool {
	return t.emboss
}

//SetEmboss set value to Emboss
func (t *TextStyle) SetEmboss(value bool) {
	if value && len(t.Document.GetChildren("emboss")) == 0 {
		n := t.Document.CreateNode("emboss")
		n.SetAttributeValue("val", "true")
		n.GetAttribute("val").NS = &xmldom.Namespace{Name: "w"}
		n.NS = &xmldom.Namespace{Name: "w"}
	}
	if value == false && len(t.Document.GetChildren("emboss")) != 0 {
		t.Document.RemoveChild(t.Document.GetChild("emboss"))
	}
	t.emboss = value

}

//GetImprint Return true if TextStyle has Imprint
func (t *TextStyle) GetImprint() bool {
	return t.imprint
}

//SetImprint set value to Imprint
func (t *TextStyle) SetImprint(value bool) {
	if value && len(t.Document.GetChildren("imprint")) == 0 {
		n := t.Document.CreateNode("imprint")
		n.SetAttributeValue("val", "true")
		n.GetAttribute("val").NS = &xmldom.Namespace{Name: "w"}
		n.NS = &xmldom.Namespace{Name: "w"}
	}
	if value == false && len(t.Document.GetChildren("imprint")) != 0 {
		t.Document.RemoveChild(t.Document.GetChild("imprint"))
	}
	t.imprint = value
}

//GetSize Return Size in half-points (half-poits = fontsize*2)
func (t *TextStyle) GetSize() int {
	return t.size
}

//SetSize set value to Size value in half-points (half-poits = fontsize*2)
func (t *TextStyle) SetSize(value int) {
	if value != t.size && len(t.Document.GetChildren("sz")) == 0 {
		n := t.Document.CreateNode("sz")
		n.SetAttributeValue("val", strconv.Itoa(value))
		n.GetAttribute("val").NS = &xmldom.Namespace{Name: "w"}
		n.NS = &xmldom.Namespace{Name: "w"}
	}
	if value != t.size && len(t.Document.GetChildren("sz")) != 0 {
		t.Document.GetChild("sz").SetAttributeValue("val", strconv.Itoa(value)).NS = &xmldom.Namespace{Name: "w"}
	}
	t.size = value
}

//GetVanish Return true if TextStyle has Vanish
func (t *TextStyle) GetVanish() bool {
	return t.vanish
}

//SetVanish set value to Vanish
func (t *TextStyle) SetVanish(value bool) {

	if value && len(t.Document.GetChildren("vanish")) == 0 {
		t.Document.CreateNode("vanish")
		//n.NS = &xmldom.Namespace{Name: "w"}
	}
	if value == false && len(t.Document.GetChildren("vanish")) != 0 {
		t.Document.RemoveChild(t.Document.GetChild("vanish"))
	}
	t.vanish = value

}

//GetVertAlign Return value of VertAlign can be: basline, subscript, superscript
func (t *TextStyle) GetVertAlign() string {
	return t.vertAlign
}

//SetVertAlign set value to VertAlign can be: basline , subscript , superscript
func (t *TextStyle) SetVertAlign(value string) {
	if value != t.vertAlign && len(t.Document.GetChildren("sz")) == 0 {
		n := t.Document.CreateNode("sz")
		n.SetAttributeValue("val", value)
		n.GetAttribute("val").NS = &xmldom.Namespace{Name: "w"}
	}
	if value != t.vertAlign && len(t.Document.GetChildren("sz")) != 0 {
		t.Document.GetChild("sz").SetAttributeValue("val", value).NS = &xmldom.Namespace{Name: "w"}
	}
	t.vertAlign = value
}

//GetUnderline return (has underline , pattern , color)
func (t *TextStyle) GetUnderline() (bool, string, string) {
	if t.underline == false {
		return false, "", ""
	}
	return true, t.structUnderline.value, t.structUnderline.color

}

//SetUnderline set value to Underline (pattern , color)
//if color is empty default value is auto
//possible values for pattern http://officeopenxml.com/WPtextFormatting.php
func (t *TextStyle) SetUnderline(pattern string, color string) {
	if pattern == "" {
		t.Document.RemoveChild(t.Document.GetChild("u"))
		t.structUnderline = sUnderline{value: "", color: ""}
		t.underline = false
		return
	}
	if len(t.Document.GetChildren("u")) == 0 {
		n := t.Document.CreateNode("u")
		n.NS = &xmldom.Namespace{Name: "w"}

	}
	n := t.Document.GetChild("u")
	n.SetAttributeValue("val", pattern)
	n.GetAttribute("val").NS = &xmldom.Namespace{Name: "w"}
	n.SetAttributeValue("color", color)
	if color == "" {
		n.SetAttributeValue("color", "auto")
	}
	n.GetAttribute("color").NS = &xmldom.Namespace{Name: "w"}
	t.structUnderline = sUnderline{value: pattern, color: color}
	t.underline = true
}

//GetFont input attribute possible attributes (ascii,hAnsi,cs) return font name
func (t *TextStyle) GetFont(attribute string) (string, error) {
	switch attribute {
	case "ascii":
		return t.font.ascii, nil
	case "hAnsi":
		return t.font.hAnsi, nil
	case "cs":
		return t.font.cs, nil
	default:
		return "", errors.New("wrong attribute " + attribute)
	}
}

//GetFonts return (ascii, hAnsi, cs)
func (t *TextStyle) GetFonts() (string, string, string) {
	return t.font.ascii, t.font.hAnsi, t.font.cs
}

//SetFont input attribute possible attributes (ascii,hAnsi,cs) return font name
func (t *TextStyle) SetFont(attribute string, FontName string) error {
	if len(t.Document.GetChildren("rFonts")) == 0 {
		n := t.Document.CreateNode("rFonts")
		n.NS = &xmldom.Namespace{Name: "w"}
	}
	n := t.Document.GetChild("rFonts")
	var a *xmldom.Attribute
	switch attribute {
	case "ascii":
		n.SetAttributeValue("ascii", FontName)
		a = n.GetAttribute("ascii")
		t.font.ascii = FontName
	case "hAnsi":
		n.SetAttributeValue("hAnsi", FontName)
		a = n.GetAttribute("hAnsi")
		t.font.hAnsi = FontName
	case "cs":
		n.SetAttributeValue("cs", FontName)
		a = n.GetAttribute("cs")
		t.font.cs = FontName
	default:
		return errors.New("wrong attribute " + attribute)
	}
	a.NS = &xmldom.Namespace{Name: "w"}
	return nil
}

//SetFonts input (ascii,hAnsi,cs)
func (t *TextStyle) SetFonts(ascii string, hAnsi string, cs string) {
	if len(t.Document.GetChildren("rFonts")) == 0 {
		n := t.Document.CreateNode("rFonts")
		n.NS = &xmldom.Namespace{Name: "w"}
	}

	n := t.Document.GetChild("rFonts")
	n.SetAttributeValue("ascii", ascii)
	n.SetAttributeValue("hAnsi", hAnsi)
	n.SetAttributeValue("cs", cs)

	n.GetAttribute("ascii").NS = &xmldom.Namespace{Name: "w"}
	n.GetAttribute("hAnsi").NS = &xmldom.Namespace{Name: "w"}
	n.GetAttribute("cs").NS = &xmldom.Namespace{Name: "w"}

	t.font.ascii = ascii
	t.font.hAnsi = hAnsi
	t.font.cs = cs

}

//IsEqual Check if has the sames settings as input
func (t *TextStyle) IsEqual(T TextStyle) bool {
	if t.bold != T.bold {
		return false
	}
	if t.italic != T.italic {
		return false
	}
	if t.caps != T.caps {
		return false
	}
	if t.color != T.color {
		return false
	}
	if t.dstrike != T.dstrike {
		return false
	}
	if t.emboss != T.emboss {
		return false
	}
	if t.imprint != T.imprint {
		return false
	}
	if t.outline != T.outline {
		return false
	}
	if t.shadow != T.shadow {
		return false
	}
	if t.smallCaps != T.smallCaps {
		return false
	}
	if t.strike != T.strike {
		return false
	}
	if t.size != T.size {
		return false
	}
	if t.vanish != T.vanish {
		return false
	}
	if t.vertAlign != T.vertAlign {
		return false
	}
	if t.underline != T.underline {
		return false
	}
	if t.structUnderline.value != T.structUnderline.value {
		return false
	}
	if t.structUnderline.color != T.structUnderline.color {
		return false
	}
	if t.font.ascii != T.font.ascii {
		return false
	}
	if t.font.hAnsi != T.font.hAnsi {
		return false
	}
	if t.font.cs != T.font.cs {
		return false
	}
	return true
}

//Clone copy every settings of input
func (t *TextStyle) Clone(T TextStyle) {
	t.SetBold(T.GetBold())
	t.SetItalic(T.GetItalic())
	t.SetCaps(T.GetCaps())
	t.SetColor(T.GetColor())
	t.SetDstrike(T.GetDstrike())
	t.SetEmboss(T.GetEmboss())
	t.SetImprint(T.GetImprint())
	t.SetSize(T.GetSize())
	t.SetVanish(T.GetVanish())
	t.SetVertAlign(T.GetVertAlign())
	V, pattern, color := T.GetUnderline()
	if V == true {
		t.SetUnderline(pattern, color)
	}
	t.SetFonts(T.GetFonts())
}
