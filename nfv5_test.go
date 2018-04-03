/*************************************************************************
 * Copyright 2018 Gravwell, Inc. All rights reserved.
 * Contact: <legal@gravwell.io>
 *
 * This software may be modified and distributed under the terms of the
 * BSD 2-clause license. See the LICENSE file for details.
 **************************************************************************/

package netflow

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

const ()

var (
	bigPkt = []byte{
		0x0, 0x5, 0x0, 0x1e, 0x95, 0x90, 0xf5, 0xc8, 0x5a, 0xc1, 0x7d, 0x28, 0x0, 0x0, 0xc7, 0xae,
		0x0, 0x3, 0x2d, 0xc0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x64, 0x8, 0x8, 0x4, 0x4,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x19, 0x0, 0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x47,
		0x95, 0x90, 0xa3, 0x3c, 0x95, 0x90, 0xa3, 0x3c, 0x87, 0x29, 0x0, 0x35, 0x0, 0x0, 0x11, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x1, 0xa, 0x0, 0x0, 0x64,
		0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0x0, 0x19, 0x0, 0x0, 0x0, 0xc, 0x0, 0x0, 0x4, 0x3e,
		0x95, 0x90, 0x74, 0xc8, 0x95, 0x90, 0xa3, 0x3c, 0x0, 0x0, 0x3, 0x3, 0x0, 0x0, 0x1, 0xc0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x64, 0x8, 0x8, 0x4, 0x4,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x19, 0x0, 0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x47,
		0x95, 0x90, 0xa3, 0x3c, 0x95, 0x90, 0xa3, 0x3c, 0x83, 0x34, 0x0, 0x35, 0x0, 0x0, 0x11, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x64, 0xa, 0x0, 0x0, 0x1,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x19, 0xff, 0xff, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0xf0,
		0x95, 0x90, 0xa3, 0x3c, 0x95, 0x90, 0xa3, 0x3c, 0xb6, 0xa6, 0x0, 0x35, 0x0, 0x0, 0x11, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x64, 0xa, 0x0, 0x0, 0x1,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x19, 0xff, 0xff, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0xf0,
		0x95, 0x90, 0xa3, 0x3c, 0x95, 0x90, 0xa3, 0x3c, 0xa5, 0xae, 0x0, 0x35, 0x0, 0x0, 0x11, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x8, 0x4, 0x4, 0xa, 0x0, 0x0, 0x64,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x19, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x9c,
		0x95, 0x90, 0xa3, 0x64, 0x95, 0x90, 0xa3, 0x64, 0x0, 0x35, 0x83, 0x34, 0x0, 0x0, 0x11, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x64, 0xac, 0xd9, 0xe, 0xae,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x19, 0x0, 0x2, 0x0, 0x0, 0x0, 0x18, 0x0, 0x0, 0x14, 0x54,
		0x95, 0x90, 0x51, 0xa0, 0x95, 0x90, 0xa4, 0x0, 0xec, 0x88, 0x1, 0xbb, 0x0, 0x18, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xac, 0xd9, 0xe, 0xae, 0xa, 0x0, 0x0, 0x64,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x19, 0x0, 0x0, 0x0, 0x1b, 0x0, 0x0, 0xe, 0xc5,
		0x95, 0x90, 0x51, 0xf4, 0x95, 0x90, 0xa4, 0x70, 0x1, 0xbb, 0xec, 0x88, 0x0, 0x18, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x62, 0xa2, 0x7d, 0x22, 0x81,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x19, 0x0, 0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x3, 0xb,
		0x95, 0x90, 0xa5, 0x20, 0x95, 0x90, 0xa5, 0x20, 0xd5, 0x7a, 0x1, 0xbb, 0x0, 0x18, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa2, 0x7d, 0x22, 0x81, 0xa, 0x0, 0x0, 0x62,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x19, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x1, 0x69,
		0x95, 0x90, 0xa5, 0x18, 0x95, 0x90, 0xa5, 0x58, 0x1, 0xbb, 0xd5, 0x7a, 0x0, 0x18, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x62, 0xd0, 0x43, 0xde, 0xde,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x19, 0x0, 0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x52,
		0x95, 0x90, 0xa6, 0x70, 0x95, 0x90, 0xa6, 0x70, 0xb2, 0x4c, 0x0, 0x35, 0x0, 0x0, 0x11, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xd0, 0x43, 0xde, 0xde, 0xa, 0x0, 0x0, 0x62,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x19, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x62,
		0x95, 0x90, 0xa6, 0xc4, 0x95, 0x90, 0xa6, 0xc4, 0x0, 0x35, 0xb2, 0x4c, 0x0, 0x0, 0x11, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x64, 0x8, 0x8, 0x8, 0x8,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x19, 0x0, 0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x47,
		0x95, 0x90, 0xa7, 0x50, 0x95, 0x90, 0xa7, 0x50, 0xe6, 0x90, 0x0, 0x35, 0x0, 0x0, 0x11, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x8, 0x8, 0x8, 0xa, 0x0, 0x0, 0x64,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x19, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x67,
		0x95, 0x90, 0xa7, 0x78, 0x95, 0x90, 0xa7, 0x78, 0x0, 0x35, 0xe6, 0x90, 0x0, 0x0, 0x11, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0xd, 0x4c, 0x65, 0xa, 0xa, 0xa, 0x49,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x5, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x34,
		0x95, 0x90, 0xa8, 0xac, 0x95, 0x90, 0xa8, 0xac, 0x1, 0xbb, 0xf0, 0x29, 0x0, 0x14, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x68, 0x7d, 0xc2, 0xc2, 0xa, 0x0, 0x0, 0x62,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x19, 0x0, 0x0, 0x0, 0x9, 0x0, 0x0, 0x13, 0x56,
		0x95, 0x90, 0xa6, 0xe8, 0x95, 0x90, 0xa9, 0xd8, 0x1, 0xbb, 0xb6, 0xf4, 0x0, 0x1a, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x62, 0x68, 0x7d, 0xc2, 0xc2,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x19, 0x0, 0x2, 0x0, 0x0, 0x0, 0xc, 0x0, 0x0, 0xd, 0xf5,
		0x95, 0x90, 0xa6, 0xc8, 0x95, 0x90, 0xaa, 0xc, 0xb6, 0xf4, 0x1, 0xbb, 0x0, 0x1a, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x34, 0x36, 0xc1, 0x27, 0xa, 0x0, 0x0, 0x64,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x19, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x40, 0xdd,
		0x95, 0x90, 0xa7, 0xe0, 0x95, 0x90, 0xaa, 0xd0, 0x1, 0xbb, 0xbb, 0xec, 0x0, 0x1a, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xd8, 0x3a, 0xda, 0x8e, 0xa, 0x0, 0x0, 0x62,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x19, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x34,
		0x95, 0x90, 0xab, 0x20, 0x95, 0x90, 0xab, 0x20, 0x1, 0xbb, 0xc8, 0xa8, 0x0, 0x10, 0x6, 0x80,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x64, 0x34, 0x36, 0xc1, 0x27,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x19, 0x0, 0x2, 0x0, 0x0, 0x0, 0xc, 0x0, 0x0, 0x5, 0x75,
		0x95, 0x90, 0xa7, 0x78, 0x95, 0x90, 0xaa, 0xfc, 0xbb, 0xec, 0x1, 0xbb, 0x0, 0x1a, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x62, 0xd8, 0x3a, 0xda, 0x8e,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x19, 0x0, 0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x34,
		0x95, 0x90, 0xaa, 0xd4, 0x95, 0x90, 0xaa, 0xd4, 0xc8, 0xa8, 0x1, 0xbb, 0x0, 0x10, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xd0, 0x47, 0x8d, 0x22, 0xa, 0x0, 0x0, 0x64,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x19, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0xd9,
		0x95, 0x90, 0xaf, 0xf4, 0x95, 0x90, 0xaf, 0xf4, 0x1, 0xbb, 0x9a, 0xee, 0x0, 0x18, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x64, 0xd0, 0x47, 0x8d, 0x22,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x19, 0x0, 0x2, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x3, 0x8d,
		0x95, 0x90, 0xaf, 0x60, 0x95, 0x90, 0xaf, 0xf4, 0x9a, 0xee, 0x1, 0xbb, 0x0, 0x18, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x62, 0xd8, 0x3a, 0xda, 0xa3,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x19, 0x0, 0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x34,
		0x95, 0x90, 0xb2, 0xd8, 0x95, 0x90, 0xb2, 0xd8, 0xbc, 0x50, 0x1, 0xbb, 0x0, 0x10, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xd8, 0x3a, 0xda, 0xa3, 0xa, 0x0, 0x0, 0x62,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x19, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x34,
		0x95, 0x90, 0xb3, 0x20, 0x95, 0x90, 0xb3, 0x20, 0x1, 0xbb, 0xbc, 0x50, 0x0, 0x10, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0xa, 0xa, 0x49, 0x17, 0xc8, 0xd1, 0x7e,
		0x48, 0x18, 0xe8, 0x1, 0x0, 0x5, 0x0, 0x2, 0x0, 0x0, 0x0, 0xf8, 0x0, 0x0, 0x4b, 0x1a,
		0x95, 0x8f, 0xc4, 0xec, 0x95, 0x90, 0xb4, 0x8c, 0xf0, 0x47, 0x1, 0xbb, 0x0, 0x18, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x17, 0xc8, 0xd1, 0x7e, 0xa, 0xa, 0xa, 0x49,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x5, 0x0, 0x0, 0x1, 0x1d, 0x0, 0x6, 0x6c, 0xa4,
		0x95, 0x8f, 0xc5, 0x20, 0x95, 0x90, 0xb4, 0x8c, 0x1, 0xbb, 0xf0, 0x47, 0x0, 0x18, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x1, 0xa, 0x0, 0x0, 0x64,
		0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0x0, 0x19, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0xbc,
		0x95, 0x90, 0xb6, 0xb0, 0x95, 0x90, 0xb6, 0xb0, 0x8, 0x1, 0x3, 0x16, 0x0, 0x18, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x64, 0xa, 0x0, 0x0, 0x1,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x19, 0xff, 0xff, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x1, 0x10,
		0x95, 0x90, 0xb6, 0xb0, 0x95, 0x90, 0xb6, 0xb0, 0x3, 0x16, 0x8, 0x1, 0x0, 0x18, 0x6, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0xa, 0xa, 0xd, 0xa, 0xa, 0xa, 0x1,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0xff, 0xff, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x1, 0x48,
		0x95, 0x90, 0xbb, 0x2c, 0x95, 0x90, 0xbb, 0x2c, 0x0, 0x44, 0x0, 0x43, 0x0, 0x0, 0x11, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}
)

