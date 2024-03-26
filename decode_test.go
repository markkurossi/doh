//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package doh

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var queries = []string{
	"q80BAAABAAAAAAAAA3d3dwdleGFtcGxlA2NvbQAAAQAB",
	"AAABAAABAAAAAAAAA3d3dwdleGFtcGxlA2NvbQAAAQAB",
	"AAABAAABAAAAAAAAAWE-NjJjaGFyYWN0ZXJsYWJlbC1tYWtlcy1iYXNlNjR1cmwtZGlzdGluY3QtZnJvbS1zdGFuZGFyZC1iYXNlNjQHZXhhbXBsZQNjb20AAAEAAQ",
}

var decodeOptions = gopacket.DecodeOptions{
	Lazy:   true,
	NoCopy: true,
}

func TestDecode(t *testing.T) {
	for _, q := range queries {
		dst := make([]byte, base64.RawURLEncoding.DecodedLen(len(q)))
		n, err := base64.RawURLEncoding.Decode(dst, []byte(q))
		if err != nil {
			t.Fatalf("decode error: %v", err)
		}
		dst = dst[:n]
		fmt.Printf("decoded:\n%s", hex.Dump(dst))

		packet := gopacket.NewPacket(dst, layers.LayerTypeDNS, decodeOptions)
		layer := packet.Layer(layers.LayerTypeDNS)
		if layer == nil {
			t.Fatalf("non-DNS message")
		}
	}
}
