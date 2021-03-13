# MILENAGE

MILENAGE algorithm implemented in the Go Programming Language.

[![CircleCI](https://circleci.com/gh/wmnsk/milenage.svg?style=shield)](https://circleci.com/gh/wmnsk/milenage)
![golangci-lint](https://github.com/wmnsk/milenage/workflows/golangci-lint/badge.svg)
[![GoDoc](https://godoc.org/github.com/wmnsk/milenage?status.svg)](https://godoc.org/github.com/wmnsk/milenage)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/wmnsk/milenage/blob/master/LICENSE)

## Quickstart

Initialize Milenage first with K, OP, RAND, SQN, and AMF.

```go
mil := milenage.New(
	// K
	[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
	// OP
	[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
	// RAND
	[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
	0x000000000001, // SQN
	0x8000,         // AMF
)
```

Fill all fields(except 5G RES*) at once using `ComputeAll()`.

```go
if err := mil.ComputeAll(); err != nil {
	// ...
}
```

Get MAC-A and MAC-S. This also fills each field.

```go
macA, err := mil.F1()
if err != nil {
	// ...
}

macS, err := mil.F1Star()
if err != nil {
	// ...
}
```

Get RES, CK, IK, AK. This also fills each field.

```go
res, ck, ik, ak, err := mil.F2345()
if err != nil {
	// ...
}
```

Get RES* for 5G with `ComputeRESStar()` by giving MCC and MNC.

```go
resStar, err := mil.ComputeRESStar("001", "01")
if err != nil {
	// ...
}
```

Get OPc from K and OP. This is not the method on `*Milenage`.

```go
opc, err := milenage.ComputeOPc(
	[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
	[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
)
if err != nil {
	// ...
}
```

Get AUTN and AUTS which are used in authentication procedure.
Be sure that the required calculation has done before calling these methods.

```go
autn, err := mil.GenerateAUTN()
if err != nil {
	// ...
}

// Note that this re-calcurates MAC-S and AKS with AMF=0x0000
// as described in 6.3.3, TS 33.102.
auts, err := mil.GenerateAUTS()
if err != nil {
	// ...
}
```

## Author

Yoshiyuki Kurauchi ([Website](https://wmnsk.com/) / [Twitter](https://twitter.com/wmnskdmms))

## License

[MIT](https://github.com/wmnsk/milenage/blob/master/LICENSE)
