// Copyright 2018-2021 milenage authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Package milenage provides functions of MILENAGE algorithm set defined in 3GPP TS 35.205.
*/
package milenage

import (
	"crypto/aes"
	"encoding/binary"

	"github.com/pkg/errors"
)

// Milenage is a set of parameters used/generated in MILENAGE algorithm.
type Milenage struct {
	// AK is a 48-bit anonymity key that is the output of either of the functions f5.
	AK []byte
	// AKS is a 48-bit anonymity key that is the output of either of the functions f5*.
	AKS []byte
	// AMF is a 16-bit authentication management field that is an input to the functions f1 and f1*.
	AMF []byte
	// CK is a 128-bit confidentiality key that is the output of the function f3.
	CK []byte
	// IK is a 128-bit integrity key that is the output of the function f4.
	IK []byte
	// K is a 128-bit subscriber key that is an input to the functions f1, f1*, f2, f3, f4, f5 and f5*.
	K []byte
	// MACA is a 64-bit network authentication code that is the output of the function f1.
	MACA []byte
	// MACS is a 64-bit resynchronisation authentication code that is the output of the function f1*.
	MACS []byte
	// OP is a 128-bit Operator Variant Algorithm Configuration Field that is a component of the
	// functions f1, f1*, f2, f3, f4, f5 and f5*.
	OP []byte
	// OPc is a 128-bit value derived from OP and K and used within the computation of the functions.
	OPc []byte
	// RAND is a 128-bit random challenge that is an input to the functions f1, f1*, f2, f3, f4, f5 and f5*.
	RAND []byte
	// RES is a 64-bit signed response that is the output of the function f2.
	RES []byte
	// SQN is a 48-bit sequence number that is an input to either of the functions f1 and f1*.
	// (For f1* this input is more precisely called SQNMS.)
	SQN []byte
	// TEMP is a 128-bit value used within the computation of the functions.
	// TEMP []byte
}

// New initializes a new MILENAGE algorithm.
func New(k, op, rand []byte, sqn uint64, amf uint16) *Milenage {
	m := &Milenage{
		K:    k,
		OP:   op,
		OPc:  nil,
		AMF:  make([]byte, 2),
		SQN:  make([]byte, 6),
		CK:   make([]byte, 16),
		IK:   make([]byte, 16),
		AK:   make([]byte, 6),
		MACA: make([]byte, 8),
		MACS: make([]byte, 8),
		RAND: rand,
		RES:  make([]byte, 8),
	}

	s := make([]byte, 8)
	binary.BigEndian.PutUint64(s, sqn)
	for i := 0; i < 6; i++ {
		m.SQN[i] = s[i+2]
	}

	binary.BigEndian.PutUint16(m.AMF, amf)

	return m
}

// NewWithOPc initializes a new MILENAGE algorithm using OPc instead of OP.
func NewWithOPc(k, opc, rand []byte, sqn uint64, amf uint16) *Milenage {
	m := &Milenage{
		K:    k,
		OP:   nil,
		OPc:  opc,
		AMF:  make([]byte, 2),
		SQN:  make([]byte, 6),
		CK:   make([]byte, 16),
		IK:   make([]byte, 16),
		AK:   make([]byte, 6),
		MACA: make([]byte, 8),
		MACS: make([]byte, 8),
		RAND: rand,
		RES:  make([]byte, 8),
	}

	s := make([]byte, 8)
	binary.BigEndian.PutUint64(s, sqn)
	for i := 0; i < 6; i++ {
		m.SQN[i] = s[i+2]
	}

	binary.BigEndian.PutUint16(m.AMF, amf)

	return m
}

// ComputeOPc is a helper that provides users to retrieve OPc value from
// the K and OP given.
func ComputeOPc(k, op []byte) ([]byte, error) {
	m := New(k, op, make([]byte, 16), 0, 0)
	if err := m.computeOPc(); err != nil {
		return nil, err
	}
	return m.OPc, nil
}

// ComputeAll fills all the fields in *Milenage struct.
func (m *Milenage) ComputeAll() error {
	if _, err := m.F1(); err != nil {
		return errors.Wrap(err, "failed F1())")
	}

	if _, err := m.F1Star(m.SQN, []byte{0x00, 0x00}); err != nil {
		return errors.Wrap(err, "failed F1Star()")
	}

	if _, _, _, _, err := m.F2345(); err != nil {
		return errors.Wrap(err, "failed F2345()")
	}

	return nil
}

// F1 is the network authentication function.
// F1 computes network authentication code MAC-A from key K, random challenge RAND,
// sequence number SQN and authentication management field AMF.
func (m *Milenage) F1() ([]byte, error) {
	mac, err := m.f1base(m.SQN, m.AMF)
	if err != nil {
		return nil, err
	}

	m.MACA = mac[:8]
	return mac[:8], nil
}

// F1Star is the re-synchronisation message authentication function.
// F1Star computes resynch authentication code MAC-S from key K, random challenge RAND,
// sequence number SQN and authentication management field AMF.
//
// Note that the AMF value should be zero to be compliant with the specification
// TS 33.102 6.3.3 (This method just computes with the given value).
func (m *Milenage) F1Star(sqn, amf []byte) ([]byte, error) {
	mac, err := m.f1base(sqn, amf)
	if err != nil {
		return nil, err
	}

	m.MACS = mac[8:]
	return mac[8:], nil
}

