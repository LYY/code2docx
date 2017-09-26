// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"log"
	"strconv"
)

type Table struct {
	CT_Table
}

func NewTable() *Table {
	ret := &Table{}
	ret.CT_Table = *NewCT_Table()
	return ret
}

func (m *Table) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "ma:table"
	return m.CT_Table.MarshalXML(e, start)
}

func (m *Table) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Table = *NewCT_Table()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "displayName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DisplayNameAttr = parsed
		}
		if attr.Name.Local == "comment" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CommentAttr = &parsed
		}
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = parsed
		}
		if attr.Name.Local == "tableType" {
			m.TableTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "headerRowCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.HeaderRowCountAttr = &pt
		}
		if attr.Name.Local == "insertRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.InsertRowAttr = &parsed
		}
		if attr.Name.Local == "insertRowShift" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.InsertRowShiftAttr = &parsed
		}
		if attr.Name.Local == "totalsRowCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TotalsRowCountAttr = &pt
		}
		if attr.Name.Local == "totalsRowShown" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TotalsRowShownAttr = &parsed
		}
		if attr.Name.Local == "published" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PublishedAttr = &parsed
		}
		if attr.Name.Local == "headerRowDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.HeaderRowDxfIdAttr = &pt
		}
		if attr.Name.Local == "dataDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DataDxfIdAttr = &pt
		}
		if attr.Name.Local == "totalsRowDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TotalsRowDxfIdAttr = &pt
		}
		if attr.Name.Local == "headerRowBorderDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.HeaderRowBorderDxfIdAttr = &pt
		}
		if attr.Name.Local == "tableBorderDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TableBorderDxfIdAttr = &pt
		}
		if attr.Name.Local == "totalsRowBorderDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TotalsRowBorderDxfIdAttr = &pt
		}
		if attr.Name.Local == "headerRowCellStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HeaderRowCellStyleAttr = &parsed
		}
		if attr.Name.Local == "dataCellStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DataCellStyleAttr = &parsed
		}
		if attr.Name.Local == "totalsRowCellStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TotalsRowCellStyleAttr = &parsed
		}
		if attr.Name.Local == "connectionId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ConnectionIdAttr = &pt
		}
	}
lTable:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "autoFilter"}:
				m.AutoFilter = NewCT_AutoFilter()
				if err := d.DecodeElement(m.AutoFilter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sortState"}:
				m.SortState = NewCT_SortState()
				if err := d.DecodeElement(m.SortState, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tableColumns"}:
				if err := d.DecodeElement(m.TableColumns, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tableStyleInfo"}:
				m.TableStyleInfo = NewCT_TableStyleInfo()
				if err := d.DecodeElement(m.TableStyleInfo, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on Table %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lTable
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Table and its children
func (m *Table) Validate() error {
	return m.ValidateWithPath("Table")
}

// ValidateWithPath validates the Table and its children, prefixing error messages with path
func (m *Table) ValidateWithPath(path string) error {
	if err := m.CT_Table.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
