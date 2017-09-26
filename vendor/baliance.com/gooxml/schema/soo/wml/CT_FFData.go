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

type CT_FFData struct {
	// Form Field Name
	Name []*CT_FFName
	// Form Field Label
	Label []*CT_DecimalNumber
	// Form Field Navigation Order Index
	TabIndex []*CT_UnsignedDecimalNumber
	// Form Field Enabled
	Enabled []*CT_OnOff
	// Recalculate Fields When Current Field Is Modified
	CalcOnExit []*CT_OnOff
	// Script Function to Execute on Form Field Entry
	EntryMacro []*CT_MacroName
	// Script Function to Execute on Form Field Exit
	ExitMacro []*CT_MacroName
	// Associated Help Text
	HelpText []*CT_FFHelpText
	// Associated Status Text
	StatusText []*CT_FFStatusText
	CheckBox   *CT_FFCheckBox
	DdList     *CT_FFDDList
	TextInput  *CT_FFTextInput
}

func NewCT_FFData() *CT_FFData {
	ret := &CT_FFData{}
	return ret
}

func (m *CT_FFData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Name != nil {
		sename := xml.StartElement{Name: xml.Name{Local: "w:name"}}
		for _, c := range m.Name {
			e.EncodeElement(c, sename)
		}
	}
	if m.Label != nil {
		selabel := xml.StartElement{Name: xml.Name{Local: "w:label"}}
		for _, c := range m.Label {
			e.EncodeElement(c, selabel)
		}
	}
	if m.TabIndex != nil {
		setabIndex := xml.StartElement{Name: xml.Name{Local: "w:tabIndex"}}
		for _, c := range m.TabIndex {
			e.EncodeElement(c, setabIndex)
		}
	}
	if m.Enabled != nil {
		seenabled := xml.StartElement{Name: xml.Name{Local: "w:enabled"}}
		for _, c := range m.Enabled {
			e.EncodeElement(c, seenabled)
		}
	}
	if m.CalcOnExit != nil {
		secalcOnExit := xml.StartElement{Name: xml.Name{Local: "w:calcOnExit"}}
		for _, c := range m.CalcOnExit {
			e.EncodeElement(c, secalcOnExit)
		}
	}
	if m.EntryMacro != nil {
		seentryMacro := xml.StartElement{Name: xml.Name{Local: "w:entryMacro"}}
		for _, c := range m.EntryMacro {
			e.EncodeElement(c, seentryMacro)
		}
	}
	if m.ExitMacro != nil {
		seexitMacro := xml.StartElement{Name: xml.Name{Local: "w:exitMacro"}}
		for _, c := range m.ExitMacro {
			e.EncodeElement(c, seexitMacro)
		}
	}
	if m.HelpText != nil {
		sehelpText := xml.StartElement{Name: xml.Name{Local: "w:helpText"}}
		for _, c := range m.HelpText {
			e.EncodeElement(c, sehelpText)
		}
	}
	if m.StatusText != nil {
		sestatusText := xml.StartElement{Name: xml.Name{Local: "w:statusText"}}
		for _, c := range m.StatusText {
			e.EncodeElement(c, sestatusText)
		}
	}
	if m.CheckBox != nil {
		secheckBox := xml.StartElement{Name: xml.Name{Local: "w:checkBox"}}
		e.EncodeElement(m.CheckBox, secheckBox)
	}
	if m.DdList != nil {
		seddList := xml.StartElement{Name: xml.Name{Local: "w:ddList"}}
		e.EncodeElement(m.DdList, seddList)
	}
	if m.TextInput != nil {
		setextInput := xml.StartElement{Name: xml.Name{Local: "w:textInput"}}
		e.EncodeElement(m.TextInput, setextInput)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FFData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FFData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "name"}:
				tmp := NewCT_FFName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Name = append(m.Name, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "label"}:
				tmp := NewCT_DecimalNumber()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Label = append(m.Label, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tabIndex"}:
				tmp := NewCT_UnsignedDecimalNumber()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TabIndex = append(m.TabIndex, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "enabled"}:
				tmp := NewCT_OnOff()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Enabled = append(m.Enabled, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "calcOnExit"}:
				tmp := NewCT_OnOff()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CalcOnExit = append(m.CalcOnExit, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "entryMacro"}:
				tmp := NewCT_MacroName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.EntryMacro = append(m.EntryMacro, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "exitMacro"}:
				tmp := NewCT_MacroName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ExitMacro = append(m.ExitMacro, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "helpText"}:
				tmp := NewCT_FFHelpText()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.HelpText = append(m.HelpText, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "statusText"}:
				tmp := NewCT_FFStatusText()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.StatusText = append(m.StatusText, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "checkBox"}:
				m.CheckBox = NewCT_FFCheckBox()
				if err := d.DecodeElement(m.CheckBox, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ddList"}:
				m.DdList = NewCT_FFDDList()
				if err := d.DecodeElement(m.DdList, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "textInput"}:
				m.TextInput = NewCT_FFTextInput()
				if err := d.DecodeElement(m.TextInput, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_FFData %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FFData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FFData and its children
func (m *CT_FFData) Validate() error {
	return m.ValidateWithPath("CT_FFData")
}

// ValidateWithPath validates the CT_FFData and its children, prefixing error messages with path
func (m *CT_FFData) ValidateWithPath(path string) error {
	for i, v := range m.Name {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Name[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Label {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Label[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.TabIndex {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TabIndex[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Enabled {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Enabled[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.CalcOnExit {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CalcOnExit[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.EntryMacro {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EntryMacro[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ExitMacro {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ExitMacro[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.HelpText {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/HelpText[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.StatusText {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/StatusText[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.CheckBox != nil {
		if err := m.CheckBox.ValidateWithPath(path + "/CheckBox"); err != nil {
			return err
		}
	}
	if m.DdList != nil {
		if err := m.DdList.ValidateWithPath(path + "/DdList"); err != nil {
			return err
		}
	}
	if m.TextInput != nil {
		if err := m.TextInput.ValidateWithPath(path + "/TextInput"); err != nil {
			return err
		}
	}
	return nil
}
