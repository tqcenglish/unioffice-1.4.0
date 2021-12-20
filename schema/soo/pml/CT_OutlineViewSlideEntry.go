// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_OutlineViewSlideEntry struct {
	IdAttr string
	// Collapsed
	CollapseAttr *bool
}

func NewCT_OutlineViewSlideEntry() *CT_OutlineViewSlideEntry {
	ret := &CT_OutlineViewSlideEntry{}
	return ret
}

func (m *CT_OutlineViewSlideEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.CollapseAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "collapse"},
			Value: fmt.Sprintf("%d", b2i(*m.CollapseAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OutlineViewSlideEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
		if attr.Name.Local == "collapse" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CollapseAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_OutlineViewSlideEntry: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_OutlineViewSlideEntry and its children
func (m *CT_OutlineViewSlideEntry) Validate() error {
	return m.ValidateWithPath("CT_OutlineViewSlideEntry")
}

// ValidateWithPath validates the CT_OutlineViewSlideEntry and its children, prefixing error messages with path
func (m *CT_OutlineViewSlideEntry) ValidateWithPath(path string) error {
	return nil
}