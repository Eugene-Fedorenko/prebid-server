package stroeerCore

import (
	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
	"testing"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "stroeercoretest", NewStroeerCoreBidder("http://localhost:8361/s2sdsh"))
}
