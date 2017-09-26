// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package word

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type Borderbottom struct {
	CT_Border
}

func NewBorderbottom() *Borderbottom {
	ret := &Borderbottom{}
	ret.CT_Border = *NewCT_Border()
	return ret
}

func (m *Borderbottom) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:word"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "borderbottom"
	return m.CT_Border.MarshalXML(e, start)
}

func (m *Borderbottom) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Border = *NewCT_Border()
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "width" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.WidthAttr = &pt
		}
		if attr.Name.Local == "shadow" {
			m.ShadowAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Borderbottom: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Borderbottom and its children
func (m *Borderbottom) Validate() error {
	return m.ValidateWithPath("Borderbottom")
}

// ValidateWithPath validates the Borderbottom and its children, prefixing error messages with path
func (m *Borderbottom) ValidateWithPath(path string) error {
	if err := m.CT_Border.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
