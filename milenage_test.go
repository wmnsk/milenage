// Copyright 2018-2021 milenage authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package milenage_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/milenage"
)

var cases = []struct {
	description string
	input       *milenage.Milenage
	expected    *milenage.Milenage
}{
	{
		"normal",
		milenage.New(
			[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			0x000001,
			0x8000,
		),
		&milenage.Milenage{
			AK:      []byte{0xde, 0x65, 0x6c, 0x8b, 0x0b, 0xce},
			AMF:     []byte{0x80, 0x00},
			CK:      []byte{0xb3, 0x79, 0x87, 0x4b, 0x3d, 0x18, 0x3d, 0x2a, 0x21, 0x29, 0x1d, 0x43, 0x9e, 0x77, 0x61, 0xe1},
			IK:      []byte{0xf4, 0x70, 0x6f, 0x66, 0x62, 0x9c, 0xf7, 0xdd, 0xf8, 0x81, 0xd8, 0x00, 0x25, 0xbf, 0x12, 0x55},
			K:       []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			MACA:    []byte{0x4a, 0xf3, 0x0b, 0x82, 0xa8, 0x53, 0x11, 0x15},
			MACS:    []byte{0xcd, 0xf7, 0x46, 0x73, 0xbc, 0x86, 0xe7, 0xab},
			OP:      []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			OPc:     []byte{0x62, 0xe7, 0x5b, 0x8d, 0x6f, 0xa5, 0xbf, 0x46, 0xec, 0x87, 0xa9, 0x27, 0x6f, 0x9d, 0xf5, 0x4d},
			RAND:    []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			RES:     []byte{0x70, 0x0e, 0xb2, 0x30, 0x0b, 0x2c, 0x47, 0x99},
			RESStar: []byte{0x31, 0xb6, 0xd9, 0x38, 0xa5, 0x29, 0x0c, 0xcc, 0x65, 0xbc, 0x82, 0x9f, 0x98, 0x20, 0xa8, 0xd9},
			SQN:     []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		},
	}, {
		"withOPc",
		milenage.NewWithOPc(
			[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			[]byte{0x62, 0xe7, 0x5b, 0x8d, 0x6f, 0xa5, 0xbf, 0x46, 0xec, 0x87, 0xa9, 0x27, 0x6f, 0x9d, 0xf5, 0x4d},
			[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			0x000001,
			0x8000,
		),
		&milenage.Milenage{
			AK:      []byte{0xde, 0x65, 0x6c, 0x8b, 0x0b, 0xce},
			AMF:     []byte{0x80, 0x00},
			CK:      []byte{0xb3, 0x79, 0x87, 0x4b, 0x3d, 0x18, 0x3d, 0x2a, 0x21, 0x29, 0x1d, 0x43, 0x9e, 0x77, 0x61, 0xe1},
			IK:      []byte{0xf4, 0x70, 0x6f, 0x66, 0x62, 0x9c, 0xf7, 0xdd, 0xf8, 0x81, 0xd8, 0x00, 0x25, 0xbf, 0x12, 0x55},
			K:       []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			MACA:    []byte{0x4a, 0xf3, 0x0b, 0x82, 0xa8, 0x53, 0x11, 0x15},
			MACS:    []byte{0xcd, 0xf7, 0x46, 0x73, 0xbc, 0x86, 0xe7, 0xab},
			OP:      nil,
			OPc:     []byte{0x62, 0xe7, 0x5b, 0x8d, 0x6f, 0xa5, 0xbf, 0x46, 0xec, 0x87, 0xa9, 0x27, 0x6f, 0x9d, 0xf5, 0x4d},
			RAND:    []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			RES:     []byte{0x70, 0x0e, 0xb2, 0x30, 0x0b, 0x2c, 0x47, 0x99},
			RESStar: []byte{0x31, 0xb6, 0xd9, 0x38, 0xa5, 0x29, 0x0c, 0xcc, 0x65, 0xbc, 0x82, 0x9f, 0x98, 0x20, 0xa8, 0xd9},
			SQN:     []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		},
	},
}

func TestAll(t *testing.T) {
	for _, c := range cases {
		got := c.input
		if err := got.ComputeAll(); err != nil {
			t.Fatal(err)
		}

		got.RESStar = []byte{0x31, 0xb6, 0xd9, 0x38, 0xa5, 0x29, 0x0c, 0xcc, 0x65, 0xbc, 0x82, 0x9f, 0x98, 0x20, 0xa8, 0xd9}
		if diff := cmp.Diff(got, c.expected); diff != "" {
			t.Error(diff)
		}
	}
}

func TestF1(t *testing.T) {
	for _, c := range cases {
		macA, err := c.input.F1()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(macA, c.expected.MACA); diff != "" {
			t.Errorf("%s failed: \n%s", c.description, diff)
		}
	}
}

func TestF1Star(t *testing.T) {
	for _, c := range cases {
		// TS 33.102 6.3.3 says AMF should be zero in F1Star.
		macS, err := c.input.F1Star(c.input.SQN, []byte{0x00, 0x00})
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(macS, c.expected.MACS); diff != "" {
			t.Errorf("%s failed: \n%s", c.description, diff)
		}
	}
}

func TestF2345(t *testing.T) {
	for _, c := range cases {
		res, ck, ik, ak, err := c.input.F2345()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(res, c.expected.RES); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/RES", diff)
		}
		if diff := cmp.Diff(ck, c.expected.CK); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/CK", diff)
		}
		if diff := cmp.Diff(ik, c.expected.IK); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/IK", diff)
		}
		if diff := cmp.Diff(ak, c.expected.AK); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/AK", diff)
		}
	}
}

func TestComputeOPc(t *testing.T) {
	c := cases[0]
	got, err := milenage.ComputeOPc(c.input.K[:], c.input.OP[:])
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(got, c.expected.OPc); diff != "" {
		t.Errorf("%s failed: \n%s", c.description, diff)
	}
}

func TestComputeRESStar(t *testing.T) {
	for _, c := range cases {
		if err := c.input.ComputeAll(); err != nil {
			t.Fatal(err)
		}

		resStar, err := c.input.ComputeRESStar("001", "01")
		//resStar, err := c.input.ComputeRESstar()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(resStar, c.expected.RESStar); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/RESStar", diff)
		}
	}
}
