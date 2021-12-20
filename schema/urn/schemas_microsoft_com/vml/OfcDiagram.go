// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package vml

import (
	"encoding/xml"
	"strconv"

	"github.com/unidoc/unioffice"
)

type OfcDiagram struct {
	OfcCT_Diagram
}

func NewOfcDiagram() *OfcDiagram {
	ret := &OfcDiagram{}
	ret.OfcCT_Diagram = *NewOfcCT_Diagram()
	return ret
}

func (m *OfcDiagram) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "o:diagram"
	return m.OfcCT_Diagram.MarshalXML(e, start)
}

func (m *OfcDiagram) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_Diagram = *NewOfcCT_Diagram()
	for _, attr := range start.Attr {
		if attr.Name.Local == "dgmstyle" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmstyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoformat" {
			m.AutoformatAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "reverse" {
			m.ReverseAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "autolayout" {
			m.AutolayoutAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dgmscalex" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmscalexAttr = &parsed
			continue
		}
		if attr.Name.Local == "dgmscaley" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmscaleyAttr = &parsed
			continue
		}
		if attr.Name.Local == "dgmfontsize" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmfontsizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "constrainbounds" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ConstrainboundsAttr = &parsed
			continue
		}
		if attr.Name.Local == "dgmbasetextscale" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmbasetextscaleAttr = &parsed
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lOfcDiagram:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "relationtable"}:
				m.Relationtable = NewOfcCT_RelationTable()
				if err := d.DecodeElement(m.Relationtable, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on OfcDiagram %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOfcDiagram
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcDiagram and its children
func (m *OfcDiagram) Validate() error {
	return m.ValidateWithPath("OfcDiagram")
}

// ValidateWithPath validates the OfcDiagram and its children, prefixing error messages with path
func (m *OfcDiagram) ValidateWithPath(path string) error {
	if err := m.OfcCT_Diagram.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}