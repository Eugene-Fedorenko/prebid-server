package gamoshi

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "gamoshitest", NewGamoshiBidder("https://rtb.gamoshi.io"))
	adapterstest.RunJSONBidderTest(t, "gamoshitest", NewGamoshiBidder(""))
}
