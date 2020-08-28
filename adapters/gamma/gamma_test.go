package gamma

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "gammatest", NewGammaBidder("https://hb.gammaplatform.com/adx/request/"))
}