func TestValidateBuffer(t *testing.T) {
	var nf NFv5
	if n, err := nf.ValidateSize(bigPkt); err != nil {
		t.Fatal(err)
	} else if len(bigPkt) != n {
		t.Fatal("Invalid validate")
	}
}

func TestDecode(t *testing.T) {
	var nf NFv5
	if err := nf.Decode(bigPkt); err != nil {
		t.Fatal(err)
	}
	if nf.Version != 5 {
		t.Fatal("Invalid version")
	}
	if nf.Count != 30 {
		t.Fatal("Invalid record count")
	}
}

func TestDecodeHeaderAlt(t *testing.T) {
	var nf NFv5Header
	var nfa NFv5Header
	if err := nf.Decode(bigPkt); err != nil {
		t.Fatal(err)
	}
	if err := nfa.DecodeAlt(bigPkt); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkValidateSize(b *testing.B) {
	var nf NFv5
	for i := 0; i < b.N; i++ {
		if n, err := nf.ValidateSize(bigPkt); err != nil {
			b.Fatal(err)
		} else if n != len(bigPkt) {
			b.Fatal("Invalid size")
		}
	}
}

func BenchmarkDecodeHeader(b *testing.B) {
	var nf NFv5Header
	for i := 0; i < b.N; i++ {
		if err := nf.Decode(bigPkt); err != nil {
			b.Fatal(err)
		}
		if nf.Version != 5 {
			b.Fatal("Invalid version")
		}
		if nf.Count != 30 {
			b.Fatal("Invalid record count")
		}
	}
}

func BenchmarkDecodeHeaderAlt(b *testing.B) {
	var nf NFv5Header
	for i := 0; i < b.N; i++ {
		if err := nf.DecodeAlt(bigPkt); err != nil {
			b.Fatal(err)
		}
		if nf.Version != 5 {
			b.Fatal("Invalid version")
		}
		if nf.Count != 30 {
			b.Fatal("Invalid record count")
		}
	}
}

func TestDecodeAlt(t *testing.T) {
	var nf NFv5
	var nfa NFv5
	if err := nf.Decode(bigPkt); err != nil {
		t.Fatal(err)
	}
	if err := nfa.DecodeAlt(bigPkt); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(nf.NFv5Header, nfa.NFv5Header) {
		fmt.Println(nf.NFv5Header)
		fmt.Println(nfa.NFv5Header)
		t.Fatal("Decode header differences")
	}
}

func BenchmarkDecode(b *testing.B) {
	var nf NFv5
	for i := 0; i < b.N; i++ {
		if err := nf.Decode(bigPkt); err != nil {
			b.Fatal(err)
		}
		if nf.Version != 5 {
			b.Fatal("Invalid version")
		}
		if nf.Count != 30 {
			b.Fatal("Invalid record count")
		}
	}
}

func BenchmarkDecodeAlt(b *testing.B) {
	var nf NFv5
	for i := 0; i < b.N; i++ {
		if err := nf.DecodeAlt(bigPkt); err != nil {
			b.Fatal(err)
		}
		if nf.Version != 5 {
			b.Fatal("Invalid version")
		}
		if nf.Count != 30 {
			b.Fatal("Invalid record count")
		}
	}
}

func BenchmarkRead(b *testing.B) {
	var nf NFv5
	bb := bytes.NewBuffer(bigPkt)
	for i := 0; i < b.N; i++ {
		if err := nf.Read(bb); err != nil {
			b.Fatal(err)
		}
		if nf.Version != 5 {
			b.Fatal("Invalid version")
		}
		if nf.Count != 30 {
			b.Fatal("Invalid record count")
		}
		bb = bytes.NewBuffer(bigPkt)
	}
}
