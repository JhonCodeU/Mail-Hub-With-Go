package main

import (
	"github/jhoncodeu/mailbox-masive-go/src/core"
	"testing"
)

func TestConvertMboxToNdjson(t *testing.T) {
	core.ConvertMboxToNdjson()
}

func BenchmarkConvertMboxToNdjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.ConvertMboxToNdjson()
	}
}
