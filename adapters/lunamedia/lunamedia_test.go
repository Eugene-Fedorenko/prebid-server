package lunamedia

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
	"github.com/eugene-fedorenko/prebid-server/config"
	"github.com/eugene-fedorenko/prebid-server/openrtb_ext"
	"github.com/stretchr/testify/assert"
)

func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderLunaMedia, config.Adapter{
		Endpoint: "http://api.lunamedia.io/xp/get?pubid={{.PublisherID}}"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	adapterstest.RunJSONBidderTest(t, "lunamediatest", bidder)
}

func TestEndpointTemplateMalformed(t *testing.T) {
	_, buildErr := Builder(openrtb_ext.BidderLunaMedia, config.Adapter{
		Endpoint: "{{Malformed}}"})

	assert.Error(t, buildErr)
}
