package yeahmobi

import (
	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
	"testing"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "yeahmobitest", NewYeahmobiBidder("https://{{.Host}}/prebid/bid"))
}
