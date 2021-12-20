// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type WdCT_PosHChoice struct {
	Align     WdST_AlignH
	PosOffset *int32
}

func NewWdCT_PosHChoice() *WdCT_PosHChoice {
	ret := &WdCT_PosHChoice{}
	return ret
}

func (m *WdCT_PosHChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Align != WdST_AlignHUnset {
		sealign := xml.StartElement{Name: xml.Name{Local: "wp:align"}}
		e.EncodeElement(m.Align, sealign)
	}
	if m.PosOffset != nil {
		seposOffset := xml.StartElement{Name: xml.Name{Local: "wp:posOffset"}}
		e.EncodeElement(m.PosOffset, seposOffset)
	}
	return nil
}

func (m *WdCT_PosHChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lWdCT_PosHChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "align"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "align"}:
				m.Align = WdST_AlignHUnset
				if err := d.DecodeElement(&m.Align, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "posOffset"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "posOffset"}:
				m.PosOffset = new(int32)
				if err := d.DecodeElement(m.PosOffset, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on WdCT_PosHChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_PosHChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_PosHChoice and its children
func (m *WdCT_PosHChoice) Validate() error {
	return m.ValidateWithPath("WdCT_PosHChoice")
}

// ValidateWithPath validates the WdCT_PosHChoice and its children, prefixing error messages with path
func (m *WdCT_PosHChoice) ValidateWithPath(path string) error {
	if err := m.Align.ValidateWithPath(path + "/Align"); err != nil {
		return err
	}
	return nil
}