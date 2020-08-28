package adtarget

import (
	"testing"

	"github.com/eugene-fedorenko/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "adtargettest", NewAdtargetBidder("http://ghb.console.adtarget.com.tr/pbs/ortb"))
}
