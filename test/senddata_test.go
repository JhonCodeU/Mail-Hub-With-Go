package main

import (
	"github/jhoncodeu/mailbox-masive-go/config"
	"github/jhoncodeu/mailbox-masive-go/src/core"
	"testing"
)

var pathFile = "src/data/output/enron.ndjson"

var pathFolder = "src/data/output/enron.jdjson"

func TestSendRequestToZincsearch(t *testing.T) {
	core.SendRequestToZincSearch(config.UrlBase, pathFile)
}

func BenchmarkSendRequestToZincsearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.SendRequestToZincSearch(config.UrlBase, pathFile)
	}
}