// F2345 takes key K and random challenge RAND, and returns response RES,
// confidentiality key CK, integrity key IK and anonymity key AK.
func (m *Milenage) F2345() (res, ck, ik, ak []byte, err error) {
	if m.OPc == nil {
		if err := m.computeOPc(); err != nil {
			return nil, nil, nil, nil, err
		}
	}

	rijndaelInput := make([]byte, 16)
	for i := 0; i < 16; i++ {
		rijndaelInput[i] = m.RAND[i] ^ m.OPc[i]
	}

	temp, err := encrypt(m.K, rijndaelInput)
	if err != nil {
		return
	}

	// To obtain output block OUT2: XOR OPc and TEMP, rotate by r2=0, and XOR on the
	// constant c2 (which is all zeroes except that the last bit is 1).
	for i := 0; i < 16; i++ {
		rijndaelInput[i] = temp[i] ^ m.OPc[i]
	}
	rijndaelInput[15] ^= 1

	out, err := encrypt(m.K, rijndaelInput)
	if err != nil {
		return
	}
	tmp := xor(out, m.OPc)
	res = tmp[8:]
	ak = tmp[:6]

	// To obtain output block OUT3: XOR OPc and TEMP, rotate by r3=32, and XOR on the
	// constant c3 (which is all zeroes except that the next to last bit is 1).
	for i := 0; i < 16; i++ {
		rijndaelInput[(i+12)%16] = temp[i] ^ m.OPc[i]
	}
	rijndaelInput[15] ^= 2

	out, err = encrypt(m.K, rijndaelInput)
	if err != nil {
		return
	}
	ck = xor(out, m.OPc)

	// To obtain output block OUT4: XOR OPc and TEMP, rotate by r4=64, and XOR on the
	// constant c4 (which is all zeroes except that the 2nd from last bit is 1).

	for i := 0; i < 16; i++ {
		rijndaelInput[(i+8)%16] = temp[i] ^ m.OPc[i]
	}
	rijndaelInput[15] ^= 4

	out, err = encrypt(m.K, rijndaelInput)
	if err != nil {
		return
	}
	ik = xor(out, m.OPc)

	m.RES = res
	m.CK = ck
	m.IK = ik
	m.AK = ak
	return res, ck, ik, ak, nil
}

// F5Star is the anonymity key derivation function for the re-synchronisation message.
// F5Star takes key K and random challenge RAND, and returns resynch anonymity key AK.
func (m *Milenage) F5Star() (aks []byte, err error) {
	if m.OPc == nil {
		if err := m.computeOPc(); err != nil {
			return nil, err
		}
	}

	rijndaelInput := make([]byte, 16)
	for i := 0; i < 16; i++ {
		rijndaelInput[i] = m.RAND[i] ^ m.OPc[i]
	}

	tmp, err := encrypt(m.K, rijndaelInput)
	if err != nil {
		return
	}

	// To obtain output block OUT5: XOR OPc and TEMP, rotate by r5=96, and XOR on the
	// constant c5 (which is all zeroes except that the 3rd from last bit is 1).
	for i := 0; i < 16; i++ {
		rijndaelInput[(i+4)%16] = tmp[i] ^ m.OPc[i]
	}
	rijndaelInput[15] ^= 8

	out, err := encrypt(m.K, rijndaelInput)
	if err != nil {
		return
	}

	aks = xor(out, m.OPc)[:6]
	m.AKS = aks
	return aks, nil
}

// computeOPc computes OPc from K and OP inside m.
func (m *Milenage) computeOPc() error {
	m.OPc = make([]byte, 16)

	block, err := aes.NewCipher(m.K)
	if err != nil {
		return err
	}
	cipherText := make([]byte, len(m.OP))
	block.Encrypt(cipherText, m.OP)

	bytes := xor(cipherText, m.OP)
	for i, b := range bytes {
		if i > len(m.OPc) {
			break
		}
		m.OPc[i] = b
	}
	return nil
}

func xor(b1, b2 []byte) []byte {
	var l int
	if len(b1)-len(b2) < 0 {
		l = len(b1)
	} else {
		l = len(b2)
	}

	// don't update b1
	out := make([]byte, l)
	for i := 0; i < l; i++ {
		out[i] = b1[i] ^ b2[i]
	}
	return out
}

func encrypt(key, plain []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, len(plain))
	block.Encrypt(encrypted, plain)
	return encrypted, nil
}

func (m *Milenage) f1base(sqn, amf []byte) ([]byte, error) {
	if m.OPc == nil {
		if err := m.computeOPc(); err != nil {
			return nil, err
		}
	}

	rijndaelInput := make([]byte, 16)
	for i := 0; i < 16; i++ {
		rijndaelInput[i] = m.RAND[i] ^ m.OPc[i]
	}

	temp, err := encrypt(m.K, rijndaelInput)
	if err != nil {
		return nil, err
	}

	in1 := make([]byte, 16)
	for i := 0; i < 6; i++ {
		in1[i] = sqn[i]
		in1[i+8] = sqn[i]
	}
	for i := 0; i < 2; i++ {
		in1[i+6] = amf[i]
		in1[i+14] = amf[i]
	}

	// XOR op_c and in1, rotate by r1=64, and XOR
	// on the constant c1 (which is all zeroes)
	for i := 0; i < 16; i++ {
		rijndaelInput[(i+8)%16] = in1[i] ^ m.OPc[i]
	}
	/* XOR on the value temp computed before */

	for i := 0; i < 16; i++ {
		rijndaelInput[i] ^= temp[i]
	}

	out, err := encrypt(m.K, rijndaelInput)
	if err != nil {
		return nil, err
	}

	return xor(out, m.OPc), nil
}
