// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcCT_ColorMenu struct {
	StrokecolorAttr    *string
	FillcolorAttr      *string
	ShadowcolorAttr    *string
	ExtrusioncolorAttr *string
	ExtAttr            ST_Ext
}

func NewOfcCT_ColorMenu() *OfcCT_ColorMenu {
	ret := &OfcCT_ColorMenu{}
	return ret
}

func (m *OfcCT_ColorMenu) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.StrokecolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "strokecolor"},
			Value: fmt.Sprintf("%v", *m.StrokecolorAttr)})
	}
	if m.FillcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillcolor"},
			Value: fmt.Sprintf("%v", *m.FillcolorAttr)})
	}
	if m.ShadowcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "shadowcolor"},
			Value: fmt.Sprintf("%v", *m.ShadowcolorAttr)})
	}
	if m.ExtrusioncolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "extrusioncolor"},
			Value: fmt.Sprintf("%v", *m.ExtrusioncolorAttr)})
	}
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_ColorMenu) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "strokecolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StrokecolorAttr = &parsed
		}
		if attr.Name.Local == "fillcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FillcolorAttr = &parsed
		}
		if attr.Name.Local == "shadowcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ShadowcolorAttr = &parsed
		}
		if attr.Name.Local == "extrusioncolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ExtrusioncolorAttr = &parsed
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_ColorMenu: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_ColorMenu and its children
func (m *OfcCT_ColorMenu) Validate() error {
	return m.ValidateWithPath("OfcCT_ColorMenu")
}

// ValidateWithPath validates the OfcCT_ColorMenu and its children, prefixing error messages with path
func (m *OfcCT_ColorMenu) ValidateWithPath(path string) error {
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
