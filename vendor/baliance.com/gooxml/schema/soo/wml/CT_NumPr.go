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
	"log"
)

type CT_NumPr struct {
	// Numbering Level Reference
	Ilvl *CT_DecimalNumber
	// Numbering Definition Instance Reference
	NumId *CT_DecimalNumber
	// Previous Paragraph Numbering Properties
	NumberingChange *CT_TrackChangeNumbering
	// Inserted Numbering Properties
	Ins *CT_TrackChange
}

func NewCT_NumPr() *CT_NumPr {
	ret := &CT_NumPr{}
	return ret
}

func (m *CT_NumPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Ilvl != nil {
		seilvl := xml.StartElement{Name: xml.Name{Local: "w:ilvl"}}
		e.EncodeElement(m.Ilvl, seilvl)
	}
	if m.NumId != nil {
		senumId := xml.StartElement{Name: xml.Name{Local: "w:numId"}}
		e.EncodeElement(m.NumId, senumId)
	}
	if m.NumberingChange != nil {
		senumberingChange := xml.StartElement{Name: xml.Name{Local: "w:numberingChange"}}
		e.EncodeElement(m.NumberingChange, senumberingChange)
	}
	if m.Ins != nil {
		seins := xml.StartElement{Name: xml.Name{Local: "w:ins"}}
		e.EncodeElement(m.Ins, seins)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NumPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_NumPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ilvl"}:
				m.Ilvl = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.Ilvl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "numId"}:
				m.NumId = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.NumId, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "numberingChange"}:
				m.NumberingChange = NewCT_TrackChangeNumbering()
				if err := d.DecodeElement(m.NumberingChange, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ins"}:
				m.Ins = NewCT_TrackChange()
				if err := d.DecodeElement(m.Ins, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_NumPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NumPr and its children
func (m *CT_NumPr) Validate() error {
	return m.ValidateWithPath("CT_NumPr")
}

// ValidateWithPath validates the CT_NumPr and its children, prefixing error messages with path
func (m *CT_NumPr) ValidateWithPath(path string) error {
	if m.Ilvl != nil {
		if err := m.Ilvl.ValidateWithPath(path + "/Ilvl"); err != nil {
			return err
		}
	}
	if m.NumId != nil {
		if err := m.NumId.ValidateWithPath(path + "/NumId"); err != nil {
			return err
		}
	}
	if m.NumberingChange != nil {
		if err := m.NumberingChange.ValidateWithPath(path + "/NumberingChange"); err != nil {
			return err
		}
	}
	if m.Ins != nil {
		if err := m.Ins.ValidateWithPath(path + "/Ins"); err != nil {
			return err
		}
	}
	return nil
}
