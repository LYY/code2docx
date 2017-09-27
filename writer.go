package code2docx

import (
	"bufio"
	"log"
	"os"
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

		filename := strings.Replace(file, o.Base, "", 1)

		o.writeToDoc(filename, file)

	}
	o.doc.SaveToFile(o.Docfile)
	return nil
}

func (o *Outer) init() {
	doc := document.New()
	o.doc = doc
	hdr := doc.AddHeader()
	para := hdr.AddParagraph()
	setParagraphAlignCenter(para)
	para.AddRun().AddText(o.Config.Header)
	doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)

	para = doc.AddParagraph()
	para.SetStyle("Title")
	setParagraphAlignCenter(para)
	para.AddRun().AddText(o.Config.Title)

	para = doc.AddParagraph()
	para.AddRun().AddText("")
}

func ensurePPr(para document.Paragraph) {
	if para.X().PPr == nil {
		para.X().PPr = wml.NewCT_PPr()
	}
}
func setParagraphAlignCenter(para document.Paragraph) {
	ensurePPr(para)
	paraAlign := wml.NewCT_Jc()
	paraAlign.ValAttr = wml.ST_JcCenter
	para.X().PPr.Jc = paraAlign
}

func (o *Outer) writeToDoc(filename, filepath string) {
	doc := o.doc
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(filename + ":")
	run.AddBreak()

	fio, err := os.Open(filepath)
	if err != nil {
		log.Fatal("open file error: ", err)
	}
	defer fio.Close()
	fbscaner := bufio.NewScanner(fio)
	for fbscaner.Scan() {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(fbscaner.Text())
	}
	if err := fbscaner.Err(); err != nil {
		log.Fatal("read file error: ", err)
	}

	para = doc.AddParagraph()
	run = para.AddRun()
	run.AddBreak()
}
