package aja

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
)

const testsBidderEndpoint = "https://localhost/bid/4"

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "ajatest", NewAJABidder(testsBidderEndpoint))
}
