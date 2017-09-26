// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/soo/wml"
)

func TestCT_NumPicBulletConstructor(t *testing.T) {
	v := wml.NewCT_NumPicBullet()
	if v == nil {
		t.Errorf("wml.NewCT_NumPicBullet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_NumPicBullet should validate: %s", err)
	}
}

func TestCT_NumPicBulletMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_NumPicBullet()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_NumPicBullet()
	xml.Unmarshal(buf, v2)
}
