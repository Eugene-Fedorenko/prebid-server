package adoppler

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	bidder := NewAdopplerBidder("http://adoppler.com")
	adapterstest.RunJSONBidderTest(t, "adopplertest", bidder)
}
