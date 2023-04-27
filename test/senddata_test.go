package main

import (
	"github/jhoncodeu/mailbox-masive-go/config"
	"github/jhoncodeu/mailbox-masive-go/src/core"
	"testing"
)

var pathFolder = "src/data/output/enron-ndjson"

func TestSendRequestToZincsearch(t *testing.T) {
	core.SendRequestToZincSearch(config.UrlBase, pathFolder)
}

func BenchmarkSendRequestToZincsearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.SendRequestToZincSearch(config.UrlBase, pathFolder)
	}
}
