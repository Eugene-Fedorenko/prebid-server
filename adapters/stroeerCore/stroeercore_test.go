package stroeerCore

import (
	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
	"github.com/eugene-fedorenko/prebid-server/config"
	"github.com/eugene-fedorenko/prebid-server/openrtb_ext"
	"testing"
)

func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderStroeerCore, config.Adapter{
		Endpoint: "http://localhost:8361/s2sdsh"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	adapterstest.RunJSONBidderTest(t, "stroeercoretest", bidder)
}
