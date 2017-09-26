// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"log"
)

type CT_TransitionSoundAction struct {
	// Start Sound Action
	StSnd *CT_TransitionStartSoundAction
	// Stop Sound Action
	EndSnd *CT_Empty
}

func NewCT_TransitionSoundAction() *CT_TransitionSoundAction {
	ret := &CT_TransitionSoundAction{}
	return ret
}

func (m *CT_TransitionSoundAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.StSnd != nil {
		sestSnd := xml.StartElement{Name: xml.Name{Local: "p:stSnd"}}
		e.EncodeElement(m.StSnd, sestSnd)
	}
	if m.EndSnd != nil {
		seendSnd := xml.StartElement{Name: xml.Name{Local: "p:endSnd"}}
		e.EncodeElement(m.EndSnd, seendSnd)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TransitionSoundAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TransitionSoundAction:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "stSnd"}:
				m.StSnd = NewCT_TransitionStartSoundAction()
				if err := d.DecodeElement(m.StSnd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "endSnd"}:
				m.EndSnd = NewCT_Empty()
				if err := d.DecodeElement(m.EndSnd, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TransitionSoundAction %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TransitionSoundAction
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TransitionSoundAction and its children
func (m *CT_TransitionSoundAction) Validate() error {
	return m.ValidateWithPath("CT_TransitionSoundAction")
}

// ValidateWithPath validates the CT_TransitionSoundAction and its children, prefixing error messages with path
func (m *CT_TransitionSoundAction) ValidateWithPath(path string) error {
	if m.StSnd != nil {
		if err := m.StSnd.ValidateWithPath(path + "/StSnd"); err != nil {
			return err
		}
	}
	if m.EndSnd != nil {
		if err := m.EndSnd.ValidateWithPath(path + "/EndSnd"); err != nil {
			return err
		}
	}
	return nil
}
