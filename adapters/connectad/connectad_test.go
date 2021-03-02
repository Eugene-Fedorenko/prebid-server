package connectad

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
	"github.com/eugene-fedorenko/prebid-server/config"
	"github.com/eugene-fedorenko/prebid-server/openrtb_ext"
)

func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderConnectAd, config.Adapter{
		Endpoint: "http://bidder.connectad.io/API?src=pbs"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	adapterstest.RunJSONBidderTest(t, "connectadtest", bidder)
}
