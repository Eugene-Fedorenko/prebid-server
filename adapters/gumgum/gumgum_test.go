package gumgum

import (
	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
	"testing"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "gumgumtest", NewGumGumBidder("https://g2.gumgum.com/providers/prbds2s/bid"))
}
