// Copyright 2018-2021 milenage authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package milenage_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/milenage"
)

type expected struct {
	mil  *milenage.Milenage
	autn []byte
	auts []byte
}

var cases = []struct {
	description string
	input       *milenage.Milenage
	*expected
}{
	{
		"withOP/dummy values",
		milenage.New(
			[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			0x000000000001,
			0x8000,
		),
		&expected{
			mil: &milenage.Milenage{
				K:       []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				OP:      []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				OPc:     []byte{0x62, 0xe7, 0x5b, 0x8d, 0x6f, 0xa5, 0xbf, 0x46, 0xec, 0x87, 0xa9, 0x27, 0x6f, 0x9d, 0xf5, 0x4d},
				RAND:    []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				SQN:     []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
				AMF:     []byte{0x80, 0x00},
				MACA:    []byte{0x4a, 0xf3, 0x0b, 0x82, 0xa8, 0x53, 0x11, 0x15},
				MACS:    []byte{0x23, 0xfc, 0x01, 0xba, 0x24, 0x03, 0x13, 0x62},
				RES:     []byte{0x70, 0x0e, 0xb2, 0x30, 0x0b, 0x2c, 0x47, 0x99},
				CK:      []byte{0xb3, 0x79, 0x87, 0x4b, 0x3d, 0x18, 0x3d, 0x2a, 0x21, 0x29, 0x1d, 0x43, 0x9e, 0x77, 0x61, 0xe1},
				IK:      []byte{0xf4, 0x70, 0x6f, 0x66, 0x62, 0x9c, 0xf7, 0xdd, 0xf8, 0x81, 0xd8, 0x00, 0x25, 0xbf, 0x12, 0x55},
				AK:      []byte{0xde, 0x65, 0x6c, 0x8b, 0x0b, 0xce},
				AKS:     []byte{0xb9, 0xac, 0x50, 0xc4, 0x8a, 0x83},
				RESStar: []byte{0x31, 0xb6, 0xd9, 0x38, 0xa5, 0x29, 0x0c, 0xcc, 0x65, 0xbc, 0x82, 0x9f, 0x98, 0x20, 0xa8, 0xd9},
			},
			autn: []byte{0xde, 0x65, 0x6c, 0x8b, 0x0b, 0xcf, 0x80, 0x00, 0x4a, 0xf3, 0x0b, 0x82, 0xa8, 0x53, 0x11, 0x15},
			auts: []byte{0xb9, 0xac, 0x50, 0xc4, 0x8a, 0x82, 0xcd, 0xf7, 0x46, 0x73, 0xbc, 0x86, 0xe7, 0xab},
		},
	}, {
		"withOPc/dummy values",
		milenage.NewWithOPc(
			[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			[]byte{0x62, 0xe7, 0x5b, 0x8d, 0x6f, 0xa5, 0xbf, 0x46, 0xec, 0x87, 0xa9, 0x27, 0x6f, 0x9d, 0xf5, 0x4d},
			[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			0x000000000001,
			0x8000,
		),
		&expected{
			mil: &milenage.Milenage{
				K:       []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				OP:      nil,
				OPc:     []byte{0x62, 0xe7, 0x5b, 0x8d, 0x6f, 0xa5, 0xbf, 0x46, 0xec, 0x87, 0xa9, 0x27, 0x6f, 0x9d, 0xf5, 0x4d},
				RAND:    []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				SQN:     []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
				AMF:     []byte{0x80, 0x00},
				MACA:    []byte{0x4a, 0xf3, 0x0b, 0x82, 0xa8, 0x53, 0x11, 0x15},
				MACS:    []byte{0x23, 0xfc, 0x01, 0xba, 0x24, 0x03, 0x13, 0x62},
				RES:     []byte{0x70, 0x0e, 0xb2, 0x30, 0x0b, 0x2c, 0x47, 0x99},
				CK:      []byte{0xb3, 0x79, 0x87, 0x4b, 0x3d, 0x18, 0x3d, 0x2a, 0x21, 0x29, 0x1d, 0x43, 0x9e, 0x77, 0x61, 0xe1},
				IK:      []byte{0xf4, 0x70, 0x6f, 0x66, 0x62, 0x9c, 0xf7, 0xdd, 0xf8, 0x81, 0xd8, 0x00, 0x25, 0xbf, 0x12, 0x55},
				AK:      []byte{0xde, 0x65, 0x6c, 0x8b, 0x0b, 0xce},
				AKS:     []byte{0xb9, 0xac, 0x50, 0xc4, 0x8a, 0x83},
				RESStar: []byte{0x31, 0xb6, 0xd9, 0x38, 0xa5, 0x29, 0x0c, 0xcc, 0x65, 0xbc, 0x82, 0x9f, 0x98, 0x20, 0xa8, 0xd9},
			},
			autn: []byte{0xde, 0x65, 0x6c, 0x8b, 0x0b, 0xcf, 0x80, 0x00, 0x4a, 0xf3, 0x0b, 0x82, 0xa8, 0x53, 0x11, 0x15},
			auts: []byte{0xb9, 0xac, 0x50, 0xc4, 0x8a, 0x82, 0xcd, 0xf7, 0x46, 0x73, 0xbc, 0x86, 0xe7, 0xab},
		},
	}, {
		"withOP/TS35207-1",
		milenage.New(
			[]byte{0x46, 0x5b, 0x5c, 0xe8, 0xb1, 0x99, 0xb4, 0x9f, 0xaa, 0x5f, 0x0a, 0x2e, 0xe2, 0x38, 0xa6, 0xbc},
			[]byte{0xcd, 0xc2, 0x02, 0xd5, 0x12, 0x3e, 0x20, 0xf6, 0x2b, 0x6d, 0x67, 0x6a, 0xc7, 0x2c, 0xb3, 0x18},
			[]byte{0x23, 0x55, 0x3c, 0xbe, 0x96, 0x37, 0xa8, 0x9d, 0x21, 0x8a, 0xe6, 0x4d, 0xae, 0x47, 0xbf, 0x35},
			0xff9bb4d0b607,
			0xb9b9,
		),
		&expected{
			mil: &milenage.Milenage{
				K:       []byte{0x46, 0x5b, 0x5c, 0xe8, 0xb1, 0x99, 0xb4, 0x9f, 0xaa, 0x5f, 0x0a, 0x2e, 0xe2, 0x38, 0xa6, 0xbc},
				OP:      []byte{0xcd, 0xc2, 0x02, 0xd5, 0x12, 0x3e, 0x20, 0xf6, 0x2b, 0x6d, 0x67, 0x6a, 0xc7, 0x2c, 0xb3, 0x18},
				OPc:     []byte{0xcd, 0x63, 0xcb, 0x71, 0x95, 0x4a, 0x9f, 0x4e, 0x48, 0xa5, 0x99, 0x4e, 0x37, 0xa0, 0x2b, 0xaf},
				RAND:    []byte{0x23, 0x55, 0x3c, 0xbe, 0x96, 0x37, 0xa8, 0x9d, 0x21, 0x8a, 0xe6, 0x4d, 0xae, 0x47, 0xbf, 0x35},
				SQN:     []byte{0xff, 0x9b, 0xb4, 0xd0, 0xb6, 0x07},
				AMF:     []byte{0xb9, 0xb9},
				MACA:    []byte{0x4a, 0x9f, 0xfa, 0xc3, 0x54, 0xdf, 0xaf, 0xb3},
				MACS:    []byte{0x01, 0xcf, 0xaf, 0x9e, 0xc4, 0xe8, 0x71, 0xe9},
				RES:     []byte{0xa5, 0x42, 0x11, 0xd5, 0xe3, 0xba, 0x50, 0xbf},
				CK:      []byte{0xb4, 0x0b, 0xa9, 0xa3, 0xc5, 0x8b, 0x2a, 0x05, 0xbb, 0xf0, 0xd9, 0x87, 0xb2, 0x1b, 0xf8, 0xcb},
				IK:      []byte{0xf7, 0x69, 0xbc, 0xd7, 0x51, 0x04, 0x46, 0x04, 0x12, 0x76, 0x72, 0x71, 0x1c, 0x6d, 0x34, 0x41},
				AK:      []byte{0xaa, 0x68, 0x9c, 0x64, 0x83, 0x70},
				AKS:     []byte{0x45, 0x1e, 0x8b, 0xec, 0xa4, 0x3b},
				RESStar: []byte{0xf2, 0x36, 0xa7, 0x41, 0x72, 0x72, 0xbf, 0xb2, 0xd6, 0x6d, 0x4d, 0x67, 0x07, 0x33, 0xb5, 0x27},
			},
			autn: []byte{0x55, 0xf3, 0x28, 0xb4, 0x35, 0x77, 0xb9, 0xb9, 0x4a, 0x9f, 0xfa, 0xc3, 0x54, 0xdf, 0xaf, 0xb3},
			auts: []byte{0xba, 0x85, 0x3f, 0x3c, 0x12, 0x3c, 0xcf, 0x44, 0xe9, 0x35, 0x96, 0xe3, 0x55, 0xc6},
		},
	},
}

func TestComputeAll(t *testing.T) {
	for _, c := range cases {
		got := c.input
		if err := got.ComputeAll(); err != nil {
			t.Fatal(err)
		}

		resStar := c.expected.mil.RESStar
		c.expected.mil.RESStar = nil
		if diff := cmp.Diff(got, c.expected.mil); diff != "" {
			t.Error(diff)
		}
		c.expected.mil.RESStar = resStar
	}
}

func TestF1(t *testing.T) {
	for _, c := range cases {
		macA, err := c.input.F1()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(macA, c.expected.mil.MACA); diff != "" {
			t.Errorf("%s failed: \n%s", c.description, diff)
		}
	}
}

func TestF1Star(t *testing.T) {
	for _, c := range cases {
		// TS 33.102 6.3.3 says AMF should be zero in F1Star,
		// but the test data 3GPP provides uses it as it is.
		macS, err := c.input.F1Star(c.input.SQN, c.input.AMF)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(macS, c.expected.mil.MACS); diff != "" {
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

		if diff := cmp.Diff(res, c.expected.mil.RES); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/RES", diff)
		}
		if diff := cmp.Diff(ck, c.expected.mil.CK); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/CK", diff)
		}
		if diff := cmp.Diff(ik, c.expected.mil.IK); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/IK", diff)
		}
		if diff := cmp.Diff(ak, c.expected.mil.AK); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/AK", diff)
		}
	}
}

