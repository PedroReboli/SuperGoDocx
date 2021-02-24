package docx

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"

	"github.com/PedroReboli/go-xmldom"
)

//Docx is the base for the module
type Docx struct {
	path       string
	zipcontent []*zip.File
	Document   wDocument
}

/*
``` go
type  Docx  struct {
	path string
	zipcontent []*zip.File
	Document wDocument
}
```
*/

//Open is to open a docx file
func Open(path string) (Docx, error) {
	d := Docx{}
	d.path = path
	reader, err := zip.OpenReader(path)
	if err != nil {
		return Docx{}, err
	}
	files := reader.File
	d.zipcontent = files

	for _, file := range d.zipcontent {
		if file.Name == "word/document.xml" {
			reader, err := file.Open()
			if err != nil {
				return Docx{}, err
			}
			defer reader.Close()

			content, err := ioutil.ReadAll(reader)
			doc := xmldom.Must(xmldom.ParseXML((string)(content)))
			root := doc.Root
			root.NS = root.GetChild("body").NS
			d.Document = wDocument{}
			d.Document.populate(root)
			/*err = xml.Unmarshal(content, &d.Document)
			if err != nil {
				fmt.Println("erro no unmarshal")
				return err
			}*/

			//d.Document = &doc
			//fmt.Println(d.Document.Body.Paragraph[3].Runs[2])
			//fmt.Println(d.Xmlstruct.XMLName)
		}
	}
	return d, nil

}

//SaveAs reconstruct zip and save in path
func (d *Docx) SaveAs(path string) error {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	for _, file := range d.zipcontent {
		f, err := w.Create(file.Name)
		if err != nil {
			return err
		}
		if file.Name == "word/document.xml" {
			data := d.Document.Document.XML()
			_, err = f.Write(([]byte)(data))
			if err != nil {
				return err
			}
		} else {
			reader, _ := file.Open()
			content, err := ioutil.ReadAll(reader)
			_, err = f.Write(content)
			if err != nil {
				return err
			}
		}

	}
	err := w.Close()
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = buf.WriteTo(f)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	//ioutil.WriteFile(path,  , 0)
	return nil

}

//Save file overwriting the original file
func (d *Docx) Save() error {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	for _, file := range d.zipcontent {
		f, err := w.Create(file.Name)
		if err != nil {
			return err
		}
		if file.Name == "word/document.xml" {
			data := d.Document.Document.XML()
			_, err = f.Write(([]byte)(data))
			if err != nil {
				return err
			}
		} else {
			reader, _ := file.Open()
			content, err := ioutil.ReadAll(reader)
			_, err = f.Write(content)
			if err != nil {
				return err
			}
		}

	}
	err := w.Close()
	if err != nil {
		return err
	}
	f, err := os.Create(d.path)
	if err != nil {
		return err
	}
	_, err = buf.WriteTo(f)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	//ioutil.WriteFile(path,  , 0)
	return nil

}

//SaveToMemory return docx file in a Byte Slice
func (d *Docx) SaveToMemory() ([]byte, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	for _, file := range d.zipcontent {
		f, err := w.Create(file.Name)
		if err != nil {
			return nil, err
		}
		if file.Name == "word/document.xml" {
			data := d.Document.Document.XML()
			_, err = f.Write(([]byte)(data))
			if err != nil {
				return nil, err
			}
		} else {
			reader, _ := file.Open()
			content, err := ioutil.ReadAll(reader)
			_, err = f.Write(content)
			if err != nil {
				return nil, err
			}
		}
	}
	w.Close()
	return buf.Bytes(), nil

}
