package code2docx

import (
	"io/ioutil"
	"strings"

	"baliance.com/gooxml/document"
	"baliance.com/gooxml/schema/soo/wml"
)

type Outer struct {
	Config  *Config
	Base    string
	Files   []string
	Docfile string
	doc     *document.Document
}

func (o *Outer) WriteDoc() error {
	o.init()
	for _, file := range o.Files {
		filebytes, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		filename := strings.Replace(file, o.Base+"/", "", 1)

		o.writeToDoc(filename, filebytes)

	}
	o.doc.SaveToFile(o.Docfile)
	return nil
}

func (o *Outer) init() {
	doc := document.New()
	o.doc = doc
	hdr := doc.AddHeader()
	para := hdr.AddParagraph()
	para.AddRun().AddText(o.Config.Header)
	doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)

	para = doc.AddParagraph()
	para.SetStyle("Title")
	para.AddRun().AddText(o.Config.Title)

	para = doc.AddParagraph()
	para.AddRun().AddBreak()
}

func (o *Outer) writeToDoc(filename string, content []byte) {
	doc := o.doc
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(filename + ":")
	run.AddBreak()

	para = doc.AddParagraph()
	run = para.AddRun()
	run.AddText(string(content))
	run.AddBreak()
	run.AddBreak()
}