func TestF5Star(t *testing.T) {
	for _, c := range cases {
		aks, err := c.input.F5Star()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(aks, c.expected.mil.AKS); diff != "" {
			t.Errorf("%s failed: \n%s", c.description, diff)
		}
	}
}

func TestComputeOPc(t *testing.T) {
	c := cases[0]
	got, err := milenage.ComputeOPc(c.input.K[:], c.input.OP[:])
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(got, c.expected.mil.OPc); diff != "" {
		t.Errorf("%s failed: \n%s", c.description, diff)
	}
}

func TestComputeRESStar(t *testing.T) {
	for _, c := range cases {
		if err := c.input.ComputeAll(); err != nil {
			t.Fatal(err)
		}

		resStar, err := c.input.ComputeRESStar("001", "01")
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(resStar, c.expected.mil.RESStar); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/RESStar", diff)
		}
	}
}

func TestGenerateAUTN(t *testing.T) {
	for _, c := range cases {
		if err := c.input.ComputeAll(); err != nil {
			t.Fatal(err)
		}

		autn, err := c.input.GenerateAUTN()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(autn, c.expected.autn); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/AUTN", diff)
		}
	}
}

func TestGenerateAUTS(t *testing.T) {
	for _, c := range cases {
		if err := c.input.ComputeAll(); err != nil {
			t.Fatal(err)
		}

		auts, err := c.input.GenerateAUTS()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(auts, c.expected.auts); diff != "" {
			t.Errorf("%s failed: \n%s", c.description+"/AUTS", diff)
		}
	}
}
