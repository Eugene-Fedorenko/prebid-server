package improvedigital

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "improvedigitaltest", NewImprovedigitalBidder("http://localhost/pbs"))
}
