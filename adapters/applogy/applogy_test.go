package applogy

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "applogytest", NewApplogyBidder("http://example.com/prebid"))
}
