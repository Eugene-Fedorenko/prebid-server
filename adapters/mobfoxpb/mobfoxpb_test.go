package mobfoxpb

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
	"github.com/eugene-fedorenko/prebid-server/config"
	"github.com/eugene-fedorenko/prebid-server/openrtb_ext"
)

func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderMobfoxpb, config.Adapter{
		Endpoint: "http://example.com/?c=__route__&m=__method__&key=__key__"})
	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}
	adapterstest.RunJSONBidderTest(t, "mobfoxpbtest", bidder)
}
