package main

import (
	"github/jhoncodeu/mailbox-masive-go/src/core"
	"testing"
)

func TestMbox(t *testing.T) {
	core.Mbox()
}

func BenchmarkMbox(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.Mbox()
	}
}
