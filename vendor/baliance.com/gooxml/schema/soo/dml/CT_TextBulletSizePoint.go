// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_TextBulletSizePoint struct {
	ValAttr int32
}

func NewCT_TextBulletSizePoint() *CT_TextBulletSizePoint {
	ret := &CT_TextBulletSizePoint{}
	ret.ValAttr = 100
	return ret
}

func (m *CT_TextBulletSizePoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextBulletSizePoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = 100
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ValAttr = int32(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextBulletSizePoint: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextBulletSizePoint and its children
func (m *CT_TextBulletSizePoint) Validate() error {
	return m.ValidateWithPath("CT_TextBulletSizePoint")
}

// ValidateWithPath validates the CT_TextBulletSizePoint and its children, prefixing error messages with path
func (m *CT_TextBulletSizePoint) ValidateWithPath(path string) error {
	if m.ValAttr < 100 {
		return fmt.Errorf("%s/m.ValAttr must be >= 100 (have %v)", path, m.ValAttr)
	}
	if m.ValAttr > 400000 {
		return fmt.Errorf("%s/m.ValAttr must be <= 400000 (have %v)", path, m.ValAttr)
	}
	return nil
}
