package admixer

import (
	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
	"testing"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "admixertest", NewAdmixerBidder("http://inv-nets.admixer.net/pbs.aspx"))
}
