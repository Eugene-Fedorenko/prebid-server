package lockerdome

import (
	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
	"testing"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "lockerdometest", NewLockerDomeBidder("https://lockerdome.com/ladbid/prebidserver/openrtb2"))
}
