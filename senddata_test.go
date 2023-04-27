package main

import (
	"github/jhoncodeu/mailbox-masive-go/config"
	"github/jhoncodeu/mailbox-masive-go/src/core"
	"testing"
)

var pathFile = "src/data/output/enron.ndjson"

func TestSendRequestToZincsearch(t *testing.T) {
	core.SendRequestToZincsearch(config.UrlBase, pathFile)
}

func BenchmarkSendRequestToZincsearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.SendRequestToZincsearch(config.UrlBase, pathFile)
	}
}
