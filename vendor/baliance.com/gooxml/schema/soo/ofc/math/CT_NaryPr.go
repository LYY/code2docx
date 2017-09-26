// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type CT_NaryPr struct {
	Chr     *CT_Char
	LimLoc  *CT_LimLoc
	Grow    *CT_OnOff
	SubHide *CT_OnOff
	SupHide *CT_OnOff
	CtrlPr  *CT_CtrlPr
}

func NewCT_NaryPr() *CT_NaryPr {
	ret := &CT_NaryPr{}
	return ret
}

func (m *CT_NaryPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Chr != nil {
		sechr := xml.StartElement{Name: xml.Name{Local: "m:chr"}}
		e.EncodeElement(m.Chr, sechr)
	}
	if m.LimLoc != nil {
		selimLoc := xml.StartElement{Name: xml.Name{Local: "m:limLoc"}}
		e.EncodeElement(m.LimLoc, selimLoc)
	}
	if m.Grow != nil {
		segrow := xml.StartElement{Name: xml.Name{Local: "m:grow"}}
		e.EncodeElement(m.Grow, segrow)
	}
	if m.SubHide != nil {
		sesubHide := xml.StartElement{Name: xml.Name{Local: "m:subHide"}}
		e.EncodeElement(m.SubHide, sesubHide)
	}
	if m.SupHide != nil {
		sesupHide := xml.StartElement{Name: xml.Name{Local: "m:supHide"}}
		e.EncodeElement(m.SupHide, sesupHide)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NaryPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_NaryPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "chr"}:
				m.Chr = NewCT_Char()
				if err := d.DecodeElement(m.Chr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "limLoc"}:
				m.LimLoc = NewCT_LimLoc()
				if err := d.DecodeElement(m.LimLoc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "grow"}:
				m.Grow = NewCT_OnOff()
				if err := d.DecodeElement(m.Grow, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "subHide"}:
				m.SubHide = NewCT_OnOff()
				if err := d.DecodeElement(m.SubHide, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "supHide"}:
				m.SupHide = NewCT_OnOff()
				if err := d.DecodeElement(m.SupHide, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "ctrlPr"}:
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_NaryPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NaryPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NaryPr and its children
func (m *CT_NaryPr) Validate() error {
	return m.ValidateWithPath("CT_NaryPr")
}

// ValidateWithPath validates the CT_NaryPr and its children, prefixing error messages with path
func (m *CT_NaryPr) ValidateWithPath(path string) error {
	if m.Chr != nil {
		if err := m.Chr.ValidateWithPath(path + "/Chr"); err != nil {
			return err
		}
	}
	if m.LimLoc != nil {
		if err := m.LimLoc.ValidateWithPath(path + "/LimLoc"); err != nil {
			return err
		}
	}
	if m.Grow != nil {
		if err := m.Grow.ValidateWithPath(path + "/Grow"); err != nil {
			return err
		}
	}
	if m.SubHide != nil {
		if err := m.SubHide.ValidateWithPath(path + "/SubHide"); err != nil {
			return err
		}
	}
	if m.SupHide != nil {
		if err := m.SupHide.ValidateWithPath(path + "/SupHide"); err != nil {
			return err
		}
	}
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
