package pubnative

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "pubnativetest", NewPubnativeBidder("http://example.com/prebid"))
}
