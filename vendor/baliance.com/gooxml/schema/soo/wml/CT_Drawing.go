// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Drawing struct {
	Anchor []*WdAnchor
	Inline []*WdInline
}

func NewCT_Drawing() *CT_Drawing {
	ret := &CT_Drawing{}
	return ret
}

func (m *CT_Drawing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Anchor != nil {
		seanchor := xml.StartElement{Name: xml.Name{Local: "wp:anchor"}}
		for _, c := range m.Anchor {
			e.EncodeElement(c, seanchor)
		}
	}
	if m.Inline != nil {
		seinline := xml.StartElement{Name: xml.Name{Local: "wp:inline"}}
		for _, c := range m.Inline {
			e.EncodeElement(c, seinline)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Drawing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Drawing:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "anchor"}:
				tmp := NewWdAnchor()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Anchor = append(m.Anchor, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "inline"}:
				tmp := NewWdInline()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Inline = append(m.Inline, tmp)
			default:
				log.Printf("skipping unsupported element on CT_Drawing %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Drawing
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Drawing and its children
func (m *CT_Drawing) Validate() error {
	return m.ValidateWithPath("CT_Drawing")
}

// ValidateWithPath validates the CT_Drawing and its children, prefixing error messages with path
func (m *CT_Drawing) ValidateWithPath(path string) error {
	for i, v := range m.Anchor {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Anchor[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Inline {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Inline[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
