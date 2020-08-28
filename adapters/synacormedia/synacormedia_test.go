package synacormedia

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "synacormediatest", NewSynacorMediaBidder("http://{{.Host}}.technoratimedia.com/openrtb/bids/{{.Host}}"))
}
